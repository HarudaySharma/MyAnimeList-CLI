package utils

import (
	"encoding/json"
	"fmt"
	"learning/server/config"
	"learning/server/types"
	"log"
	"net/http"
)

func FetchAnimeDetails(animeId string, fields []AnimeDetailField) *types.NativeAnimeDetails {
	client := http.Client{}

	fieldsStr := ConvertToCommaSeperatedString(fields)
	url := fmt.Sprintf("%s/anime/%s?fields=%s",
		config.C.MAL_API_URL,
		animeId,
		fieldsStr,
	)
	log.Println(url)
	req := CreateHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var data types.MALAnimeDetails
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatalf("ERROR in FetchAnimeDetails decoding json body \n %v", err)
	}

    defer res.Body.Close()

	return convertToNativeAnimeDetailsType(&data)
}

func convertToNativeAnimeDetailsType(data *types.MALAnimeDetails) *types.NativeAnimeDetails {
	nativeDetails := types.NativeAnimeDetails{
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
		nra := types.NativeRelatedAnime{}
		nra.Node.ID = d.Node.ID
		nra.Node.Title = d.Node.Title
		nativeDetails.RelatedAnime = append(nativeDetails.RelatedAnime, nra)
	}

	for _, d := range data.Recommendations {
		nrm := types.NativeRecommendation{}
		nrm.Node.ID = d.Node.ID
		nrm.Node.Title = d.Node.Title
		nativeDetails.Recommendations = append(nativeDetails.Recommendations, nrm)
	}

	return &nativeDetails
}
