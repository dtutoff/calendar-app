package cmd

import (
	"strings"
	"sync"

	"github.com/dtutoff/app/storage"
)

type Logger struct {
	logs    []string
	mu      sync.Mutex
	storage storage.Store
}

func NewLogger(s storage.Store) *Logger {
	logger := &Logger{
		logs:    make([]string, 0),
		storage: s,
	}

	if s != nil {
		data, err := s.Load()
		if err == nil && len(data) > 0 {
			logger.logs = strings.Split(string(data), "\n")
		}
	}

	return logger
}

func (l *Logger) Add(entry string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logs = append(l.logs, entry)

	if l.storage != nil {
		data := []byte(strings.Join(l.logs, "\n"))
		return l.storage.Save(data)
	}

	return nil
}
