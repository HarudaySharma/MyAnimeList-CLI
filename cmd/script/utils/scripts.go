package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
        		kitty icat --clear --transfer-mode=stream --unicode-placeholder --stdin=no --place="$dim@0x0" "$file" | sed '$d' | sed $'$s/$/\e[m/'

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

var userPfpLoc = imageDir + "/user/pfp"
var userDataLoc = dataDir + "/user/details"

func SaveUserPreviewData(userD *types.NativeUserDetails) error {
	// NOTE: all the saving are being done according to the preview script written

	// save preview images to cache
	go func() {
		url := fmt.Sprintf("%v", userD.Picture)
		if err := DownloadImage(url, userPfpLoc); err != nil {
			fmt.Println(err)
			fmt.Println("error dowloading image")
		}
	}()

	// NOTE: Won't be caching the user details as it is subject to frequent change
	dataScript := GenerateUserDataScript(userD)

	filePath := userDataLoc

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

	_, err = out.Write([]byte(dataScript))

	return err
}

func SaveAnimePreviewData(key string, node *types.AnimeListDataNode) error {
	// NOTE: all the saving are being done according to the preview script written

	mainPicture := node.CustomFields["main_picture"]
	if mainPicture, ok := mainPicture.(map[string]interface{}); ok {
		medium, _ := mainPicture["medium"]
		url := fmt.Sprintf("%v", medium)
		fileName := strings.ReplaceAll(key, " ", "")
		fileName = strings.ReplaceAll(fileName, "\t", "")
		fileName += filepath.Ext(url)

		go func() {
			// save preview images to cache
			if err := DownloadImage(url, imageDir+"/"+fileName); err != nil {
				fmt.Println(err)
				fmt.Println("error dowloading image")
			}
		}()
	}

	dataFileName := strings.ReplaceAll(key, " ", "")
	dataFileName = strings.ReplaceAll(dataFileName, "\t", "")
	dataFilePath := dataDir + "/" + dataFileName

	if checkFileExists(dataFilePath) {
		return nil
	}

	dataScript := GenerateAnimeDataPreviewScript(node)

	dir := dataFilePath[:len(dataFilePath)-len("/"+dataFilePath[strings.LastIndex(dataFilePath, "/")+1:])]

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	out, err := os.Create(dataFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write([]byte(dataScript))

	return err
}

func SaveUserAnimePreviewData(key string, node *types.UserAnimeListDataNode) error {
	// NOTE: all the saving are being done according to the preview script written

	mainPicture := node.Node.CustomFields["main_picture"]
	if mainPicture, ok := mainPicture.(map[string]interface{}); ok {
		medium, _ := mainPicture["medium"]
		url := fmt.Sprintf("%v", medium)
		fileName := strings.ReplaceAll(key, " ", "")
		fileName = strings.ReplaceAll(fileName, "\t", "")
		fileName += filepath.Ext(url)

		go func() {
			// save preview images to cache
			if err := DownloadImage(url, imageDir+"/"+fileName); err != nil {
				fmt.Println(err)
				fmt.Println("error dowloading image")
			}
		}()
	}

	dataFileName := strings.ReplaceAll(key, " ", "")
	dataFileName = strings.ReplaceAll(dataFileName, "\t", "")

	animeDataFilePath := dataDir + "/" + dataFileName // NOTE:
	if !checkFileExists(animeDataFilePath) {
		// creating animeData files

		animeDataScript := GenerateAnimeDataPreviewScript(&node.Node)
		/* dataScript := GenerateAnimeDataPreviewScript(node.Node)
		dataScript += GenerateUserListStatusScript(node.AnimeStatus) */

		dir := animeDataFilePath[:len(animeDataFilePath)-len("/"+animeDataFilePath[strings.LastIndex(animeDataFilePath, "/")+1:])]

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directories: %v", err)
		}

		out, err := os.Create(animeDataFilePath)
		if err != nil {
			return err
		}

		defer out.Close()
		_, err = out.Write([]byte(animeDataScript))

	}

	// creating userAnimeData files everytime as this data is prone to change actively

	userAnimeDataFilePath := dataDir + "/user/" + dataFileName // NOTE:

	userAnimeDataScript := GenerateUserListStatusScript(node.AnimeStatus)

	dir := userAnimeDataFilePath[:len(userAnimeDataFilePath)-len("/"+userAnimeDataFilePath[strings.LastIndex(userAnimeDataFilePath, "/")+1:])]

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	out, err := os.Create(userAnimeDataFilePath)
	if err != nil {
		return err
	}
    defer out.Close()

	_, err = out.Write([]byte(userAnimeDataScript))

	return err
}

func GenerateUserDataScript(userD *types.NativeUserDetails) string {
	stats := strings.Builder{}

	stats.WriteString(fmt.Sprintf(`
        echo "%s 📊 Anime Stats 📊 %s"
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "-"
            ll=$((ll + 1))
        done
        echo
        echo "%sTotal Episodes%s: %.0f"
        echo "%sMean Score%s: %.2f"
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "-"
            ll=$((ll + 1))
        done
        echo
        echo "%sTotal Days Watched%s: %.0f"
        echo "%sTotal Days Watching%s: %.0f"
        echo "%sTotal Days Completed%s: %.0f"
        echo "%sTotal Days OnHold%s: %.0f"
        echo "%sTotal Days Dropped%s: %.0f"
        echo "%sTotal Days Count%s: %.0f"
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "-"
            ll=$((ll + 1))
        done
        echo
        echo "%sTotal Anime Completed%s: %.0f"
        echo "%sTotal Anime Watching%s: %.0f"
        echo "%sTotal Anime Plan To Watch%s: %.0f"
        echo "%sTotal Anime OnHold%s: %.0f"
        echo "%sTotal Anime Dropped%s: %.0f"
        echo "%sTotal Anime Rewatched%s: %.0f"
        echo "%sTotal Anime Count%s: %.0f"
        echo
        `,
		colors.Red, colors.Reset,

		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumEpisodes,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.MeanScore,

		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumDaysWatched,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumDaysWatching,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumDaysCompleted,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumDaysOnHold,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumDaysDropped,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumDays,

		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumItemsCompleted,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumItemsWatching,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumItemsPlanToWatch,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumItemsDropped,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumItemsOnHold,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumTimesRewatched,
		colors.Cyan, colors.Reset,
		userD.AnimeStatistics.NumItems,
	))

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
        echo "%sID:%s %d" | wrap $(($cols + 10))
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "─"
            ll=$((ll + 1))
        done
        echo
        echo "%sName:%s %s" | wrap $(($cols + 10))
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "─"
            ll=$((ll + 1))
        done
        echo
        echo "%sJoined At:%s %s" | wrap $(($cols + 10))
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "─"
            ll=$((ll + 1))
        done
        echo
        echo "%sLocation:%s %s" | wrap $(($cols + 10))
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "─"
            ll=$((ll + 1))
        done
        ll=2
        while [ "$ll" -le "$cols" ]; do
            echo -n -e  "─"
            ll=$((ll + 1))
        done
        echo
        %s
    `,

		colors.Red, colors.Reset,
		userD.Id,
		colors.Red, colors.Reset,
		userD.Name,
		colors.Red, colors.Reset,
		userD.JoinedAt,
		colors.Red, colors.Reset,
		userD.Location,
		stats.String(),
	)

	return script
}

func GenerateAnimeDataPreviewScript(node *types.AnimeListDataNode) string {
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

	return script
}

func GenerateUserListStatusScript(userListStatus types.NativeUserAnimeStatus) string {
	if userListStatus == (types.NativeUserAnimeStatus{}) {
		return ""
	}

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

            echo "%sUser List Status%s"

            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n -e  "-"
                ll=$((ll + 1))
            done
            echo
            echo "%sStatus:%s %s" | wrap $(($cols + 10))
            echo "%sScore:%s %d" | wrap $(($cols + 10))
            echo "%sEpisodes Watched:%s %d" | wrap $(($cols + 10))
            ll=2
            while [ "$ll" -le "$cols" ]; do
                echo -n -e  "-"
                ll=$((ll + 1))
            done
            echo
            echo "%sRe-Watching:%s %v" | wrap $(($cols + 10))
            echo "%sLast Updated:%s %s" | wrap $(($cols + 10))
            echo
        `,
		colors.Red, colors.Reset,
		colors.Cyan, colors.Reset,
		userListStatus.Status,
		colors.Cyan, colors.Reset,
		userListStatus.Score,
		colors.Cyan, colors.Reset,
		userListStatus.NumWatchedEpisodes,
		colors.Cyan, colors.Reset,
		userListStatus.IsRewatching,
		colors.Cyan, colors.Reset,
		userListStatus.UpdatedAt,
	)

	return script
}

func GenerateUserPreviewScript() string {

	previewScript := fmt.Sprintf(`
        if [ -s "%s" ]; then
            fzf-preview "%s"
        else
            echo "Loading User Image..."
        fi

        if [ -s "%s" ]; then
            source "%s"
        else
            echo "Loading User Data..."
        fi
    `,
		userPfpLoc,
		userPfpLoc,
		userDataLoc,
		userDataLoc,
	)

	return previewScript
}

func GenerateAnimePreviewScript() string {
	userAnimeDataDir := dataDir + "/user"

	previewScript := fmt.Sprintf(`
            title=$(echo {} | tr -d '[:space:]')
            show_image_previews="%s"
            if [ "${show_image_previews}" = "true" ];then
                if [ -s "%s/${title}.jpg" ]; then
                    fzf-preview "%s/${title}.jpg"
                elif [ -s "%s/${title}.png" ]; then
                    fzf-preview "%s/${title}.png"
                elif [ -s "%s/${title}.webp" ]; then
                    fzf-preview "%s/${title}.webp"
                else
                    echo Loading Image...
                fi
            fi

            #animeData
            if [ -s "%s/${title}" ]; then
                source "%s/${title}"
            else
                echo Loading Data...
            fi
            #userAnimeData
            if [ -s "%s/${title}" ]; then
                source "%s/${title}"
            fi
    `,
		"true",
		imageDir, imageDir,
		imageDir, imageDir,
		imageDir, imageDir,
		dataDir, dataDir,
		userAnimeDataDir, userAnimeDataDir,
	)

	return previewScript
}
