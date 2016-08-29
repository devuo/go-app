package service

import "fmt"

type logger struct {
	Debug string
}

// Info logs an info log
func (l *logger) Info(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// Err logs an error if the error is not nil
func (l *logger) Err(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
