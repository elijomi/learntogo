package tasktracker

import "fmt"

func NewTask(title string, priority int) *Task {
	return &Task{
		Title:    title,
		Priority: priority,
		Done:     false,
	}
}

func (t Task) String() string {
	x := " "
	if t.Done {
		x = "X"
	}
	return fmt.Sprintf("[%s] #ID:%d \"%s\" (priority=%d)", x, t.ID, t.Title, t.Priority)
}
