package core

import "time"

const (
	PENDING   string = "❌"
	COMPLETED string = "✅"
	PROGRESS  string = "⌛"
)

type Task struct {
	Id          int       `json:"ID"`
	CreatedAt   time.Time `json:"CreatedAt"`
	Description string    `json:"Description"`
	Status      string    `json:"Status"`
	CompletedAt time.Time `json:"CompletedAt"`
}

func NewTask(id int, desc string) *Task {
	return &Task{
		Id:          id,
		CreatedAt:   time.Now(),
		Description: desc,
		Status:      PENDING,
	}
}

// validate Description
func ValidateDescription(desc string) error {
	if desc == "" {
		return ErrEmptyDescription
	}

	return nil
}

// validate Id
func ValidateID(id int) error {
	if id < 0 {
		return ErrInvalidId
	}

	return nil
}
