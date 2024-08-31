package help

import (
	"context"

	"github.com/aleksandersh/task-tui/app/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	title       = " Help "
	status      = " Press [yellow]Esc[white] to go back"
	description = ` Press [yellow]Enter[white] to execute the selected task
 Press [yellow]/[white] to enter the filtering mode
 Press [yellow]s[white] to show the task summary

 Press [yellow]Ctrl+C[white] to exit`
)

func New(ctx context.Context, controller ui.Controller) *tview.Grid {
	var textView = tview.NewTextView().SetText(description)
	textView.SetDynamicColors(true).
		SetRegions(true).
		SetBorder(true).
		SetTitle(title)

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
