package storage

import (
	"fmt"
	"os"
)

type JsonStorage struct {
	*Storage
}

func NewStorage(filename string) (*JsonStorage, error) {
	s := &JsonStorage{
		&Storage{filename},
	}

	_, err := os.Stat(s.GetFilename())

	if err == nil {
		return s, nil
	}

	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check file: %w", err)
	}

	if err := os.WriteFile(s.GetFilename(), []byte("{}"), 0644); err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return s, nil
}

func (s *JsonStorage) Save(data []byte) error {
	return os.WriteFile(s.GetFilename(), data, 0644)
}

func (s *JsonStorage) Load() ([]byte, error) {
	return os.ReadFile(s.GetFilename())
}
