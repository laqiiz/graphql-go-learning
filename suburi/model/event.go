package model

import (
	"errors"
	"github.com/google/uuid"
	"os/user"
	"time"
)

type Event struct {
	EventId        string       `json:"eventId"`
	User           *user.User   `json:"user"`
	EventName      string       `json:"eventName"`
	Description    string       `json:"description"`
	Location       string       `json:"location"`
	StartTime      string       `json:"startTime"`
	EndTime        string       `json:"endTime"`
	Participants   []*user.User `json:"participants"`
	RegisteredTime int64        `json:"registeredTime"`
}

func NewEvent(givenUser *user.User, eventName, description, location, startTime, endTime string) (*Event, error) {
	if eventName == "" {
		return nil, errors.New("eventName is empty")
	}
	if description == "" {
		return nil, errors.New("description is empty")
	}
	if location == "" {
		return nil, errors.New("location is empty")
	}
	if startTime == "" {
		return nil, errors.New("startTime is empty")
	}
	if endTime == "" {
		return nil, errors.New("endTime is empty")
	}
	if givenUser == nil {
		return nil, errors.New("user is nil")
	}

	return &Event{
		EventId:        uuid.New().String(),
		User:           givenUser,
		EventName:      eventName,
		Description:    description,
		Location:       location,
		StartTime:      startTime,
		EndTime:        endTime,
		Participants:   []*user.User{},
		RegisteredTime: time.Now().Unix()}, nil
}

func (event *Event) Equals(other *Event) bool {
	return event.EventId == other.EventId
}
