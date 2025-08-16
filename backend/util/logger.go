package util

import (
	"fmt"
	"os"
)

type Logger struct {
	LogLocation string
}

func (l *Logger) LogError(method, message string) {
	msg := fmt.Sprintf("ERROR in method: %s - %s", method, message)
	fmt.Println(msg)
	// Write to some log location
	f, err := os.Create(l.LogLocation)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
