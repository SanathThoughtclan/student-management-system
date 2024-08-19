package utils

import (
	"log"
	"os"
)

func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func LogInfo(message string, id string) {
	log.Printf("INFO: %s - ID: %s", message, id)
}

func LogError(message string, err error) {
	log.Printf("ERROR: %s - Error: %v", message, err)
}
func LogInfo2(message string) {
	log.Printf("INFO: %s", message)
}

func LogInfo3(message string, id string) {
	log.Printf("INFO: %s - Name: %s", message, id)
}
