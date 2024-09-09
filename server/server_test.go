package server_test

/*
import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test_Client(t *testing.T) {
	client := &http.Client{}


	//req, err := http.NewRequest("GET", cfg.URL + "/anime?q=one&limit=4", nil)
    url := fmt.Sprintf("%v/oauth2/authorize?"+
		"response_type=%v"+
		"&client_id=%v"+
		"&code_challenge=%v",
		cfg.SITE_URL,
		"code",
		cfg.CLIENT_ID,
		"jaksbcaskcbasjkcbasjkcbasjkcbasjckio12826392938347jk")

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("client_id", cfg.CLIENT_ID)
	req.Header.Add("response_type", "code")
	req.Header.Add("code_challenge", "abcdef")
	req.Header.Add("code_challenge_method", "plain")

	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		t.Fatalf("ERROR SENDING REQ %v", err)
	}
	t.Log(res)

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("ERROR READING REQ %v", err)
	}

	t.Log(string(body))

} */
