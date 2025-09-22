package util

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	logLocation string
}

func NewLogger(logLocation string) *Logger {
	return &Logger{logLocation: logLocation}
}

func (l *Logger) LogError(method, message string) {
	msg := fmt.Sprintf("%s - ERROR in method: %s - %s", time.Now(), method, message)
	fmt.Println(msg)
	// Write to some log location
	err := os.WriteFile(l.logLocation, []byte(msg), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
