package ui

import (
	// "bytes"
	// "encoding/base64"
	"fmt"
	// "image/jpeg"
	"slices"
	"strconv"
	"strings"

	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	c "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui/components"
	u "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
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

func (ui *AnimeDetailsUI) CreateStatus() *tview.TextView {
	statusBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Status",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorIndianRed,
		Text:       ui.Details.Status,
		TextAlign:  tview.AlignCenter,
	})

	return statusBox
}

func (ui *AnimeDetailsUI) CreateNumEpisodes() *tview.TextView {
	text := strings.Builder{}
	text.WriteString("[" + tcell.ColorLightSkyBlue.String() + "]")

	episodes := ui.Details.NumEpisodes
	if episodes == 0 {
		text.WriteString("Unknown")
	} else {
		text.WriteString(strconv.Itoa(episodes))
	}
	text.WriteString(" eps")
	text.WriteString("[-]")

	episodesBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Episodes",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorCadetBlue,
		Text:       text.String(),
		TextAlign:  tview.AlignCenter,
	})

	return episodesBox
}

func (ui *AnimeDetailsUI) CreateAverageEpisodeDuration() *tview.TextView {
	duration := ui.Details.AverageEpisodeDuration / 60 // in minutes
	durationStr := strings.Builder{}
	durationStr.WriteString("[" + tcell.ColorLightSkyBlue.String() + "]")
	durationStr.WriteString(strconv.FormatInt(duration, 10))
	durationStr.WriteString(" min")
	durationStr.WriteString("[-]")

	durationBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Avg. Episode Duration",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorLightBlue,
		Text:       durationStr.String(),
		TextAlign:  tview.AlignCenter,
	})

	return durationBox

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

func (ui *AnimeDetailsUI) CreateBackground() (*tview.TextView, int, int) {
	const maxRowLen = 56
	text := strings.Builder{}

	for i, ch := range ui.Details.Background {
		text.WriteRune(ch)
		if (i+1)%maxRowLen == 0 {
			text.WriteString("\n")
		}
	}

	backgroundBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Background",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignLeft,
	})

	w := u.CalMaxWidth(text.String()) + 3
	h := u.CalMaxHeight(text.String()) + 2

	return backgroundBox, w, h
}

func (ui *AnimeDetailsUI) CreateStartSeason() (*tview.TextView, int, int) {
	s := ui.Details.StartSeason

	text := strings.Builder{}
	text.WriteString(s.Season)
	text.WriteString(", ")
	text.WriteString(strconv.Itoa(s.Year))

	seasonBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Season",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignLeft,
	})

	// as using flex container, so don't really have full control over the item's width nd height
	w := u.CalMaxWidth(text.String()) + 3
	h := u.CalMaxHeight(text.String()) + 2

	return seasonBox, w, h
}

func (ui *AnimeDetailsUI) CreateBroadcast() (*tview.TextView, int, int) {
	broadcast := strings.Builder{}
	dayOfTheWeek := strings.ReplaceAll(ui.Details.Broadcast.DayOfTheWeek, "\n", "")
	airingTime := strings.ReplaceAll(ui.Details.Broadcast.StartTime, "\n", "")

	broadcast.WriteString("Every " + strings.ToUpper(dayOfTheWeek))
	broadcast.WriteString(" [ " + airingTime + " ]")

	//fmt.Println(broadcast.String())
	boradcastBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Broadcast",
		TitleAlign: tview.AlignLeft,
		TitleColor: tcell.ColorGreenYellow,
		Text:       broadcast.String(),
		TextAlign:  tview.AlignCenter,
	})
	boradcastBox.SetSize(5, 40)

	w := u.CalMaxWidth(broadcast.String()) + 3
	h := u.CalMaxHeight(broadcast.String()) + 2

	return boradcastBox, w, h
}

func (ui *AnimeDetailsUI) CreateMediaType() (*tview.TextView, int, int) {
	mediatype := c.NewTextView(c.NewTextViewParams{
		Title:      "Media-Type",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       ui.Details.MediaType,
		TextAlign:  tview.AlignCenter,
	})

	w := u.CalMaxWidth(ui.Details.MediaType) + 16
	h := u.CalMaxHeight(ui.Details.MediaType) + 2

	return mediatype, w, h
}

/*will show both NumListUsers & NumScoringUsers
 */
func (ui *AnimeDetailsUI) CreateUsersCount() (*tview.TextView, int, int) {
	listUsers := ui.Details.NumListUsers
	scoringUsers := ui.Details.NumScoringUsers

	text := strings.Builder{}
	text.WriteString("Total Users: ")
	text.WriteString("[red]")
	text.WriteString(u.FormatNumberWithSeparator(listUsers, ","))
	text.WriteString("[-]")
	text.WriteString("\n")
	text.WriteString("Total Scoring Users: ")
	text.WriteString("[red]")
	text.WriteString(u.FormatNumberWithSeparator(scoringUsers, ","))
	text.WriteString("[-]")

	usersCountBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Users Count",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignCenter,
	})

	usersCountBox.SetDynamicColors(true)

	w := u.CalMaxWidth(text.String()) + 3
	h := u.CalMaxHeight(text.String()) + 2

	return usersCountBox, w, h
}

func (ui *AnimeDetailsUI) CreateRating() (*tview.TextView, int, int) {
	rating := ui.Details.Rating

	text := strings.Builder{}
	text.WriteString(enums.RatingMap()[rating])

	ratingBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Rating",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignCenter,
	})

	w := u.CalMaxWidth(text.String()) + 3
	h := u.CalMaxHeight(text.String()) + 2

	return ratingBox, w, h
}

func (ui *AnimeDetailsUI) CreateRank() (*tview.TextView, int, int) {
	rank := ui.Details.Rank

	text := strings.Builder{}
	text.WriteString(u.FormatNumberWithSeparator(int64(rank), ","))

	rankBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Rank",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignCenter,
	})

	w := u.CalMaxWidth(text.String()) + 8
	h := u.CalMaxHeight(text.String()) + 2

	return rankBox, w, h
}

func (ui *AnimeDetailsUI) CreatePopularity() (*tview.TextView, int, int) {
	popularity := ui.Details.Popularity

	text := strings.Builder{}
	text.WriteString(u.FormatNumberWithSeparator(int64(popularity), ","))

	popularityBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Popularity",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignCenter,
	})

	w := u.CalMaxWidth(text.String()) + 8
	h := u.CalMaxHeight(text.String()) + 2

	return popularityBox, w, h
}

func (ui *AnimeDetailsUI) CreateStatistics() (*tview.Flex, int, int) {
	stats := ui.Details.Statistics

	text := strings.Builder{}
	text.WriteString("Watching: ")
	text.WriteString("[red]")

	// type assertion statistics sub fields
	text.WriteString(u.FormatNumberInterfaceWithSeparator(stats.Status.Watching, ","))
	text.WriteString("[-]")
	text.WriteString("\n")
	text.WriteString("Completed: ")
	text.WriteString("[red]")
	text.WriteString(u.FormatNumberInterfaceWithSeparator(stats.Status.Completed, ","))
	text.WriteString("[-]")
	text.WriteString("\n")
	text.WriteString("On Hold: ")
	text.WriteString("[red]")
	text.WriteString(u.FormatNumberInterfaceWithSeparator(stats.Status.OnHold, ","))
	text.WriteString("[-]")
	text.WriteString("\n")
	text.WriteString("Dropped: ")
	text.WriteString("[red]")
	text.WriteString(u.FormatNumberInterfaceWithSeparator(stats.Status.Dropped, ","))
	text.WriteString("[-]")
	text.WriteString("\n")
	text.WriteString("Plan to Watch: ")
	text.WriteString("[red]")
	text.WriteString(u.FormatNumberInterfaceWithSeparator(stats.Status.PlanToWatch, ","))
	text.WriteString("[-]")

	statusBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Status",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       text.String(),
		TextAlign:  tview.AlignCenter,
	})
	statusBox.SetDynamicColors(true)

	totalUsersBox := c.NewTextView(c.NewTextViewParams{
		Title:      "Total Users",
		TitleAlign: tview.AlignCenter,
		TitleColor: tcell.ColorGreenYellow,
		Text:       "Total Users: " + "[red]" + u.FormatNumberWithSeparator(stats.NumListUsers, ",") + "[-]",
		TextAlign:  tview.AlignCenter,
	})
	totalUsersBox.SetBorder(false)

	statisticsBox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(totalUsersBox, 1, 1, false).
		AddItem(statusBox, u.CalMaxHeight(text.String())+2, 1, false)

	statisticsBox.SetTitle("Statistics")
	statisticsBox.SetBorder(true).SetBorderPadding(1, 0, 0, 0)

	w := u.CalMaxWidth(text.String()) + 3
	h := u.CalMaxHeight(text.String()) + 6

	return statisticsBox, w, h
}

// TODO: work on this
func (ui *AnimeDetailsUI) CreateAdditionalInfo() *tview.Flex {
	// aggregrate of all the user defined anime detail fields
	additionalInfobox := tview.NewFlex().SetDirection(tview.FlexRow)

	additionalInfobox.SetBorder(true).SetTitle("Additional Information")

	if contain := slices.Contains(*ui.DetailFields, es.MainPicture); contain {
		if ui.Details.MainPicture.Large != "" || ui.Details.MainPicture.Medium != "" {
			additionalInfobox.AddItem(ui.CreateImage(), 0, 1, false)
		}
	}

	newRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	nextRow := 0.0
	rowW := 1
	rowH := 1

	for _, field := range *(ui.DetailFields) {
		// all the default details should be skipped
		if isDefault, _ := (*e.DefaultDetailFieldsMap())[field]; isDefault == true {
			continue
		}

		var textView *tview.TextView
		var w int
		var h int
		var flexBox *tview.Flex

		switch field {
		case es.Background:
			textView, w, h = ui.CreateBackground()
			nextRow += 3
		case es.StartSeason:
			textView, w, h = ui.CreateStartSeason()
			nextRow += 1
		case es.Broadcast:
			textView, w, h = ui.CreateBroadcast()
			nextRow += 1
		case es.MediaType:
			textView, w, h = ui.CreateMediaType()
			nextRow += 1
			// NOTE: this is dumb
		case es.NumListUsers:
			if ui.Details.NumListUsers != 0 && ui.Details.NumScoringUsers != 0 {
				textView, w, h = ui.CreateUsersCount()
				nextRow += 1.5
			}
		case es.Rating:
			textView, w, h = ui.CreateRating()
			nextRow += 3
		case es.Rank:
			textView, w, h = ui.CreateRank()
			nextRow += 1
		case es.Popularity:
			textView, w, h = ui.CreatePopularity()
			nextRow += 1
		case es.Statistics:
			flexBox, w, h = ui.CreateStatistics()
			nextRow += 3
		}

		if flexBox != nil {
			newRow.AddItem(flexBox, min(w, 100), 1, false)
		}
		if textView != nil {
			newRow.AddItem(textView, min(w, 100), 1, false)
		}
		rowH = max(rowH, h)
		rowW = max(rowW, min(w, 100))

		if nextRow >= 5 {
			additionalInfobox.AddItem(newRow, rowH, 1, false)
			newRow = tview.NewFlex().SetDirection(tview.FlexColumn)
			nextRow = 0.0
			rowH = 0
		}
	}

	if nextRow != 0.0 {
		additionalInfobox.AddItem(newRow, rowH, 1, false)
	}

	return additionalInfobox
}

func (ui *AnimeDetailsUI) CreateImage() *tview.Image {

	if contain := slices.Contains(*ui.DetailFields, es.MainPicture); !contain {
		return tview.NewImage().SetLabel("No-Image")
	}

	// prefers large picture size
	image_link := ui.Details.MainPicture.Large
	if image_link == "" {
		image_link = ui.Details.MainPicture.Medium
	}

	image := tview.NewImage()
	photo, mimetype := u.FetchImage(image_link)
	if mimetype != "jpeg" && mimetype != "jpg" && mimetype != "png" {
		return image.SetLabel("Unsupported Image Format")
	}

	image.SetImage(photo).
		SetColors(tview.TrueColor).
		SetDithering(tview.DitheringFloydSteinberg)

	return image
}

func (ui *AnimeDetailsUI) CreateLayout() *tview.Flex {
	layout := tview.NewFlex().
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(ui.CreateTitle(), 5, 1, false).
				AddItem(
					tview.NewFlex().SetDirection(tview.FlexColumn).
						AddItem(ui.CreateStatus(), 0, 1, false).
						AddItem(ui.CreateNumEpisodes(), 0, 1, false).
						AddItem(ui.CreateAverageEpisodeDuration(), 0, 1, false), 3, 1, false).
				AddItem(ui.CreateGenres(), 3, 1, false).
				AddItem(ui.CreateSynopsis(), 0, 3, false).
				AddItem(
					tview.NewFlex().SetDirection(tview.FlexColumn).
						AddItem(ui.CreateStudios(), 0, 1, false), 3, 1, false), 0, 1, false)

		// Additional Field Check
	defaultFieldCount := 0
	for _, field := range *(ui.DetailFields) {
		if isDefault, _ := (*e.DefaultDetailFieldsMap())[field]; isDefault == true {
			defaultFieldCount++
		}
	}

	if defaultFieldCount == len(*ui.DetailFields) {
		// no Additional Fields Present
		return layout
	}

	layout.AddItem(ui.CreateAdditionalInfo(), 0, 1, false)

	//layout.SetBorder(true).SetTitle("Anime Details")

	return layout
}

/* TODO: Pending Constructors

	ID                     int                    `json:"id"`
	Source                 string                 `json:"source"`

	CreatedAt              time.Time              `json:"created_at"`
    UpdatedAt              time.Time              `json:"updated_at"`

	StartDate              string                 `json:"start_date"`
	EndDate                string                 `json:"end_date"`

	Pictures               []Picture              `json:"pictures"`

	Mean                   float64                `json:"mean"`
	NSFW                   string                 `json:"nsfw"`

	Recommendations        []NativeRecommendation `json:"recommendations"`
	RelatedAnime           []NativeRelatedAnime   `json:"related_anime"`

*/
