package utils

import (
	"fmt"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateTitle(animeDetails *types.NativeAnimeDetails) *tview.TextView {
	alternativeTitles := strings.Builder{}
	alternativeTitles.WriteString("ENG:\t")
	alternativeTitles.WriteString(fmt.Sprintln(animeDetails.AlternativeTitles.EN))
	alternativeTitles.WriteString("JP:\t")
	alternativeTitles.WriteString(fmt.Sprintln(animeDetails.AlternativeTitles.JA))

	titleBox := NewTextView(NewTextViewParams{
		Title:      "Title",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorLightCyan,
		Text:       alternativeTitles.String(),
		TextAlign:  tview.AlignLeft,
	})

	return titleBox
}

func CreateSynopsis(animeDetails *types.NativeAnimeDetails) *tview.TextView {
	synopsisBox := NewTextView(NewTextViewParams{
		Title:      "Synopsis",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorMistyRose,
		Text:       animeDetails.Synopsis,
		TextAlign:  tview.AlignLeft,
	})

	return synopsisBox
}

func CreateGenres(animeDetails *types.NativeAnimeDetails) *tview.TextView {
	genres := strings.Builder{}
	for i, genre := range animeDetails.Genres {
		if i != 0 {
			genres.WriteString(", ")
		}
		genres.WriteString("`" + genre.Name + "`")
	}

	genresBox := NewTextView(NewTextViewParams{
		Title:      "Genres",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorYellowGreen,
		Text:       genres.String(),
		TextAlign:  tview.AlignLeft,
	})

	return genresBox
}

func CreateStudios(animeDetails *types.NativeAnimeDetails) *tview.TextView {
	studios := strings.Builder{}
	for i, studio := range animeDetails.Studios {
		if i != 0 {
			studios.WriteString(", ")
		}
		studios.WriteString("`" + studio.Name + "`")
	}

	studioBox := NewTextView(NewTextViewParams{
		Title:      "Studios",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       studios.String(),
		TextAlign:  tview.AlignLeft,
	})

    return studioBox
}
