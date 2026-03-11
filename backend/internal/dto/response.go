package dto

// MeetingResponse is returned after creating a meeting
type MeetingResponse struct {
	MeetingID  string `json:"meeting_id"`
	Token      string `json:"token,omitempty"`
	LiveKitURL string `json:"livekit_url,omitempty"`
	Message    string `json:"message"`
}

// JoinMeetingResponse is returned after joining a meeting
type JoinMeetingResponse struct {
	Token      string `json:"token"`
	LiveKitURL string `json:"livekit_url"`
	Message    string `json:"message"`
}

// MeetingDetailsResponse is returned when checking meeting info
type MeetingDetailsResponse struct {
	MeetingID    string   `json:"meeting_id"`
	HostName     string   `json:"host_name"`
	IsActive     bool     `json:"is_active"`
	Participants []string `json:"participants"`
	CreatedAt    string   `json:"created_at"`
}

// ErrorResponse is the standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}
