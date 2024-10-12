package utils

import (
	"fmt"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NOTE: refactor it to be more dynamic
func AnimeDetailsFlexBox(animeDetails *types.NativeAnimeDetails, detailFields *[]es.AnimeDetailField) *tview.Flex {

	additionalInfobox := tview.NewFlex().SetDirection(tview.FlexRow)
	additionalInfobox.SetTitle("Additional Information")
	// All the other detail fields will be added dynamically
	fmt.Print(detailFields)
	for _, field := range *detailFields {
		// all the default details should be skipped
		if isDefault, _ := (*enums.DefaultDetailFieldsMap())[field]; isDefault == true {
			continue
		}

		box := NewTextView(NewTextViewParams{
			Title:      string(field),
			TitleAlign: tview.AlignCenter,
			TitleColor: tcell.ColorMistyRose,
            Text:       "",// TODO: need to create handler for each possible field
			TextAlign:  tview.AlignCenter,
		})
		additionalInfobox.AddItem(box, 0, 1, false)
	}

	titleBox := CreateTitle(animeDetails)
	synopsisBox := CreateSynopsis(animeDetails)
	genresBox := CreateGenres(animeDetails)
	studioBox := CreateStudios(animeDetails)

	flex := tview.NewFlex().
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(titleBox, 5, 1, false).
				AddItem(synopsisBox, 0, 3, false).
				AddItem(
					tview.NewFlex().SetDirection(tview.FlexColumn).
						AddItem(genresBox, 0, 2, false).
						AddItem(studioBox, 0, 1, false), 0, 1, false), 0, 1, false).
		AddItem(additionalInfobox, 0, 1, false)

	flex.SetTitle("Anime Details")
	return flex
}
