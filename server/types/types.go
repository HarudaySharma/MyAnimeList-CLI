package types

type MALAnimeList struct {
	Data []struct {
		Node struct {
			Id          int    `json:"id"`
			Title       string `json:"title"`
			MainPicture struct {
				Large  string `json:"large"`
				Medium string `json:"medium"`
			} `json:"main_picture"`
		} `json:"node"`
	} `json:"data"`
	Paging struct {
		Next string `json:"next"`
	} `json:"paging"`
}

type AnimeListDataNode struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type NativeAnimeList struct {
	Data []AnimeListDataNode `json:"data"`
	Paging struct {
		Next string `json:"next"`
	} `json:"paging"`
}
