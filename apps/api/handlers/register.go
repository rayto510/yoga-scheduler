package handlers

import (
	"net/http"
	"os"
	"time"

	"yoga/api/db"
	"yoga/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if user already exists
    var existing models.User
    if err := db.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Create studio
    studio := models.Studio{Name: input.Email + "'s Studio"}
    if err := db.DB.Create(&studio).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create studio"})
        return
    }

    // Create user
    user := models.User{
        Email:        input.Email,
        PasswordHash: string(hashedPassword),
        StudioID:     studio.ID,
    }

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    // Create JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  user.ID,
        "studio_id": user.StudioID,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
        "user": gin.H{
            "id":        user.ID,
            "email":     user.Email,
            "studio_id": user.StudioID,
        },
    })
}
