package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Continue Continue `json:"continue"`
	Query    Query    `json:"query"`
}
type Continue struct {
	Plcontinue string `json:"plcontinue"`
	Continue   string `json:"continue"`
}
type Links struct {
	Ns    int    `json:"ns"`
	Title string `json:"title"`
}
type ID struct {
	PageID int     `json:"pageid"`
	NS     int     `json:"ns"`
	Title  string  `json:"title"`
	Links  []Links `json:"links"`
}
type Pages struct {
	ID ID
}
type Query struct {
	Pages map[int]json.RawMessage
}

func GetOptions(category string) (this string, that string) {
	reqUrl := "https://en.wikipedia.org/w/api.php?action=query&titles=" + category + "&prop=links&pllimit=max&format=json"

	res, err := http.Get(reqUrl)

	if err != nil {
		log.Fatal("Error retrieving response from wikipedia API", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Error reading response body", err)
	}

	var response Response

	json.Unmarshal(body, &response)

	var key int
	for k, _ := range response.Query.Pages {
		key = k
	}

	var ID ID
	json.Unmarshal(response.Query.Pages[key], &ID)
	// fmt.Println(response.Query.Pages)

	var options [500]string
	for i, val := range ID.Links {
		options[i] = val.Title
	}

	rand.Seed(time.Now().Unix())
	rand1 := rand.Intn(500)
	rand2 := rand.Intn(500)
	for rand1 == rand2 {
		rand2 = rand.Intn(500)
	}

	return options[rand1], options[rand2]
}
