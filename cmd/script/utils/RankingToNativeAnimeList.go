package utils

import "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"


func RankingToNativeAnimeList(rankingAnimeList *types.NativeAnimeRanking) *types.NativeAnimeList {
    return &types.NativeAnimeList {
        //Data: rankingAnimeList.Data,
        Paging: rankingAnimeList.Paging,
    }
}
