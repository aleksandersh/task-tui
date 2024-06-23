package tasks

import (
	"strings"

	"github.com/aleksandersh/task-tui/domain"
)

type tasksViewList struct {
	items []tasksViewItem
}

type tasksViewItem struct {
	Index int
	Text  string
	Task  domain.Task
}

func createTasksViewList(tasfile *domain.Taskfile) *tasksViewList {
	tasks := make([]tasksViewItem, 0, len(tasfile.Tasks))
	for index, task := range tasfile.Tasks {
		tasks = append(tasks, tasksViewItem{
			Index: index,
			Text:  strings.ToLower(task.Name),
			Task:  task,
		})
	}
	return &tasksViewList{items: tasks}
}
