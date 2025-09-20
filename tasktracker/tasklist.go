package tasktracker

type TaskList struct {
	nextID	int
	Name  	string
	Tasks 	[]*Task
}
