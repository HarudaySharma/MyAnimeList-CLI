package utils

import "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"

func UserToNativeAnimeList(userAnimeList *types.NativeUserAnimeList) *types.NativeAnimeList {
	data := make([]types.AnimeListDataNode, 0)
	for _, node := range userAnimeList.Data {
		data = append(data, node.Node)
	}

	return &types.NativeAnimeList{
		Data:   data,
		Paging: userAnimeList.Paging,
	}
}
