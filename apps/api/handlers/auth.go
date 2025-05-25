package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-very-secret-key") // Put in env var in prod!

// Claims struct
type Claims struct {
    UserID   uint   `json:"user_id"`
    Role     string `json:"role"`
    StudioID uint   `json:"studio_id"`
    jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(userID uint, role string, studioId uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        Role:   role,
        StudioID: studioId,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // token expires in 72h
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "yoga-app",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
