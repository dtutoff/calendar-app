package storage

import "github.com/SamiRemi/project/app/logger"

type Store interface {
	Save(data []byte) error
	Load() ([]byte, error)
	GetFilename() string
}

type Storage struct {
	filename string
}

func (s *Storage) GetFilename() string {
	logger.Info("Запуск функции GetFilename")
	return s.filename
}
