package storage

import (
	"log"
	"os"
)

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	s := &Storage{filename}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := os.WriteFile(filename, []byte("{}"), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	return s
}

func (s *Storage) Save(data []byte) error {
	return os.WriteFile(s.filename, data, 0644)
}

func (s *Storage) Load() ([]byte, error) {
	return os.ReadFile(s.filename)
}
