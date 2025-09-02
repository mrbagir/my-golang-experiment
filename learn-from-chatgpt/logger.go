package main

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

type FileLogger struct {
	Writer io.Writer
}

func (f FileLogger) Log(message string) {
	f.Writer.Write([]byte("File: " + message + "\n"))
}

func Process(logger Logger) {
	logger.Log("Starting process")
	// proses lainnya...
	logger.Log("Process finished")
}

func mainLogger() {
	cl := ConsoleLogger{}
	fl := FileLogger{Writer: os.Stdout}

	Process(cl)
	Process(fl)
}
