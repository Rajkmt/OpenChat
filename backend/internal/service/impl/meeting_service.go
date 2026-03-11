package impl

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"meetroom/internal/dto"
	"meetroom/internal/repository"

	"github.com/livekit/protocol/auth"
)

// MeetingServiceImpl contains the business logic for meetings
type MeetingServiceImpl struct {
	Repo          repository.IMeetingRepository
	LiveKitAPIKey string
	LiveKitSecret string
	LiveKitURL    string
}

// NewMeetingService creates a new MeetingServiceImpl with all dependencies
func NewMeetingService(
	repo repository.IMeetingRepository,
	apiKey string,
	apiSecret string,
	livekitURL string,
) *MeetingServiceImpl {
	log.Println("Initializing meeting service")
	return &MeetingServiceImpl{
		Repo:          repo,
		LiveKitAPIKey: apiKey,
		LiveKitSecret: apiSecret,
		LiveKitURL:    livekitURL,
	}
}

// CreateMeeting creates a new meeting room and returns meeting details with host token
func (s *MeetingServiceImpl) CreateMeeting(hostName string) (*dto.MeetingResponse, error) {
	methodName := "CreateMeeting"
	log.Println("Inside", methodName, "host:", hostName)

	meetingID := generateMeetingID()

	meeting := &repository.Meeting{
		ID:           meetingID,
		HostName:     hostName,
		IsActive:     true,
		Participants: []string{hostName},
		CreatedAt:    time.Now(),
	}
	s.Repo.Save(meeting)
	log.Println("Inside", methodName, "meeting created with ID:", meetingID)

	token, err := s.generateToken(meetingID, hostName)
	if err != nil {
		log.Println("Inside", methodName, "token generation failed:", err)
		return nil, errors.New("failed to generate access token")
	}

	return &dto.MeetingResponse{
		MeetingID:  meetingID,
		Token:      token,
		LiveKitURL: s.LiveKitURL,
		Message:    "Meeting created successfully",
	}, nil
}

// JoinMeeting lets a participant join an existing meeting
func (s *MeetingServiceImpl) JoinMeeting(meetingID string, participantName string) (*dto.JoinMeetingResponse, error) {
	methodName := "JoinMeeting"
	log.Println("Inside", methodName, "meeting:", meetingID, "participant:", participantName)

	meeting, exists := s.Repo.FindByID(meetingID)
	if !exists {
		log.Println("Inside", methodName, "meeting not found:", meetingID)
		return nil, errors.New("meeting not found")
	}

	if !meeting.IsActive {
		log.Println("Inside", methodName, "meeting has ended:", meetingID)
		return nil, errors.New("meeting has ended")
	}

	s.Repo.AddParticipant(meetingID, participantName)

	token, err := s.generateToken(meetingID, participantName)
	if err != nil {
		log.Println("Inside", methodName, "token generation failed:", err)
		return nil, errors.New("failed to generate access token")
	}

	log.Println("Inside", methodName, participantName, "joined meeting:", meetingID)
	return &dto.JoinMeetingResponse{
		Token:      token,
		LiveKitURL: s.LiveKitURL,
		Message:    "Joined meeting successfully",
	}, nil
}

// GetMeeting returns the details of a meeting
func (s *MeetingServiceImpl) GetMeeting(meetingID string) (*dto.MeetingDetailsResponse, error) {
	methodName := "GetMeeting"
	log.Println("Inside", methodName, "meeting:", meetingID)

	meeting, exists := s.Repo.FindByID(meetingID)
	if !exists {
		log.Println("Inside", methodName, "meeting not found:", meetingID)
		return nil, errors.New("meeting not found")
	}

	return &dto.MeetingDetailsResponse{
		MeetingID:    meeting.ID,
		HostName:     meeting.HostName,
		IsActive:     meeting.IsActive,
		Participants: meeting.Participants,
		CreatedAt:    meeting.CreatedAt.Format(time.RFC3339),
	}, nil
}

// generateToken creates a LiveKit JWT token for a participant
func (s *MeetingServiceImpl) generateToken(meetingID string, participantName string) (string, error) {
	log.Println("Inside generateToken for", participantName, "in room", meetingID)

	at := auth.NewAccessToken(s.LiveKitAPIKey, s.LiveKitSecret)

	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     meetingID,
	}

	at.SetIdentity(participantName).
		SetValidFor(24 * time.Hour).
		AddGrant(grant)

	token, err := at.ToJWT()
	if err != nil {
		return "", err
	}

	log.Println("Inside generateToken token generated for", participantName)
	return token, nil
}

// generateMeetingID creates a random 8-character meeting ID
func generateMeetingID() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
