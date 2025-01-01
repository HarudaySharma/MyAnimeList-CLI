package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewApplication(animeDetailsUI *AnimeDetailsUI) *tview.Application {
	formShown := false
	app := tview.NewApplication()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			app.Stop() // Terminate the application
			return nil
		}

		switch event.Rune() {
		case 'q': // Quit the application
			if formShown {
				app.SetRoot(animeDetailsUI.CreateLayout(), true)
				formShown = false
			} else {
				app.Stop()
			}
		case 'f': // toggles a form
			if formShown {
				app.SetRoot(animeDetailsUI.CreateLayout(), true)
				formShown = false
			} else {
				app.SetRoot(animeDetailsUI.UserAnimeStatusForm(app), true)
				formShown = true
			}
		}
		return event
	})

	app.SetRoot(animeDetailsUI.CreateLayout(), true).EnableMouse(true)

	return app
}
