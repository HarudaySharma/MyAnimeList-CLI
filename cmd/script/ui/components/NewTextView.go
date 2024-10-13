package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NewTextViewParams struct {
	Title      string
	TitleAlign int
	TitleColor tcell.Color
	Text       string
	TextAlign  int
}

func NewTextView(p NewTextViewParams) *tview.TextView {
	textView := tview.NewTextView().
		SetText(p.Text).
		SetTextAlign(p.TextAlign).
		SetWrap(true).
		SetScrollable(true).
		SetDynamicColors(true)

	textView.SetBackgroundColor(tcell.ColorDefault).
		SetTitle(p.Title).
		SetTitleAlign(p.TitleAlign).
		SetTitleColor(p.TitleColor).
		SetBorder(true)

	return textView

}
