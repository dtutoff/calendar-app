package storage

import (
	"fmt"
	"os"
)

type LogStorage struct {
	*Storage
}

func NewLogStorage(filename string) (*LogStorage, error) {
	l := &LogStorage{
		&Storage{filename: filename},
	}

	_, err := os.Stat(l.GetFilename())

	if err == nil {
		return l, nil
	}

	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check file: %w", err)
	}

	if err := os.WriteFile(l.GetFilename(), []byte(""), 0644); err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return l, nil
}

func (l *LogStorage) Save(data []byte) {
	err := os.WriteFile(l.GetFilename(), data, 0644)
	if err != nil {
		return
	}
}

func (l *LogStorage) Load() ([]byte, error) {
	return os.ReadFile(l.GetFilename())
}
