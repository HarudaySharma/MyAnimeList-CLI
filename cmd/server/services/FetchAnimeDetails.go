package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

func FetchAnimeDetails(animeId string, fields []e.AnimeDetailField) *t.NativeAnimeDetails {
	client := http.Client{}

	fieldsStr := u.ConvertToCommaSeperatedString(fields)
	url := fmt.Sprintf("%s/anime/%s?fields=%s",
		config.C.MalApiUrl,
		animeId,
		fieldsStr,
	)
	log.Println(url)
	req := u.CreateHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var data t.MALAnimeDetails
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatalf("ERROR in FetchAnimeDetails decoding json body \n %v", err)
	}

    defer res.Body.Close()

	return convertToNativeAnimeDetailsType(&data)
}

func convertToNativeAnimeDetailsType(data *t.MALAnimeDetails) *t.NativeAnimeDetails {
	nativeDetails := t.NativeAnimeDetails{
		ID:                     data.ID,
		AlternativeTitles:      data.AlternativeTitles,
		AverageEpisodeDuration: data.AverageEpisodeDuration,
		Background:             data.Background,
		Broadcast:              data.Broadcast,
		CreatedAt:              data.CreatedAt,
		EndDate:                data.EndDate,
		Genres:                 data.Genres,
		MainPicture:            data.MainPicture,
		Mean:                   data.Mean,
		MediaType:              data.MediaType,
		NSFW:                   data.NSFW,
		NumEpisodes:            data.NumEpisodes,
		NumListUsers:           data.NumListUsers,
		NumScoringUsers:        data.NumScoringUsers,
		Pictures:               data.Pictures,
		Popularity:             data.Popularity,
		Rank:                   data.Rank,
		Rating:                 data.Rating,
		Source:                 data.Source,
		StartDate:              data.StartDate,
		StartSeason:            data.StartSeason,
		Statistics:             data.Statistics,
		Status:                 data.Status,
		Studios:                data.Studios,
		Synopsis:               data.Synopsis,
		Title:                  data.Title,
		UpdatedAt:              data.UpdatedAt,
	}

	for _, d := range data.RelatedAnime {
		nra := t.NativeRelatedAnime{}
		nra.Node.ID = d.Node.ID
		nra.Node.Title = d.Node.Title
		nativeDetails.RelatedAnime = append(nativeDetails.RelatedAnime, nra)
	}

	for _, d := range data.Recommendations {
		nrm := t.NativeRecommendation{}
		nrm.Node.ID = d.Node.ID
		nrm.Node.Title = d.Node.Title
		nativeDetails.Recommendations = append(nativeDetails.Recommendations, nrm)
	}

	return &nativeDetails
}
