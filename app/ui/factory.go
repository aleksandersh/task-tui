package ui

import "github.com/rivo/tview"

func CreateStatusTextView(text string) *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetMaxLines(1).
		SetText(text)
}

func CreateContainerGrid(contentView tview.Primitive, statusView *tview.TextView) *tview.Grid {
	return tview.NewGrid().
		SetRows(0, 2).
		AddItem(contentView, 0, 0, 1, 1, 0, 0, true).
		AddItem(statusView, 1, 0, 1, 1, 2, 0, false)
}
