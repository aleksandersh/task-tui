package ui

import (
	"github.com/aleksandersh/task-tui/domain"
	"github.com/rivo/tview"
)

type Controller interface {
	StartUi()
	ShowTasks()
	ShowTaskSummary(task *domain.Task)
	ShowHelp()
	Focus(view tview.Primitive)
	PostDraw(f func())
	Back()
	Close()
}
