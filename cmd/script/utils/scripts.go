package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/colors"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

func fzfPreview() string {
	return `
        # Dependencies:
        # - https://github.com/hpjansson/chafa
        # - https://iterm2.com/utilities/imgcat
        #
        fzf-preview() {
        	file=${1/#\~\//$HOME/}
        	dim=${FZF_PREVIEW_COLUMNS}x${FZF_PREVIEW_LINES}
        	if [[ $dim = x ]]; then
        		dim=$(stty size </dev/tty | awk '{print $2 "x" $1}')
        	elif ! [[ $KITTY_WINDOW_ID ]] && ((FZF_PREVIEW_TOP + FZF_PREVIEW_LINES == $(stty size </dev/tty | awk '{print $1}'))); then
        		# Avoid scrolling issue when the Sixel image touches the bottom of the screen
        		# * https://github.com/junegunn/fzf/issues/2544
        		dim=${FZF_PREVIEW_COLUMNS}x$((FZF_PREVIEW_LINES - 1))
        	fi

        	# 1. Use kitty icat on kitty terminal
        	if [[ $KITTY_WINDOW_ID ]]; then
        		# 1. 'memory' is the fastest option but if you want the image to be scrollable,
        		#    you have to use 'stream'.
        		#
        		# 2. The last line of the output is the ANSI reset code without newline.
        		#    This confuses fzf and makes it render scroll offset indicator.
        		#    So we remove the last line and append the reset code to its previous line.
        		kitty icat --clear --transfer-mode=memory --unicode-placeholder --stdin=no --place="$dim@0x0" "$file" | sed '$d' | sed $'$s/$/\e[m/'

            # 2. Use chafa with Sixel output
        	elif command -v chafa >/dev/null; then
                case "$(uname -a)" in
                    # termux does not support sixel graphics
                    # and produces weird output
                    *ndroid*) chafa -s "$dim" "$file";;
                    *) chafa -f sixel -s "$dim" "$file";;
                esac
        		# Add a new line character so that fzf can display multiple images in the preview window
        		echo

            # 3. If chafa is not found but imgcat is available, use it on iTerm2
        	elif command -v imgcat >/dev/null; then
        		# NOTE: We should use https://iterm2.com/utilities/it2check to check if the
        		# user is running iTerm2. But for the sake of simplicity, we just assume
        		# that's the case here.
        		imgcat -W "${dim%%x*}" -H "${dim##*x}" "$file"

            # 4. Cannot find any suitable method to preview the image
        	else
                echo install chafa or imgcat or install kitty terminal so you can enjoy image previews
        	fi
        }
    `
}

func SavePreviewData(filePath string, node types.AnimeListDataNode) error {
	titleJP := "-"
	titleEN := "-"
	altTitlesInter, _ := node.CustomFields[string(enums.AlternativeTitles)]
	if altTitlesMap, ok := altTitlesInter.(map[string]interface{}); ok {
		if en, ok := altTitlesMap["en"].(string); ok {
			if len(en) == 0 {
				titleEN = node.Title
			} else {
				titleEN = en
			}
		}
		if jp, ok := altTitlesMap["ja"].(string); ok {
			titleJP = jp
		}
	}
	titleJP = strings.ReplaceAll(titleJP, "\"", "'")
	titleEN = strings.ReplaceAll(titleEN, "\"", "'")

	genresB := strings.Builder{}
	genresInter, _ := node.CustomFields[string(enums.Genres)]
	if genreArr, ok := genresInter.([]interface{}); ok {
		for _, genreInterface := range genreArr {
			if genreMap, ok := genreInterface.(map[string]interface{}); ok {
				if genre, ok := genreMap["name"].(string); ok {
					if genresB.Len() != 0 {
						genresB.WriteString(", ")
					}
					genre = strings.ReplaceAll(genre, "\"", "'")
					genresB.WriteString(genre)
				}
			}
		}
	}

	statusInter, ok := node.CustomFields[string(enums.Status)]
	animeStatusStr, ok := statusInter.(string)
	if !ok {
		animeStatusStr = "-"
	}
	animeStatusStr = strings.ReplaceAll(animeStatusStr, "_", " ")

	meanStr := "-"
	meanInter, ok := node.CustomFields[string(enums.Mean)]
	if mean, ok := meanInter.(float64); ok {
		meanStr = fmt.Sprintf("%.2f", mean)
	}

	numEpisodesStr := "-"
	numEpisodesInter, ok := node.CustomFields[string(enums.NumEpisodes)]
	if numEpisodes, ok := numEpisodesInter.(float64); ok {
		numEpisodesStr = fmt.Sprintf("%.f", numEpisodes)
	}

	avgEpDurationStrB := strings.Builder{}
	avgEpDurationStrB.WriteString("-")
	avgEpDurationInter, ok := node.CustomFields[string(enums.AverageEpisodeDuration)]
	if avgEpDuration, ok := avgEpDurationInter.(float64); ok {
		duration := avgEpDuration / 60 // in minutes
		avgEpDurationStrB.Reset()
		avgEpDurationStrB.WriteString(fmt.Sprintf("%.0f", duration))
		avgEpDurationStrB.WriteString(" min")
	}

	startDateInter, ok := node.CustomFields[string(enums.StartDate)]
	startDate, ok := startDateInter.(string)
	if !ok {
		startDate = "-"
	}

	endDateInter, ok := node.CustomFields[string(enums.EndDate)]
	endDate, ok := endDateInter.(string)
	if !ok {
		endDate = "-"
	}

	broadCastInter, ok := node.CustomFields["broadcast"]
	broadcast := types.Broadcast{}
	if broadcastMap, ok := broadCastInter.(map[string]interface{}); ok {
		if day, ok := broadcastMap["day_of_the_week"].(string); ok {
			broadcast.DayOfTheWeek = day
		} else {
			broadcast.DayOfTheWeek = "-"
		}

		if start, ok := broadcastMap["start_time"].(string); ok {
			broadcast.StartTime = start
		} else {
			broadcast.StartTime = "-"
		}
	} else {
		broadcast.DayOfTheWeek = "-"
		broadcast.StartTime = "-"
	}
	dayOfTheWeek := strings.ReplaceAll(broadcast.DayOfTheWeek, "\n", "")
	airingTime := strings.ReplaceAll(broadcast.StartTime, "\n", "")
	broadcastStrB := strings.Builder{}
	broadcastStrB.WriteString(dayOfTheWeek)
	broadcastStrB.WriteString(" [ " + airingTime + " ]")

	script := fmt.Sprintf(`
            if command -v fold > /dev/null 2>&1; then
                wrap() {
                    fold -w "$1"
                }
            else
                wrap() {
                    cat
                }
            fi

            cols=$FZF_PREVIEW_COLUMNS
            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n -e  "─"
                ll=$((ll + 1))
            done
            echo
            echo "%sTitle(en):%s %s" | wrap $(($cols + 10))
            echo "%sTitle(jp):%s %s" | wrap $(($cols + 10))
            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n -e "─"
                ll=$((ll + 1))
            done
            echo
            echo "%sGenres:%s %s" | wrap $(($cols + 10))
            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n -e "─"
                ll=$((ll + 1))
            done
            echo
            echo "%sScore:%s %s ⭐" | wrap $(($cols + 10))
            echo "%sStatus:%s %s" | wrap $(($cols + 10))
            echo "%sAiring:%s %s" | wrap $(($cols + 10))
            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n "─"
                ll=$((ll + 1))
            done
            echo
            echo "%sEpisodes:%s %s, (%s)" | wrap $(($cols + 10))
            echo "%sStart Date:%s %s" | wrap $(($cols + 10))
            echo "%sEnd Date:%s %s" | wrap $(($cols + 10))
            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n "─"
                ll=$((ll + 1))
            done
            echo
            `,
		colors.Red, colors.Reset,
		titleEN,
		colors.Red, colors.Reset,
		titleJP,
		colors.Red, colors.Reset,
		genresB.String(),
		colors.Red, colors.Reset,
		meanStr,
		colors.Red, colors.Reset,
		animeStatusStr,
		colors.Red, colors.Reset,
		broadcastStrB.String(),
		colors.Red, colors.Reset,
		numEpisodesStr,
		avgEpDurationStrB.String(),
		colors.Red, colors.Reset,
		startDate,
		colors.Red, colors.Reset,
		endDate,
	)
	/*
	        #            while [ $ll -le $FZF_PREVIEW_COLUMNS ];do
	        #        #echo -n -e "{get_true_fg("─",*SEPARATOR_COLOR,bold=False)}"
	        #                echo -n -e "─"
	        #                ((ll++))
	        #            done
	        #            echo
	        #            echo "{get_true_fg('Episodes:',*HEADER_COLOR)} {(anime['episodes']) or 'UNKNOWN'}"
	        #            echo "{get_true_fg('Start Date:',*HEADER_COLOR)} {anilist_data_helper.format_anilist_date_object(anime['startDate']).replace('"',SINGLE_QUOTE)}"
	        #            echo "{get_true_fg('End Date:',*HEADER_COLOR)} {anilist_data_helper.format_anilist_date_object(anime['endDate']).replace('"',SINGLE_QUOTE)}"
	        #            ll=2
	        #            while [ $ll -le $FZF_PREVIEW_COLUMNS ];do
	        #        #echo -n -e "{get_true_fg("─",*SEPARATOR_COLOR,bold=False)}"
	        #                echo -n -e "─"
	        #                ((ll++))
	        #            done
	        #            echo
	        #            echo "{get_true_fg('Media List:',*HEADER_COLOR)} {mediaListName.replace('"',SINGLE_QUOTE)}"
	        #            echo "{get_true_fg('Progress:',*HEADER_COLOR)} {progress}"
	        #            ll=2
	        #            while [ $ll -le $FZF_PREVIEW_COLUMNS ];do
	        #        #echo -n -e "{get_true_fg("─",*SEPARATOR_COLOR,bold=False)}"
	        #                echo -n -e "─"
	        #                ((ll++))
	        #            done
		)) */

	dir := filePath[:len(filePath)-len("/"+filePath[strings.LastIndex(filePath, "/")+1:])]

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write([]byte(script))

	return err

}
