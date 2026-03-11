package controller

import (
	"log"
	"net/http"

	"meetroom/internal/dto"
	"meetroom/internal/service"

	"github.com/gin-gonic/gin"
)

// MeetingController handles HTTP requests for meetings
type MeetingController struct {
	MeetingService service.IMeetingService
}

// NewMeetingController creates a new controller with the given service
func NewMeetingController(meetingService service.IMeetingService) *MeetingController {
	return &MeetingController{
		MeetingService: meetingService,
	}
}

// CreateMeeting handles POST /api/v1/meetings
func (mc *MeetingController) CreateMeeting() gin.HandlerFunc {
	return func(c *gin.Context) {
		methodName := "CreateMeeting"
		log.Println("Inside", methodName)

		var request dto.CreateMeetingRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("Inside", methodName, "invalid request body:", err)
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error: "host_name is required",
			})
			return
		}

		response, err := mc.MeetingService.CreateMeeting(request.HostName)
		if err != nil {
			log.Println("Inside", methodName, "service error:", err)
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		log.Println("Inside", methodName, "meeting created:", response.MeetingID)
		c.JSON(http.StatusCreated, response)
	}
}

// JoinMeeting handles POST /api/v1/meetings/:id/join
func (mc *MeetingController) JoinMeeting() gin.HandlerFunc {
	return func(c *gin.Context) {
		methodName := "JoinMeeting"
		meetingID := c.Param("id")
		log.Println("Inside", methodName, "meeting:", meetingID)

		var request dto.JoinMeetingRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("Inside", methodName, "invalid request body:", err)
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error: "participant_name is required",
			})
			return
		}

		response, err := mc.MeetingService.JoinMeeting(meetingID, request.ParticipantName)
		if err != nil {
			log.Println("Inside", methodName, "service error:", err)
			statusCode := http.StatusInternalServerError
			if err.Error() == "meeting not found" {
				statusCode = http.StatusNotFound
			} else if err.Error() == "meeting has ended" {
				statusCode = http.StatusGone
			}
			c.JSON(statusCode, dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		log.Println("Inside", methodName, "participant joined:", meetingID)
		c.JSON(http.StatusOK, response)
	}
}

// GetMeeting handles GET /api/v1/meetings/:id
func (mc *MeetingController) GetMeeting() gin.HandlerFunc {
	return func(c *gin.Context) {
		methodName := "GetMeeting"
		meetingID := c.Param("id")
		log.Println("Inside", methodName, "meeting:", meetingID)

		response, err := mc.MeetingService.GetMeeting(meetingID)
		if err != nil {
			log.Println("Inside", methodName, "service error:", err)
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
