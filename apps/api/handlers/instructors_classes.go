package handlers

import (
	"net/http"
	"strconv"
	"yoga/api/db"
	"yoga/api/models"

	"github.com/gin-gonic/gin"
)

// GetClassesByInstructor returns all classes taught by a specific instructor within the studio.
func GetClassesByInstructor(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	// Check if instructor exists within this studio
	var instructor models.Instructor
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&instructor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	// Fetch classes for the instructor within the studio
	var classes []models.Class
	if err := db.DB.Where("instructor_id = ? AND studio_id = ?", id, studioID).Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query classes: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}
