package utils

import (
	"fmt"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

type UpdateAnimeStatusCacheParams struct {
	ListNode    interface{}
	AnimeStatus *types.NativeUserAnimeStatus
}

func UpdateUserAnimeStatusCache(p UpdateAnimeStatusCacheParams) {
	// TODO: update the cache for anime status only as it is the only part that can be changed by user.
	if node, ok := p.ListNode.(*types.AnimeListDataNode); ok {
		cacheKey, _ := GenerateUserAnimePreviewKeys(node)

		SaveUserAnimePreviewData(cacheKey, &types.UserAnimeListDataNode{
			Node:        *node,
			AnimeStatus: *p.AnimeStatus,
		})
	} else if node, ok := p.ListNode.(*types.AnimeRankingDataNode); ok {
		cacheKey, _ := GenerateRankingAnimePreviewKeys(node)

		SaveUserAnimePreviewData(cacheKey, &types.UserAnimeListDataNode{
			Node:        node.Node,
			AnimeStatus: *p.AnimeStatus,
		})
	} else if node, ok := p.ListNode.(*types.UserAnimeListDataNode); ok {
		cacheKey, _ := GenerateUserAnimePreviewKeys(&node.Node)

		if p.AnimeStatus != nil {
			node.AnimeStatus = *p.AnimeStatus
		}
		SaveUserAnimePreviewData(cacheKey, node)
	} else {
		// just the same type checking but for by-value data
		if node, ok := p.ListNode.(types.AnimeListDataNode); ok {
			cacheKey, _ := GenerateUserAnimePreviewKeys(&node)

			SaveUserAnimePreviewData(cacheKey, &types.UserAnimeListDataNode{
				Node:        node,
				AnimeStatus: *p.AnimeStatus,
			})
		} else if node, ok := p.ListNode.(types.AnimeRankingDataNode); ok {
			cacheKey, _ := GenerateRankingAnimePreviewKeys(&node)

			SaveUserAnimePreviewData(cacheKey, &types.UserAnimeListDataNode{
				Node:        node.Node,
				AnimeStatus: *p.AnimeStatus,
			})
		} else if node, ok := p.ListNode.(types.UserAnimeListDataNode); ok {
			cacheKey, _ := GenerateUserAnimePreviewKeys(&node.Node)

			if p.AnimeStatus != nil {
				node.AnimeStatus = *p.AnimeStatus
			}
			SaveUserAnimePreviewData(cacheKey, &node)
		}
	}

	return
}

type DeleteUserAnimeStatusCacheParams struct {
	ListNode interface{}
}

func DeleteUserAnimeStatusCache(p DeleteUserAnimeStatusCacheParams) {
	// TODO: update the cache for anime status only as it is the only part that can be changed by user.
	var cacheKey string
	if node, ok := p.ListNode.(*types.AnimeListDataNode); ok {
		cacheKey, _ = GenerateUserAnimePreviewKeys(node)
	} else if node, ok := p.ListNode.(*types.AnimeRankingDataNode); ok {
		cacheKey, _ = GenerateRankingAnimePreviewKeys(node)
	} else if node, ok := p.ListNode.(*types.UserAnimeListDataNode); ok {
		cacheKey, _ = GenerateUserAnimePreviewKeys(&node.Node)
	} else {
		// just the same type checking but for by-value data
		if node, ok := p.ListNode.(types.AnimeListDataNode); ok {
			cacheKey, _ = GenerateUserAnimePreviewKeys(&node)
		} else if node, ok := p.ListNode.(types.AnimeRankingDataNode); ok {
			cacheKey, _ = GenerateRankingAnimePreviewKeys(&node)
		} else if node, ok := p.ListNode.(types.UserAnimeListDataNode); ok {
			cacheKey, _ = GenerateUserAnimePreviewKeys(&node.Node)
		}
	}

	if cacheKey == "" {
		return
	}

	dataFileName := strings.ReplaceAll(cacheKey, " ", "")
	dataFileName = strings.ReplaceAll(dataFileName, "\t", "")
	userAnimeDataFilePath := dataDir + "/user/" + dataFileName // NOTE:

	if err := DeleteFile(userAnimeDataFilePath); err != nil {
		fmt.Println("-W: failed to remove the anime status file from cache")
	}

	return
}
