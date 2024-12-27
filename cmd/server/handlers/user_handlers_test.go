package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	e "github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	p "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

func Test_GETUserAnimeList(t *testing.T) {
    // TODO: test the sort param too
	var status e.UserAnimeListStatus = ""

	sort := make([]e.UserAnimeListSortOption, 0)
	sort = append(sort, e.UALSort_ListUpdatedAt)

	detailFields := make([]es.AnimeDetailField, 0)
	detailFields = append(detailFields, es.Title, es.StartSeason)

	{ // all user anime list
		fmt.Println("Fetching all user anime list...")
		jsonData := fetchUserAnimeList(status, sort, detailFields, t)

		count := strings.Count(string(jsonData), `"id"`)
		fmt.Println("Total Anime Count := ", count)
		fmt.Println("--------------------------------------------------------")
	}
	{ // dropped anime list
		status = e.ULS_Dropped
		fmt.Println("Fetching dropped user anime list...")
		jsonData := fetchUserAnimeList(status, sort, detailFields, t)

		count := strings.Count(string(jsonData), `"id"`)
		fmt.Println("Total Anime Count := ", count)
		fmt.Println("--------------------------------------------------------")
	}
	{ // watching anime list
		status = e.ULS_Watching
		fmt.Println("Fetching watching user anime list...")
		jsonData := fetchUserAnimeList(status, sort, detailFields, t)

		count := strings.Count(string(jsonData), `"id"`)
		fmt.Println("Total Anime Count := ", count)
		fmt.Println("--------------------------------------------------------")
	}
	{ // onHold anime list
		status = e.ULS_OnHold
		fmt.Println("Fetching On Hold user anime list...")
		jsonData := fetchUserAnimeList(status, sort, detailFields, t)

		count := strings.Count(string(jsonData), `"id"`)
		fmt.Println("Total Anime Count := ", count)
		fmt.Println("--------------------------------------------------------")
	}
	{ // plan_to_watch anime list
		status = e.ULS_PlanToWatch
		fmt.Println("Fetching plan to watch user anime list...")
		jsonData := fetchUserAnimeList(status, sort, detailFields, t)

		count := strings.Count(string(jsonData), `"id"`)
		fmt.Println("Total Anime Count := ", count)
		fmt.Println("--------------------------------------------------------")
	}

	{ // error prone
		fmt.Println("Fetching not listed user anime list...")
		status = "empty"
		jsonData := fetchUserAnimeList(status, sort, detailFields, t)

		count := strings.Count(string(jsonData), `"id"`)
		fmt.Println("Total Anime Count := ", count)
		fmt.Println("--------------------------------------------------------")
        if count == 0 {
            t.Fail();
            t.Log("no anime should be expected")
        }

	}
}


func Test_GETUserInfo(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:42069/api/user"))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	var data any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	jsonData, _ := json.MarshalIndent(data, "\t", "")
	fmt.Println(string(jsonData))
}

func Test_UserHandlers(t *testing.T) {
	url := fmt.Sprintf("https://myanimelist.net/v1/oauth2/authorize?response_type=code&client_id=%s&code_challenge=%s&state=Request123&code_challenge_method=plain",
		config.C.MalClientId,
		config.C.MalCodeChallenge,
	)

	fmt.Println("Redirecting to your browser for authorization...")

	err := p.OpenURL(url)
	if err != nil {
		fmt.Println("Failed to open browser:", err)
		fmt.Println("Please manually visit the URL below to authorize:")
		fmt.Println(url)
	}

	return
	/* fmt.Print("Enter the code you recieved in the callback url query paramater: ")

	    input := "def50200b49692b1a53d9d1452a825cd1a8eded1300fe99238888dc64bbbd2e3a7c82d235f3f6496cdb08aa9464fe108dd76182b336f4d0bf1d3953b9ea5d65907543bb064c67460d83a8f66fc10f190e50b633b0ce4606803e5fed91e63ba891beed5060ccb6aecf9429ba6380c72d53b38d9caa3dd0d7a0bbea803645560c78ac229ff91d205545bdb87e1056ca68c0e5f19c566d4df602b4febb5d0d4de2f5075e0ccc5d30806c90674d77df4afa3f2d6281876cbe35e4a4f406603a6ea95220130bafd10744daaa7aa04bb712c9be839ef995f9be84a042e0ff624a328f762ca0a4617eb6c56f5b2bfc7e5c9923a54d5c48e0b98cfcbc0967d1faf516b2796cb88a74bcd8a0f1db0f2faac19659f64d4bb80d461d1bd5536deda76468741d0a298a3b4bd9f9b1130e856f4cfa2440c641b230beb3872b8b29335f34e88072b500aaa325c593971caa64cbd693264d03072bbb3f7fc7d094ffba5963bf5d998a2733bdffc27d21f0d8a46460d4284f19de21354ce9088a35cc854d6ab96283b37711d977bcf8fed0e2eba91650134d3885d38964a67b08940ad60ececdf1daab0c0f6ee48aac503e6\n"

		reader := bufio.NewReader(strings.NewReader(input))
		authCode := ""

		for len(authCode) == 0 {
			input, _ := reader.ReadString('\n')
			authCode = strings.TrimSpace(input)
		}

		_, err := http.Get(fmt.Sprintf("http://localhost:42069/api/auth?code=%s", authCode))
		if err != nil {
			fmt.Println(err)
		} */
}

func fetchUserAnimeList(status e.UserAnimeListStatus, sort []e.UserAnimeListSortOption, detailFields []es.AnimeDetailField, t *testing.T) string {
	url := fmt.Sprintf("http://localhost:42069/api/user/anime-list?status=%s&sort=%s&limit=%d&fields=%s",
		status,
		p.ConvertToCommaSeperatedString(sort),
		1000,
		p.ConvertToCommaSeperatedString(detailFields),
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return ""
	}

	var data any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	jsonData, _ := json.MarshalIndent(data, "\t", " ")
	return string(jsonData)
}

