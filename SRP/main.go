package main

import (
    "fmt"
)

type FileLogger struct{}

func (f FileLogger) Log(message string) {
    fmt.Println("Logging to file:", message)
}

func main() {
    logger := FileLogger{}
    logger.Log("Application started")
}