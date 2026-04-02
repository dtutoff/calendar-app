package cmd

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/SamiRemi/project/app/calendar"
)

type Logger struct {
	entries  []string
	mutex    sync.Mutex
	filePath string
}

func (l *Logger) Log(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.entries = append(l.entries, message)
}

func NewLogger(filePath string) *Logger {
	return &Logger{
		entries:  make([]string, 0),
		filePath: filePath,
	}
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
		logger:   NewLogger("app.logger"),
	}
}

func (l *Logger) SaveToFile() error {
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

func (l *Logger) LoadFromFile() error {
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

func (l *Logger) logWithoutLock(message string) {
	l.entries = append(l.entries, message)
}
