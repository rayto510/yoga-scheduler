package handlers

import (
	"net/http"
	"strconv"
	"yoga/api/db"
	"yoga/api/models"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
    studioID := c.MustGet("studio_id").(uint)
	var classes []models.Class
	if err := db.DB.Where("studio_id = ?", studioID).Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func CreateClass(c *gin.Context) {
    studioID := c.MustGet("studio_id").(uint)
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var instructor models.Instructor
	if err := db.DB.Where("id = ? AND studio_id = ?", class.InstructorID, studioID).First(&instructor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instructor not found or doesn't belong to your studio"})
		return
	}

	class.StudioID = studioID
	if err := db.DB.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}

	c.JSON(http.StatusOK, class)
}

func GetClass(c *gin.Context) {
    studioID := c.MustGet("studio_id").(uint)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	var class models.Class
	if err := db.DB.Where("studio_id = ?", studioID).First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	c.JSON(http.StatusOK, class)
}

func UpdateClass(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	// Parse class ID from URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	// Fetch the existing class for this studio
	var existing models.Class
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&existing).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	// Parse input JSON into a temporary struct
	var input models.Class
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: validate that the new instructor belongs to the studio
	if input.InstructorID != 0 && input.InstructorID != existing.InstructorID {
		var instructor models.Instructor
		if err := db.DB.Where("id = ? AND studio_id = ?", input.InstructorID, studioID).First(&instructor).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Instructor not found or doesn't belong to your studio"})
			return
		}
	}

	// Update only fields allowed to change
	existing.Name = input.Name
	existing.Description = input.Description
	existing.InstructorID = input.InstructorID
	existing.StartTime = input.StartTime
	existing.EndTime = input.EndTime

	if err := db.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update class"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func DeleteClass(c *gin.Context) {
    studioID := c.MustGet("studio_id").(uint)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	var class models.Class
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&class).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	if err := db.DB.Delete(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete class"})
		return
	}

	c.Status(http.StatusNoContent)
}

