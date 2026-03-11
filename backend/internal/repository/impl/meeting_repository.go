package impl

import (
	"log"
	"meetroom/internal/repository"
	"sync"
)

// MeetingRepositoryImpl is the in-memory implementation of IMeetingRepository
type MeetingRepositoryImpl struct {
	mu       sync.Mutex
	meetings map[string]*repository.Meeting
}

// NewMeetingRepository creates a new in-memory meeting repository
func NewMeetingRepository() *MeetingRepositoryImpl {
	log.Println("Initializing in-memory meeting repository")
	return &MeetingRepositoryImpl{
		meetings: make(map[string]*repository.Meeting),
	}
}

// Save stores a new meeting in the map
func (r *MeetingRepositoryImpl) Save(meeting *repository.Meeting) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.meetings[meeting.ID] = meeting
	log.Println("Inside Save meeting saved with ID:", meeting.ID)
}

// FindByID looks up a meeting by its ID
func (r *MeetingRepositoryImpl) FindByID(id string) (*repository.Meeting, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	meeting, exists := r.meetings[id]
	if !exists {
		log.Println("Inside FindByID meeting not found:", id)
	}
	return meeting, exists
}

// AddParticipant adds a person's name to the meeting's participant list
func (r *MeetingRepositoryImpl) AddParticipant(meetingID string, name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if meeting, exists := r.meetings[meetingID]; exists {
		meeting.Participants = append(meeting.Participants, name)
		log.Println("Inside AddParticipant", name, "added to meeting:", meetingID)
	}
}
