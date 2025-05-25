package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"yoga/api/db"
	"yoga/api/handlers"
	"yoga/api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
func TestLoginSuccess(t *testing.T) {
    db.DB = db.SetupTestDB() // <--- This assigns the test DB to the global db.DB used in your handlers

    router := gin.Default()
    router.POST("/login", handlers.Login)

    password := "testpass"
    hashed, _ := models.HashPassword(password)

    user := models.User{
        Email:        "test@example.com",
        PasswordHash: hashed,
        Role:         "owner",
        StudioID:     1,
    }

    db.DB.Create(&user) // now db.DB is properly initialized

    payload := `{"email": "test@example.com", "password": "testpass"}`
    req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "token")
}
