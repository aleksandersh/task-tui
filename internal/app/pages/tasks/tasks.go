package tasks

import (
	"context"

	"github.com/aleksandersh/task-tui/internal/app/ui"
	"github.com/aleksandersh/task-tui/internal/domain"
	"github.com/aleksandersh/task-tui/internal/task"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type view struct {
	container    *tview.Grid
	tasks        *tview.List
	filter       *tview.TextArea
	status       *tview.TextView
	filterActive bool
}

func New(ctx context.Context, task *task.Task, config *ui.Config, uiController ui.Controller, taskfile *domain.Taskfile) *tview.Grid {
	tasksView := createTasksView(config)
	filterView := createFilterView()
	statusView := ui.CreateStatusTextView(" Press [yellow]h[white] to show the help page")
	container := createContainerView(tasksView, statusView)

	v := &view{container: container, tasks: tasksView, filter: filterView, status: statusView}
	c := newController(ctx, task, uiController, v, taskfile)
	c.start()

	startInputHandling(v, c)
	v.startFilterChangesHandling(c)

	return container
}

func startInputHandling(v *view, c *controller) {
	filterViewVisible := false
	filterViewActive := false
	v.container.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !filterViewActive {
			if event.Key() == tcell.KeyRune && event.Rune() == ui.RuneSlash {
				if !filterViewVisible {
					v.showFilterView()
					filterViewVisible = true
				}
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
			if event.Key() == tcell.KeyRune && event.Rune() == ui.RuneH {
				c.onClickHelp()
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

func (v *view) showFilterView() {
	v.container.RemoveItem(v.filter)
	v.container.RemoveItem(v.status)
	v.container.AddItem(v.filter, 1, 0, 1, 1, 2, 0, false)
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

func createContainerView(tasksView *tview.List, statusView *tview.TextView) *tview.Grid {
	view := tview.NewGrid().
		SetRows(0, 2).
		AddItem(tasksView, 0, 0, 1, 1, 0, 0, true).
		AddItem(statusView, 1, 0, 1, 1, 2, 0, false)
	return view
}
