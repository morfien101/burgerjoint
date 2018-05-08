package logger

import (
	"fmt"
	"sync"
)

// Logger is actions the internal logger needs to satisfy
type Logger interface {
	Print(string)
	Printf(string, ...interface{})
}

// Toggle allows us to turn the logger on and off
type Toggle interface {
	Toggle(bool)
}

// FullLogger is all the actions a logger can perform
type FullLogger interface {
	Logger
	Toggle
	Drain()
}

// InternalLogger is a logger we can use with trace messages.
// It will serialise the messages to normizile the order from go routines.
type InternalLogger struct {
	mu        sync.RWMutex
	on        bool
	printChan chan string
	Done      chan interface{}
	print     func(string) string
	printf    func(string, ...interface{}) string
}

// Print will run the print function and pass the result to serial printer
func (l *InternalLogger) Print(s string) {
	if l.output() {
		l.printChan <- l.print(s)
	}
}

// Printf will format the message by calling printf and then pass the result to the serial printer.
func (l *InternalLogger) Printf(format string, args ...interface{}) {
	if l.output() {
		l.printChan <- l.printf(format, args...)
	}

}

func (l *InternalLogger) output() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.on
}

func (l *InternalLogger) serialPrinter() {
	for {
		select {
		case msg, ok := <-l.printChan:
			if !ok {
				l.Done <- true
				return
			}
			fmt.Print(msg)
		}
	}
}

// Toggle will turn the logger on or off. When off logs will not be produced.
func (l *InternalLogger) Toggle(enabled bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.on = enabled
}

// Drain will flush the last of the messages and will return a channel that will
// signal when all logs that have been stored have been printed.
func (l *InternalLogger) Drain() chan interface{} {
	l.Toggle(false)
	close(l.printChan)
	return l.Done
}

// NewLogger will return a basic logger type.
func NewLogger() *InternalLogger {
	p := func(s string) string { return fmt.Sprintln(s) }
	pf := func(format string, args ...interface{}) string { return fmt.Sprintf(format, args...) }
	return newLogger(p, pf)
}

// FancyLogger will let you define your own functions for logging. Provided that your functions
// match the signiture of the functions.
func FancyLogger(p func(string) string, pf func(string, ...interface{}) string) *InternalLogger {
	return newLogger(p, pf)
}

func newLogger(p func(string) string, pf func(string, ...interface{}) string) *InternalLogger {
	il := &InternalLogger{
		printChan: make(chan string, 10000),
		Done:      make(chan interface{}, 1),
		on:        true,
		print:     p,
		printf:    pf,
	}
	go il.serialPrinter()
	return il
}
