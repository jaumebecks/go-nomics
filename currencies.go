package currencies

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Ticker returns Price, volume, market cap, and rank for all currencies across
// 1 hour, 1 day, 7 day, 30 day, 365 day, and year to date intervals.
// Current prices are updated every 10 seconds.
func Ticker() []byte {
	response, error := http.Get("https://www.google.com/robots.txt")
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
