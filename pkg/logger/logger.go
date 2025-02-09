package logger

import (
	"fmt"
	"log"
)

// Observer interface
type LogObserver interface {
	Log(level, message string)
}

// ConsoleLogger (default observer)
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(level, message string) {
	log.Printf("[%s] %s", level, message)
}

// Logger Manager
type LogEventManager struct {
	observers []LogObserver
}

func (l *LogEventManager) Register(observer LogObserver) {
	l.observers = append(l.observers, observer)
}

func (l *LogEventManager) Notify(level, message string) {
	for _, observer := range l.observers {
		observer.Log(level, message)
	}
}

// Global Logger
var Logger *LogEventManager

func init() {
	Logger = &LogEventManager{}
	Logger.Register(ConsoleLogger{}) // Default console logger
}

// Usage functions
func Info(msg string, args ...interface{}) {
	Logger.Notify("INFO", formatMessage(msg, args...))
}

func Error(msg string, args ...interface{}) {
	Logger.Notify("ERROR", formatMessage(msg, args...))
}

// Format message
func formatMessage(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
}
