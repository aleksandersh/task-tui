package tasks

import (
	"context"

	"github.com/aleksandersh/task-tui/app/ui"
	"github.com/aleksandersh/task-tui/domain"
	"github.com/aleksandersh/task-tui/task"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type view struct {
	tasks  *tview.List
	filter *tview.TextArea
}

func New(ctx context.Context, task *task.Task, config *ui.Config, uiController ui.Controller, taskfile *domain.Taskfile) *tview.Grid {
	tasksView := createTasksView(config)
	filterView := createFilterView()

	v := &view{tasks: tasksView, filter: filterView}
	c := newController(ctx, task, uiController, v, taskfile)
	c.start()

	container := createContainerView(tasksView, filterView)
	startInputHandling(container, c)
	v.startFilterChangesHandling(c)

	return container
}

func startInputHandling(v *tview.Grid, c *controller) {
	filterViewActive := false
	v.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !filterViewActive {
			if event.Key() == tcell.KeyRune && event.Rune() == ui.RuneSlash {
				c.onActivateFilter()
				filterViewActive = true
				return nil
			}
			if event.Key() == tcell.KeyEsc {
				c.onResetFilter()
				return nil
			}
			if event.Key() == tcell.KeyRune && event.Rune() == ui.RuneS {
				c.onClickSummary()
				return nil
			}
		}
		if filterViewActive {
			if event.Key() == tcell.KeyEsc {
				c.onCancelFilter()
				filterViewActive = false
				return nil
			}
			if event.Key() == tcell.KeyEnter {
				c.onFinishFilter()
				filterViewActive = false
				return nil
			}
		}
		return event
	})
}

func (v *view) startFilterChangesHandling(c *controller) {
	v.filter.SetChangedFunc(func() {
		c.onFilterChanged(v.filter.GetText())
	})
}

func createTasksView(config *ui.Config) *tview.List {
	view := tview.NewList()
	view.SetHighlightFullLine(true).
		ShowSecondaryText(config.SecondLineEnabled).
		SetSecondaryTextColor(tcell.Color16).
		SetWrapAround(false).
		SetTitle(" Taskfile ").
		SetBorder(true)

	return view
}

func createFilterView() *tview.TextArea {
	filterView := tview.NewTextArea()
	filterView.SetDisabled(true)
	filterView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune || event.Key() == tcell.KeyBackspace2 {
			return event
		}
		return nil
	})
	return filterView
}

func createContainerView(tasksView *tview.List, filterView *tview.TextArea) *tview.Grid {
	view := tview.NewGrid().
		SetRows(0, 2).
		AddItem(tasksView, 0, 0, 1, 1, 0, 0, true).
		AddItem(filterView, 1, 0, 1, 1, 2, 0, false)
	return view
}
