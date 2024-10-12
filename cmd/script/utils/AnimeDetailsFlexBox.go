package utils

import (
	"fmt"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func AnimeDetailsFlexBox(animeDetails *types.NativeAnimeDetails, detailFields *[]enums.AnimeDetailField) *tview.Flex {
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

	synopsisBox := NewTextView(NewTextViewParams{
		Title:      "Synopsis",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorMistyRose,
		Text:       animeDetails.Synopsis,
		TextAlign:  tview.AlignLeft,
	})

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

	leftBox := tview.NewFlex()
	// All the other detail fields will be added dynamically
	for _, field := range *detailFields {
		box := NewTextView(NewTextViewParams{
			Title:      "Synopsis",
			TitleAlign: tview.AlignCenter,
			TitleColor: tcell.ColorMistyRose,
			Text:       string(field), // animeDetails[field],
			TextAlign:  tview.AlignLeft,
		})
		leftBox.AddItem(box, 0, 1, false)
	}

	rightBox := NewBox(NewBoxParams{
		Title: "Left (20 cols)",
	})

	flex := tview.NewFlex().
		AddItem(leftBox, 20, 1, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(titleBox, 5, 1, false).
				AddItem(synopsisBox, 0, 3, false).
				AddItem(
					tview.NewFlex().SetDirection(tview.FlexColumn).
						AddItem(genresBox, 0, 2, false).
						AddItem(studioBox, 0, 1, false), 0, 1, false), 0, 1, false).
		AddItem(rightBox, 20, 1, false)

	flex.SetTitle("Anime Details")
	return flex
}
