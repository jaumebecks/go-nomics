package currencies

import (
	"encoding/json"
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

type tickerInterval struct {
	Volume             string `json:"volume"`
	PriceChange        string `json:"price_change"`
	PriceChangePct     string `json:"price_change_pct"`
	VolumeChange       string `json:"volume_change"`
	VolumeChangePct    string `json:"volume_change_pct"`
	MarketCapChange    string `json:"market_cap_change"`
	MarketCapChangePct string `json:"market_cap_change_pct"`
}

// TickerResponse https://nomics.com/docs/#operation/getCurrenciesTicker
type TickerResponse struct {
	ID                string         `json:"id"`
	Currency          string         `json:"currency"`
	Symbol            string         `json:"symbol"`
	Name              string         `json:"name"`
	LogoURL           string         `json:"logo_url"`
	Status            string         `json:"status"`
	Price             string         `json:"price"`
	PriceDate         string         `json:"price_date"`
	PriceTimestamp    string         `json:"price_timestamp"`
	CirculatingSupply string         `json:"circulating_supply"`
	MaxSupply         string         `json:"max_supply"`
	MarketCap         string         `json:"market_cap"`
	NumExchanges      string         `json:"num_exchanges"`
	NumPairs          string         `json:"num_pairs"`
	NumPairsUnmapped  string         `json:"num_pairs_unmapped"`
	FirstCandle       string         `json:"first_candle"`
	FirstTrade        string         `json:"first_trade"`
	FirstOrderBook    string         `json:"first_order_book"`
	Rank              string         `json:"rank"`
	RankDelta         string         `json:"rank_delta"`
	High              string         `json:"high"`
	HighTimestamp     string         `json:"high_timestamp"`
	Interval1Hour     tickerInterval `json:"1h"`
	Interval1Day      tickerInterval `json:"1d"`
	Interval7Days     tickerInterval `json:"7d"`
	Interval30Days    tickerInterval `json:"30d"`
	Interval365Days   tickerInterval `json:"365d"`
	IntervalYTD       tickerInterval `json:"ytd"`
}

// Ticker returns Price, volume, market cap, and rank for all currencies across
// 1 hour, 1 day, 7 day, 30 day, 365 day, and year to date intervals.
// Current prices are updated every 10 seconds.
func Ticker(apiKey string, query TickerQuery) []TickerResponse {
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

	tickerResponses := []TickerResponse{}
	error = json.Unmarshal(body, &tickerResponses)
	if error != nil {
		log.Fatal(error)
	}

	return tickerResponses
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
