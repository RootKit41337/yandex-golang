package main

import (
	"time"
)

type Task struct{
	summary string
	description string
	deadline time.Time 
	priority int
}

func (t Task) IsOverdue() bool {
    return t.deadline.Before(time.Now())
}

func (t Task) IsTopPriority() bool {
	return t.priority >= 4 && t.priority <= 5
}