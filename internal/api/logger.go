package api

import (
	"io"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLoggers() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multi := io.MultiWriter(os.Stdout, file)

	InfoLogger = log.New(multi, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(multi, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multi, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// output to both file and console
	// InfoLogger.SetOutput(multi)
	// WarningLogger.SetOutput(multi)
	// ErrorLogger.SetOutput(multi)

}
