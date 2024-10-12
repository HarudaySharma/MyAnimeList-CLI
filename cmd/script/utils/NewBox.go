package utils

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NewBoxParams struct {
	Title string
}

func NewBox(p NewBoxParams) *tview.Box {
	return tview.NewBox().
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true).
		SetTitle(p.Title)
}
