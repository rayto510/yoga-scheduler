package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func StudioIDMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        studioIDStr := c.GetHeader("X-Studio-ID")
        if studioIDStr == "" {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Studio-ID header missing"})
            return
        }
        id64, err := strconv.ParseUint(studioIDStr, 10, 32)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid X-Studio-ID"})
            return
        }
        c.Set("studio_id", uint(id64))
        c.Next()
    }
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
            c.Abort()
            return
        }

        tokenStr := parts[1]

        token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(*Claims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        // Save user info in context for handlers to use
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Set("studioId", claims.StudioID)

        c.Next()
    }
}
