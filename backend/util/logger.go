package util

import "fmt"

type Logger struct {
	LogLocation string
}

func (l *Logger) LogError(method, message string) {
	msg := fmt.Sprintf("ERROR in method: %s - %s", method, message)
	fmt.Println(msg)
	// Write to some log location
}
