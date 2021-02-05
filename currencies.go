package currencies

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Ticker returns Price, volume, market cap, and rank for all currencies across
// 1 hour, 1 day, 7 day, 30 day, 365 day, and year to date intervals.
// Current prices are updated every 10 seconds.
func Ticker(apiKey string, params map[string]string) []byte {
	URL, error := url.Parse("https://api.nomics.com/v1/currencies/ticker")
	if error != nil {
		log.Fatal("Error building API base URL")
	}
	buildQuery(URL, apiKey, params)

	response, error := http.Get(URL.String())
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}

	return body
}

func buildQuery(URL *url.URL, apiKey string, params map[string]string) {
	query := url.Values{}
	query.Add("key", apiKey)
	if _, ok := params["ids"]; ok == true {
		query.Add("ids", params["ids"])
	}
	if _, ok := params["convert"]; ok == true {
		query.Add("convert", params["convert"])
	}
	URL.RawQuery = query.Encode()
}
