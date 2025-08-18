package util

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	LogLocation string
}

func (l *Logger) LogError(method, message string) {
	msg := fmt.Sprintf("%s - ERROR in method: %s - %s", time.Now(), method, message)
	fmt.Println(msg)
	// Write to some log location
	err := os.WriteFile(l.LogLocation, []byte(msg), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
