package repository

import "time"

// Meeting represents a single meeting room stored in the repository
type Meeting struct {
	ID           string
	HostName     string
	IsActive     bool
	Participants []string
	CreatedAt    time.Time
}

// IMeetingRepository defines the contract for meeting data access
type IMeetingRepository interface {
	Save(meeting *Meeting)
	FindByID(id string) (*Meeting, bool)
	AddParticipant(meetingID string, name string)
}
