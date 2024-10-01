package main

import ("time")


type Note struct{
	title string
	text string
}

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

type ToDoList struct{
	name string
	tasks []Task
	notes []Note
}

func (to ToDoList) TasksCount() int{
	return len(to.tasks)
}

func (to ToDoList) NotesCount() int{
	return len(to.notes)
}


func (t ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, task := range t.tasks {
		if task.IsTopPriority() {
			count++
		}
	}
	return count
}

func (t ToDoList) CountOverdueTasks() int {
	count := 0
	for _, task := range t.tasks {
		if task.IsOverdue() {
			count++
		}
	}
	return count
}