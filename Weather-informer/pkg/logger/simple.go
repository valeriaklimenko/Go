package logger

import "fmt"

type SimpleLogger struct{}

func NewSimpleLogger() *SimpleLogger {
return &SimpleLogger{}
}

func (l *SimpleLogger) Info(msg string) {
fmt.Printf("[INFO] %s\n", msg)
}

func (l *SimpleLogger) Debug(msg string) {
fmt.Printf("[DEBUG] %s\n", msg)
}

func (l *SimpleLogger) Error(msg string) {
fmt.Printf("[ERROR] %s\n", msg)
}
