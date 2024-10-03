package main

import (
	"fmt"
)

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Logger interface {
	Log()
}

type Log struct {
	Level LogLevel
}

func (l Log) Log(txt string) {
	if l.Level == "Error" {
		fmt.Printf("ERROR: %s\n", txt)
	} else if l.Level == "Info" {
		fmt.Printf("INFO: %s\n", txt)
	}
}