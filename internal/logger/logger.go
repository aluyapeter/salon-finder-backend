package logger

import "log"

type Logger interface{
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type stdLogger struct{}

// func New() *Logger {
// 	return &Logger{}
// }

func (l *stdLogger) Info(msg string, args ...interface{}) {
	log.Printf("[INFO]"+ msg, args...)
}

func (l *stdLogger) Error(msg string, args ...interface{}) {
	log.Printf("[ERROR]"+msg, args...)
}
func (l *stdLogger) Fatal(msg string, args ...interface{}) {
	log.Fatalf("[FATAL]"+msg, args...)
}

func New() Logger {
	return &stdLogger{}
}