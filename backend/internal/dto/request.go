package dto

// CreateMeetingRequest is the JSON body sent by frontend when creating a meeting
type CreateMeetingRequest struct {
	HostName string `json:"host_name" binding:"required"`
}

// JoinMeetingRequest is the JSON body sent by frontend when joining a meeting
type JoinMeetingRequest struct {
	ParticipantName string `json:"participant_name" binding:"required"`
}
