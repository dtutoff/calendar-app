package storage

import "os"

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (s *Storage) Save(data []byte) error {
	err := os.WriteFile(s.filename, data, 0644)
	return err
}

func (s *Storage) Load() ([]byte, error) {
	err, data := os.ReadFile(s.filename)
	return err, data
}
