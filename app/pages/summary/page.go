package summary

import (
	"context"
	"strings"

	"github.com/aleksandersh/task-tui/app/ui"
	"github.com/aleksandersh/task-tui/domain"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	title  = " Task summary "
	status = " Press Esc to go back"
)

func New(ctx context.Context, controller ui.Controller, task *domain.Task) *tview.Grid {
	var textView = tview.NewTextView().SetText(getSummary(task))
	textView.SetBorder(true).SetTitle(title)

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			controller.Back()
			return nil
		}

		return event
	})

	statusView := ui.CreateStatusTextView(status)

	return ui.CreateContainerGrid(textView, statusView)
}

func getSummary(task *domain.Task) string {
	summary := task.Name
	if isNotBlank(task.Summary) {
		summary = summary + "\n\n" + task.Summary
	} else if isNotBlank(task.Description) {
		summary = summary + "\n\n" + task.Description
	}
	return summary
}

func isNotBlank(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}
