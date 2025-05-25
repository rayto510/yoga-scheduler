package handlers

import (
	"net/http"
	"strconv"
	"yoga/api/db"
	"yoga/api/models"

	"github.com/gin-gonic/gin"
)

func CreateInstructor(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	var instructor models.Instructor
	if err := c.ShouldBindJSON(&instructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructor.StudioID = studioID

	if err := db.DB.Create(&instructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create instructor"})
		return
	}

	c.JSON(http.StatusCreated, instructor)
}

func GetInstructor(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	var instructor models.Instructor
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&instructor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	c.JSON(http.StatusOK, instructor)
}

func UpdateInstructor(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	// Fetch existing instructor
	var existing models.Instructor
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&existing).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	// Bind input JSON to separate struct
	var input models.Instructor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update allowed fields explicitly
	existing.FirstName = input.FirstName
	existing.LastName = input.LastName
	existing.Bio = input.Bio

	if err := db.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update instructor"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func DeleteInstructor(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	var instructor models.Instructor
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&instructor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	if err := db.DB.Delete(&instructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete instructor"})
		return
	}

	c.Status(http.StatusNoContent)
}

func GetInstructors(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	var instructors []models.Instructor
	if err := db.DB.Where("studio_id = ?", studioID).Find(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch instructors"})
		return
	}

	c.JSON(http.StatusOK, instructors)
}