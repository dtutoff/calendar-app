package storage

import (
	"os"

	"github.com/SamiRemi/project/app/logger"
)

type JsonStorage struct {
	*Storage
}

func NewJsonStorage(filename string) *JsonStorage {
	logger.Info("Запуск функции  NewJsonStorage")
	return &JsonStorage{
		&Storage{filename: filename},
	}
}

func (s *JsonStorage) Save(data []byte) error {
	logger.Info("Запуск функции  Save")
	err := os.WriteFile(s.GetFilename(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *JsonStorage) Load() ([]byte, error) {
	logger.Info("Запуск функции  Load")
	data, err := os.ReadFile(s.GetFilename())
	if err != nil {
		return nil, err
	}
	return data, nil
}
