package router

import (
	"meetroom/internal/controller"
	"meetroom/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter sets up all routes and returns the Gin engine
func NewRouter(meetingController *controller.MeetingController, frontendURL string) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := router.Group("/api/v1")
	{
		api.POST("/meetings", meetingController.CreateMeeting())
		api.GET("/meetings/:id", meetingController.GetMeeting())
		api.POST("/meetings/:id/join", meetingController.JoinMeeting())
	}

	return router
}
