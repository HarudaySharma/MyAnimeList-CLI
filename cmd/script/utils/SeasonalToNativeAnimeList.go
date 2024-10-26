package utils

import "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"


func SeasonalToNativeAnimeList(seasonAnimeList *types.NativeSeasonalAnime) *types.NativeAnimeList {
    return &types.NativeAnimeList {
        Data: seasonAnimeList.Data,
        Paging: seasonAnimeList.Paging,
    }
}
