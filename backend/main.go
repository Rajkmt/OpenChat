package main

import (
	"log"
	"os"

	"meetroom/internal/controller"
	repoImpl "meetroom/internal/repository/impl"
	"meetroom/internal/router"
	serviceImpl "meetroom/internal/service/impl"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("========================================")
	log.Println("  OpenChat Backend Starting...")
	log.Println("========================================")

	// Load .env file for local development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	liveKitURL := getEnv("LIVEKIT_URL", "")
	liveKitAPIKey := getEnv("LIVEKIT_API_KEY", "")
	liveKitSecret := getEnv("LIVEKIT_API_SECRET", "")
	frontendURL := getEnv("FRONTEND_URL", "http://localhost:5173")
	port := getEnv("PORT", "8080")

	// Fail fast if credentials are missing
	if liveKitURL == "" || liveKitAPIKey == "" || liveKitSecret == "" {
		log.Fatal("Missing required environment variables: LIVEKIT_URL, LIVEKIT_API_KEY, LIVEKIT_API_SECRET")
	}

	log.Println("LiveKit URL: ", liveKitURL)
	log.Println("Frontend URL:", frontendURL)
	log.Println("Port:        ", port)

	meetingRepo := repoImpl.NewMeetingRepository()
	meetingService := serviceImpl.NewMeetingService(meetingRepo, liveKitAPIKey, liveKitSecret, liveKitURL)
	meetingController := controller.NewMeetingController(meetingService)

	r := router.NewRouter(meetingController, frontendURL)

	log.Println("Server listening on port", port)
	log.Println("========================================")

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// getEnv reads an environment variable, returns fallback if not set
func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
