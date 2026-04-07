package cmd

import (
	"encoding/json"
	"os"
	"sync"
)

type Log struct {
	entries  []string
	mutex    sync.Mutex
	filePath string
}

func (l *Log) Log(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.entries = append(l.entries, message)
}

func NewLogger(filePath string) *Log {
	return &Log{
		entries:  make([]string, 0),
		filePath: filePath,
	}
}

func (l *Log) SaveToFile() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	data, err := json.MarshalIndent(l.entries, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(l.filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (l *Log) LoadFromFile() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	data, err := os.ReadFile(l.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var loadedEntries []string
	err = json.Unmarshal(data, &loadedEntries)
	if err != nil {
		return err
	}

	l.entries = append(l.entries, loadedEntries...)
	return nil
}
