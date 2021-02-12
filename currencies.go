package currencies

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://api.nomics.com/v1/currencies/ticker"

// TickerQuery https://nomics.com/docs/#operation/getCurrenciesTicker
type TickerQuery struct {
	Ids          []string
	Interval     []string
	Convert      string
	Status       string
	Filter       string
	Sort         string
	Transparency bool
	Limit        int
	Page         int
}

// Ticker returns Price, volume, market cap, and rank for all currencies across
// 1 hour, 1 day, 7 day, 30 day, 365 day, and year to date intervals.
// Current prices are updated every 10 seconds.
func Ticker(apiKey string, query TickerQuery) []byte {
	URL, error := url.Parse(baseURL)
	if error != nil {
		log.Fatal("Error building API base URL")
	}
	buildQuery(URL, apiKey, query)

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

func buildQuery(URL *url.URL, apiKey string, query TickerQuery) {
	urlQuery := url.Values{}
	urlQuery.Add("key", apiKey)
	if len(query.Ids) != 0 {
		urlQuery.Add("ids", strings.Join(query.Ids, ","))
	}
	if len(query.Interval) != 0 {
		urlQuery.Add("interval", strings.Join(query.Interval, ","))
	}
	if len(query.Convert) != 0 {
		urlQuery.Add("convert", query.Convert)
	}
	if len(query.Status) != 0 {
		urlQuery.Add("status", query.Status)
	}
	if len(query.Filter) != 0 {
		urlQuery.Add("filter", query.Filter)
	}
	if len(query.Sort) != 0 {
		urlQuery.Add("sort", query.Sort)
	}
	if query.Transparency {
		urlQuery.Add("include-transparency", strconv.FormatBool(query.Transparency))
	}
	if query.Limit != 0 {
		urlQuery.Add("per-page", strconv.FormatInt(int64(query.Limit), 10))
	}
	if query.Page != 0 {
		urlQuery.Add("page", strconv.FormatInt(int64(query.Page), 10))
	}

	URL.RawQuery = urlQuery.Encode()
}
