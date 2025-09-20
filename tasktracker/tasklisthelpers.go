package tasktracker

import (
	"fmt"
	"strings"
)

func NewTaskList(name string) *TaskList {
	return &TaskList{
		nextID: 1,
		Name:   name,
	}
}

func (tl *TaskList) AddTask(t *Task) {
	t.ID = tl.nextID
	tl.nextID++
	tl.Tasks = append(tl.Tasks, t)
}

func (tl TaskList) Find(taskID int) *Task {
	for _, t := range tl.Tasks {
		if t.ID == taskID {
			return t
		}
	}

	return nil
}

func (tl *TaskList) Delete(taskID int) *Task {
	var deletedTask *Task
	n := 0
	for _, t := range tl.Tasks {
		if t.ID == taskID {
			deletedTask = t
			continue
		}
		tl.Tasks[n] = t
		n++
	}
	tl.Tasks = tl.Tasks[:n]
	return deletedTask
}

func (tl TaskList) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Task List: %s\n", tl.Name)
	for _, t := range tl.Tasks {
		fmt.Fprintln(&b, t)
	}
	return b.String()
}
