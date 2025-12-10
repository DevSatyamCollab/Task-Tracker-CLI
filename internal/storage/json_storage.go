package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-list/internal/core"
)

func GetStorage() (*Storage, error) {
	var StorageErr error

	once.Do(func() {
		file, err := unsureStorage()

		if err != nil {
			StorageErr = fmt.Errorf("failed to initialize storage: %w", err)
			instance = nil
		} else {
			instance = &Storage{fileName: file}
		}

	})

	return instance, StorageErr
}

// save method
func (s *Storage) Save(tracker *core.TaskTracker) error {
	if s.fileName == "" {
		return errEmptyStorage
	}

	jsonData, err := json.MarshalIndent(tracker.Tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(s.fileName, jsonData, 0200); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func (s *Storage) Load(tracker *core.TaskTracker) error {
	if s.fileName == "" {
		return errEmptyStorage
	}

	// read
	content, err := os.ReadFile(s.fileName)
	if err != nil {
		return fmt.Errorf("failed to read the file: %w", err)
	}

	if len(content) == 0 {
		return nil
	}

	// unmarshal
	if err = json.Unmarshal(content, &tracker.Tasks); err != nil {
		return fmt.Errorf("failed to load the data: %w", err)
	}

	// update the nextId
	tracker.NextId = tracker.UpdateNextID()

	return nil
}
