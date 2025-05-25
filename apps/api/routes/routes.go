package routes

import (
	"yoga/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Instructors routes
	r.GET("/instructors", handlers.GetInstructors)
	r.GET("/instructors/:id", handlers.GetInstructor)
	r.GET("/instructors/:id/classes", handlers.GetClassesByInstructor)

	// Classes routes
	r.GET("/classes", handlers.GetClasses)
	r.GET("/classes/:id", handlers.GetClass)

    // Locations routes
    r.GET("/locations", handlers.GetLocations)
    r.GET("/locations/:id", handlers.GetLocation)

    // Auth routes
    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    // Protected routes
    authorized := r.Group("/")
    authorized.Use(handlers.AuthMiddleware())
    {
        authorized.POST("/instructors", handlers.CreateInstructor)
        authorized.PUT("/instructors/:id", handlers.UpdateInstructor)
        authorized.DELETE("/instructors/:id", handlers.DeleteInstructor)

        authorized.POST("/classes", handlers.CreateClass) // Only logged in users can create classes
        authorized.PUT("/classes/:id", handlers.UpdateClass)
        authorized.DELETE("/classes/:id", handlers.DeleteClass)

        authorized.POST("/locations", handlers.CreateLocation)
        authorized.PUT("/locations/:id", handlers.UpdateLocation)
        authorized.DELETE("/locations/:id", handlers.DeleteLocation)
    }
}