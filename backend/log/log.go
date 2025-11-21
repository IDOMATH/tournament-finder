package log

import (
	"fmt"
	"time"
)

func Event(message string) {
	msg := fmt.Sprintf("%s - %s", time.Now(), message)
	fmt.Println(msg)
}

func Error(method, message string, e error) {
	msg := fmt.Sprintf("%s - ERROR in method: %s - %s", time.Now(), method, message)
	fmt.Println(msg)
	// Write to some log location
	// err := os.WriteFile(l.logLocation, []byte(msg), 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
