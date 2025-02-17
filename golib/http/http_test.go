package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type Holder struct {
	Address string `json:"address"`
	Amount uint64 `json:"amount"`
	Decimals int `json:"decimals"`
	Owner string `json:"owner"`
	Rank int `json:"rank"`
}

type Data struct {
	Total int `json:"total"`
	Result []Holder `json:"result"`
}

type Response struct {
	Success bool `json:"succcess"`
	Data Data `json:"data"`
}

func TestGet1(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.solscan.io/token/holders", nil)
	if err != nil {
		panic(err)
	}


	checks := make([]string, 0)
	for _, check := range checks {
		q := url.Values{}
		q.Add("token", check)
		q.Add("offset", "0")
		q.Add("size", "20")

		req.Header.Set("Accepts", "application/json")
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			panic(fmt.Errorf("response status code: %d", resp.StatusCode))
		}
		respBody, _ := ioutil.ReadAll(resp.Body)
		var body Response
		err = json.Unmarshal(respBody, &body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\"%s\":\"%s\",\n", check, body.Data.Result[10].Address)
	}
}

func TestGet(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.solscan.io/token/holders", nil)
	if err != nil {
		panic(err)
	}

	{
		q := url.Values{}
		q.Add("token", "JEFFSQ3s8T3wKsvp4tnRAsUBW7Cqgnf8ukBZC4C8XBm1")
		q.Add("offset", "0")
		q.Add("size", "20")

		req.Header.Set("Accepts", "application/json")
		req.URL.RawQuery = q.Encode()


		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			panic(fmt.Errorf("response status code: %d", resp.StatusCode))
		}
		respBody, _ := ioutil.ReadAll(resp.Body)
		var body Response
		err = json.Unmarshal(respBody, &body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", body)
	}
}
