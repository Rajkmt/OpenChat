package service

import "meetroom/internal/dto"

// IMeetingService defines the contract for meeting business logic
type IMeetingService interface {
	CreateMeeting(hostName string) (*dto.MeetingResponse, error)
	JoinMeeting(meetingID string, participantName string) (*dto.JoinMeetingResponse, error)
	GetMeeting(meetingID string) (*dto.MeetingDetailsResponse, error)
}
