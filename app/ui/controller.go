package ui

import "github.com/rivo/tview"

type Controller interface {
	StartUi()
	ShowTasks()
	Focus(view tview.Primitive)
	PostDraw(f func())
	Close()
}
