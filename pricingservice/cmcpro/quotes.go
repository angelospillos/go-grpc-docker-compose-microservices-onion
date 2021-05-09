package cmcpro

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type GetLatestPricesBySymbolsResponse struct {
	Status struct {
		Timestamp    time.Time `json:"timestamp,omitempty"`
		ErrorCode    int       `json:"error_code,omitempty"`
		ErrorMessage string    `json:"error_message,omitempty"`
		Elapsed      int       `json:"elapsed,omitempty"`
		CreditCount  int       `json:"credit_count,omitempty"`
		Notice       string    `json:"notice,omitempty"`
	}
	Data map[string]struct {
		ID                int         `json:"id,omitempty"`
		Name              string      `json:"name,omitempty"`
		Symbol            string      `json:"symbol,omitempty"`
		Slug              string      `json:"slug,omitempty"`
		IsActive          int         `json:"is_active,omitempty"`
		IsFiat            int         `json:"is_fiat,omitempty"`
		CirculatingSupply float32     `json:"circulating_supply,omitempty"`
		TotalSupply       float32     `json:"total_supply,omitempty"`
		MaxSupply         float32     `json:"max_supply,omitempty"`
		DateAdded         time.Time   `json:"date_added,omitempty"`
		NumMarketPairs    int         `json:"num_market_pairs,omitempty"`
		CmcRank           int         `json:"cmc_rank,omitempty"`
		LastUpdated       time.Time   `json:"last_updated,omitempty"`
		Tags              []string    `json:"tags,omitempty"`
		Platform          interface{} `json:"platform,omitempty"`
		Quote             map[string]struct {
			Price            float32   `json:"price,omitempty"`
			Volume24h        float32   `json:"volume_24h,omitempty"`
			PercentChange1h  float32   `json:"percent_change_1h,omitempty"`
			PercentChange24h float32   `json:"percent_change_24h,omitempty"`
			PercentChange7d  float32   `json:"percent_change_7d,omitempty"`
			PercentChange30d float32   `json:"percent_change_30d,omitempty"`
			MarketCap        float32   `json:"market_cap,omitempty"`
			LastUpdated      time.Time `json:"last_updated,omitempty"`
		} `json:"quote"`
	} `json:"data"`
}

type GetLatestPricesBySymbolsRequest struct {
	Symbols []string
}

type Api struct {
	URL string
	Key string
}

func (api Api) GetLatestPricesBySymbols(request GetLatestPricesBySymbolsRequest) (GetLatestPricesBySymbolsResponse, error) {

	client := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, api.URL+"/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", api.Key)

	q := req.URL.Query()
	q.Add("symbol", strings.Join(request.Symbols[:], ","))
	req.URL.RawQuery = q.Encode()

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	getLatestPricesBySymbolsResponse := GetLatestPricesBySymbolsResponse{}

	jsonErr := json.Unmarshal(body, &getLatestPricesBySymbolsResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return getLatestPricesBySymbolsResponse, nil
}
