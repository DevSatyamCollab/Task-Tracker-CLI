package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"todo-list/internal/core"
)

type IStorage interface {
	Save(tracker *core.TaskTracker) error
	Load(tracker *core.TaskTracker) error
}

const (
	dirPath string = "todo-list/data/"
	fname   string = "data.json"
)

type Storage struct {
	fileName string
}

var (
	instance        *Storage
	once            sync.Once
	errEmptyStorage error = errors.New("storage filepath is empty")
)

// unsure storage
func unsureStorage() (string, error) {

	// check the stat of dir
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			return "", fmt.Errorf("failed to create the directory: %w", err)
		}
	}

	// custom filepath
	filePath := filepath.Join(dirPath, fname)

	// check the stat of the file
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to create a file: %w", err)
		}

		defer file.Close()
	}

	return filePath, nil
}
