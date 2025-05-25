package handlers

import (
	"net/http"
	"strconv"
	"yoga/api/db"
	"yoga/api/models"

	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location.StudioID = studioID

	if err := db.DB.Create(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
		return
	}

	c.JSON(http.StatusCreated, location)
}

func GetLocations(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)
	var locations []models.Location
	if err := db.DB.Where("studio_id = ?", studioID).Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch locations"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func GetLocation(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	var location models.Location
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func UpdateLocation(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	// Fetch location scoped by studio
	var location models.Location
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).First(&location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	var input models.Location
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update allowed fields only
	location.Name = input.Name
	location.Address = input.Address
	// Do NOT update StudioID on update to avoid assignment errors

	if err := db.DB.Save(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func DeleteLocation(c *gin.Context) {
	studioID := c.MustGet("studio_id").(uint)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	// Ensure we only delete location belonging to this studio
	if err := db.DB.Where("id = ? AND studio_id = ?", id, studioID).Delete(&models.Location{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete location"})
		return
	}

	c.Status(http.StatusNoContent)
}
