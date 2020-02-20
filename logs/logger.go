package logger

import (
	"fmt"
	"log"
	"os"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// RequestLogger exported
var RequestLogger *log.Logger

func init() {
	generalLog, err := os.OpenFile("../logs/gcp.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	requestLog, err := os.OpenFile("../logs/requests.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	RequestLogger = log.New(requestLog, "Request Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
