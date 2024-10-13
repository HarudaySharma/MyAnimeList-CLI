package ui

import (
	"fmt"
	"strings"

	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	c "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui/components"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type AnimeDetailsUI struct {
	Details      *types.NativeAnimeDetails
	DetailFields *[]es.AnimeDetailField
}

func (ui *AnimeDetailsUI) CreateTitle() *tview.TextView {
	alternativeTitles := strings.Builder{}
	alternativeTitles.WriteString("ENG:\t")
	alternativeTitles.WriteString(fmt.Sprintln(ui.Details.AlternativeTitles.EN))
	alternativeTitles.WriteString("JP:\t")
	alternativeTitles.WriteString(fmt.Sprintln(ui.Details.AlternativeTitles.JA))

	titleBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Title",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorLightCyan,
		Text:       alternativeTitles.String(),
		TextAlign:  tview.AlignLeft,
	})

	return titleBox
}

func (ui *AnimeDetailsUI) CreateSynopsis() *tview.TextView {
	synopsisBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Synopsis",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorMistyRose,
		Text:       ui.Details.Synopsis,
		TextAlign:  tview.AlignLeft,
	})

	return synopsisBox
}

func (ui *AnimeDetailsUI) CreateGenres() *tview.TextView {
	genres := strings.Builder{}
	for i, genre := range ui.Details.Genres {
		if i != 0 {
			genres.WriteString(", ")
		}
		genres.WriteString("`" + genre.Name + "`")
	}

	genresBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Genres",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorYellowGreen,
		Text:       genres.String(),
		TextAlign:  tview.AlignLeft,
	})

	return genresBox
}

func (ui *AnimeDetailsUI) CreateStudios() *tview.TextView {
	studios := strings.Builder{}
	for i, studio := range ui.Details.Studios {
		if i != 0 {
			studios.WriteString(", ")
		}
		studios.WriteString("`" + studio.Name + "`")
	}

	studioBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Studios",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       studios.String(),
		TextAlign:  tview.AlignLeft,
	})

	return studioBox
}

func (ui *AnimeDetailsUI) CreateAvgEpDuration() *tview.TextView {
	durationBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Avg. Episode Duration",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       string(ui.Details.AverageEpisodeDuration),
		TextAlign:  tview.AlignLeft,
	})

	return durationBox
}

func (ui *AnimeDetailsUI) CreateBackground() *tview.TextView {
	backgroundBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Background",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       ui.Details.Background,
		TextAlign:  tview.AlignLeft,
	})

	return backgroundBox
}

func (ui *AnimeDetailsUI) CreateBroadcast() *tview.TextView {
	broadcast := strings.Builder{}
    dayOfTheWeek := strings.ReplaceAll(ui.Details.Broadcast.DayOfTheWeek, "\n", "")
    airingTime := strings.ReplaceAll(ui.Details.Broadcast.StartTime, "\n", "")

	broadcast.WriteString("Every " + dayOfTheWeek)
	broadcast.WriteString(" : " + airingTime)

    fmt.Println(broadcast.String())
	boradcastBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Broadcast",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       broadcast.String(),
		TextAlign:  tview.AlignLeft,
	})
    boradcastBox.SetSize(5, 10)

	return boradcastBox
}

/*
	CreatedAt              time.Time              `json:"created_at"`
	EndDate                string                 `json:"end_date"`
	ID                     int                    `json:"id"`
	MainPicture            Picture                `json:"main_picture"`
	Mean                   float64                `json:"mean"`
	MediaType              string                 `json:"media_type"`
	NSFW                   string                 `json:"nsfw"`
	NumEpisodes            int                    `json:"num_episodes"`
	NumListUsers           int                    `json:"num_list_users"`
	NumScoringUsers        int                    `json:"num_scoring_users"`
	Pictures               []Picture              `json:"pictures"`
	Popularity             int                    `json:"popularity"`
	Rank                   int                    `json:"rank"`
	Rating                 string                 `json:"rating"`
	Recommendations        []NativeRecommendation `json:"recommendations"`
	RelatedAnime           []NativeRelatedAnime   `json:"related_anime"`
	Source                 string                 `json:"source"`
	StartDate              string                 `json:"start_date"`
	StartSeason            struct {
		Season string `json:"season"`
		Year   int    `json:"year"`
	} `json:"start_season"`
	Statistics Statistics `json:"statistics"`
	Status     string     `json:"status"`
	UpdatedAt  time.Time  `json:"updated_at"` */

// TODO: work on this
func (ui *AnimeDetailsUI) CreateAdditionalInfo() *tview.Flex {
	// aggregrate of all the user defined anime detail fields
	additionalInfobox := tview.NewFlex().SetDirection(tview.FlexRow)
	additionalInfobox.SetTitle("Additional Information")
	// All the other detail fields will be added dynamically
	fmt.Print(ui.DetailFields)
	for _, field := range *(ui.DetailFields) {
		// all the default details should be skipped
		if isDefault, _ := (*e.DefaultDetailFieldsMap())[field]; isDefault == true {
			continue
		}

		box := c.NewTextView(c.NewTextViewParams{
			Title:      string(field),
			TitleAlign: tview.AlignCenter,
			TitleColor: tcell.ColorMistyRose,
			Text:       "", // TODO: need to create handler for each possible field
			TextAlign:  tview.AlignCenter,
		})
		additionalInfobox.AddItem(box, 0, 1, false)
	}

	return additionalInfobox
}

func (ui *AnimeDetailsUI) CreateLayout() *tview.Flex {
	layout := tview.NewFlex().
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(ui.CreateTitle(), 5, 1, false).
				AddItem(ui.CreateSynopsis(), 0, 3, false).
				AddItem(
					tview.NewFlex().SetDirection(tview.FlexColumn).
						AddItem(ui.CreateGenres(), 0, 1, false).
						AddItem(ui.CreateStudios(), 0, 1, false).
						AddItem(ui.CreateBroadcast(), 0, 2, false), 0, 1, false), 0, 1, false).
		AddItem(ui.CreateAdditionalInfo(), 0, 1, false)

	layout.SetTitle("Anime Details")

	return layout
}
