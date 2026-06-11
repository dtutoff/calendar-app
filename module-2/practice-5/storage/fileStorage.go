package storage

import (
	"fmt"
	"os"
)

type FileStorage struct {
	*Storage
}

func NewFileStorage(filename string, initData []byte) (*FileStorage, error) {
	s := &FileStorage{
		&Storage{filename},
	}

	_, err := os.Stat(s.GetFilename())

	if err == nil {
		return s, nil
	}

	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check file: %w", err)
	}

	if err := os.WriteFile(s.GetFilename(), initData, 0644); err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return s, nil
}

func NewJSONFileStorage(filename string) (Store, error) {
	return NewFileStorage(filename, []byte("{}"))
}

func NewTextFileStorage(filename string) (Store, error) {
	return NewFileStorage(filename, []byte(""))
}

func (f *FileStorage) Save(data []byte) error {
	return os.WriteFile(f.GetFilename(), data, 0644)
}

func (f *FileStorage) Load() ([]byte, error) {
	return os.ReadFile(f.GetFilename())
}
