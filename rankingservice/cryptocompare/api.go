package cryptocompare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TopByMarketCapResponse struct {
	Message  string `json:"Message,omitempty"`
	Type     int    `json:"Type,omitempty"`
	Metadata struct {
		Count int `json:"Count,omitempty"`
	} `json:"MetaData,omitempty"`
	SponsoredData []struct {
		CoinInfo struct {
			ID        string `json:"Id,omitempty"`
			Name      string `json:"Name,omitempty"`
			FullName  string `json:"FullName,omitempty"`
			Internal  string `json:"Internal,omitempty"`
			ImageUrl  string `json:"ImageUrl,omitempty"`
			URL       string `json:"Url,omitempty"`
			Algorithm string `json:"Algorithm,omitempty"`
			ProofType string `json:"ProofType,omitempty"`
			Rating    struct {
				Weiss struct {
					Rating                   string `json:"Rating,omitempty"`
					TechnologyAdoptionRating string `json:"TechnologyAdoptionRating,omitempty"`
					MarketPerformanceRating  string `json:"MarketPerformanceRating,omitempty"`
				} `json:"Weiss,omitempty"`
			} `json:"Rating,omitempty"`
			NetHashesPerSecond float32    `json:"NetHashesPerSecond,omitempty"`
			BlockNumber        int    `json:"BlockNumber,omitempty"`
			BlockTime          int    `json:"BlockTime,omitempty"`
			BlockReward        int    `json:"BlockReward,omitempty"`
			AssetLaunchDate    string `json:"AssetLaunchDate,omitempty"`
			MaxSupply          int    `json:"MaxSupply,omitempty"`
			Type               int    `json:"Type,omitempty"`
			DocumentType       string `json:"DocumentType,omitempty"`
		} `json:"CoinInfo,omitempty"`
	} `json:"SponsoredData,omitempty"`
	Data []struct {
		CoinInfo struct {
			ID        string `json:"Id,omitempty"`
			Name      string `json:"Name,omitempty"`
			FullName  string `json:"FullName,omitempty"`
			Internal  string `json:"Internal,omitempty"`
			ImageUrl  string `json:"ImageUrl,omitempty"`
			URL       string `json:"Url,omitempty"`
			Algorithm string `json:"Algorithm,omitempty"`
			ProofType string `json:"ProofType,omitempty"`
			Rating    struct {
				Weiss struct {
					Rating                   string `json:"Rating,omitempty"`
					TechnologyAdoptionRating string `json:"TechnologyAdoptionRating,omitempty"`
					MarketPerformanceRating  string `json:"MarketPerformanceRating,omitempty"`
				} `json:"Weiss,omitempty"`
			} `json:"Rating,omitempty"`
			NetHashesPerSecond float64   `json:"NetHashesPerSecond,omitempty"`
			BlockNumber        int     `json:"BlockNumber,omitempty"`
			BlockTime          float64 `json:"BlockTime,omitempty"`
			BlockReward        float64 `json:"BlockReward,omitempty"`
			AssetLaunchDate    string  `json:"AssetLaunchDate,omitempty"`
			MaxSupply          float64 `json:"MaxSupply,omitempty"`
			Type               int     `json:"Type,omitempty"`
			DocumentType       string  `json:"DocumentType,omitempty"`
		} `json:"CoinInfo,omitempty"`
		Raw map[string]struct {
			Type                    string  `json:"TYPE,omitempty"`
			Market                  string  `json:"MARKET,omitempty"`
			FromSymbol              string  `json:"FROMSYMBOL,omitempty"`
			ToSymbol                string  `json:"TOSYMBOL,omitempty"`
			Flags                   string  `json:"FLAGS,omitempty"`
			Price                   float64 `json:"PRICE,omitempty"`
			LastUpdate              int     `json:"LASTUPDATE,omitempty"`
			Median                  float64 `json:"MEDIAN,omitempty"`
			LastVolume              float64 `json:"LASTVOLUME,omitempty"`
			LastVolumeTo            float64 `json:"LASTVOLUMETO,omitempty"`
			LastTradeId             string  `json:"LASTTRADEID,omitempty"`
			VolumeDay               float64 `json:"VOLUMEDAY,omitempty"`
			VolumeDayTo             float64 `json:"VOLUMEDAYTO,omitempty"`
			Volume24Hour            float64 `json:"VOLUME24HOUR,omitempty"`
			Volume24HourTo          float64 `json:"VOLUME24HOURTO,omitempty"`
			OpenDay                 float64 `json:"OPENDAY,omitempty"`
			HighDay                 float64 `json:"HIGHDAY,omitempty"`
			LowDay                  float64 `json:"LOWDAY,omitempty"`
			Open24Hour              float64 `json:"OPEN24HOUR,omitempty"`
			High24Hour              float64 `json:"HIGH24HOUR,omitempty"`
			Low24Hour               float64 `json:"LOW24HOUR,omitempty"`
			LastMarket              string  `json:"LASTMARKET,omitempty"`
			VolumeHour              float64 `json:"VOLUMEHOUR,omitempty"`
			VolumeHourTo            float64 `json:"VOLUMEHOURTO,omitempty"`
			OpenHour                float64 `json:"OPENHOUR,omitempty"`
			HighHour                float64 `json:"HIGHHOUR,omitempty"`
			LowHour                 float64 `json:"LOWHOUR,omitempty"`
			TopTierVolume24Hour     float64 `json:"TOPTIERVOLUME24HOUR,omitempty"`
			TopTierVolume24HourTo   float64 `json:"TOPTIERVOLUME24HOURTO,omitempty"`
			Change24Hour            float64 `json:"CHANGE24HOUR,omitempty"`
			ChangePct24Hour         float64 `json:"CHANGEPCT24HOUR,omitempty"`
			ChangeDay               float64 `json:"CHANGEDAY,omitempty"`
			ChangePctDay            float64 `json:"CHANGEPCTDAY,omitempty"`
			ChangeHour              float64 `json:"CHANGEHOUR,omitempty"`
			ChangePctHour           float64 `json:"CHANGEPCTHOUR,omitempty"`
			ConversionType          string  `json:"CONVERSIONTYPE,omitempty"`
			ConversionSymbol        string  `json:"CONVERSIONSYMBOL,omitempty"`
			Supply                  float64     `json:"SUPPLY,omitempty"`
			MktCap                  float64 `json:"MKTCAP,omitempty"`
			MktCapPenalty           int     `json:"MKTCAPPENALTY,omitempty"`
			TotalVolume24H          float64 `json:"TOTALVOLUME24H,omitempty"`
			TotalVolume24Hto        float64 `json:"TOTALVOLUME24HTO,omitempty"`
			TotalTopTierVolume24H   float64 `json:"TOTALTOPTIERVOLUME24H,omitempty"`
			TotalTopTierVolume24HTo float64 `json:"TOTALTOPTIERVOLUME24HTO,omitempty"`
			ImageUrl                string  `json:"IMAGEURL,omitempty"`
		} `json:"RAW,omitempty"`
		Display map[string]struct {
			FromSymbol              string `json:"FROMSYMBOL,omitempty"`
			ToSymbol                string `json:"TOSYMBOL,omitempty"`
			Market                  string `json:"MARKET,omitempty"`
			Price                   string `json:"PRICE,omitempty"`
			LastUpdate              string `json:"LASTUPDATE,omitempty"`
			LastVolume              string `json:"LASTVOLUME,omitempty"`
			LastVolumeTo            string `json:"LASTVOLUMETO,omitempty"`
			LastTradeId             string `json:"LASTTRADEID,omitempty"`
			VolumeDay               string `json:"VOLUMEDAY,omitempty"`
			VolumeDayTo             string `json:"VOLUMEDAYTO,omitempty"`
			Volume24Hour            string `json:"VOLUME24HOUR,omitempty"`
			Volume24HourTo          string `json:"VOLUME24HOURTO,omitempty"`
			OpenDay                 string `json:"OPENDAY,omitempty"`
			HighDay                 string `json:"HIGHDAY,omitempty"`
			LowDay                  string `json:"LOWDAY,omitempty"`
			Open24Hour              string `json:"OPEN24HOUR,omitempty"`
			High24Hour              string `json:"HIGH24HOUR,omitempty"`
			Low24Hour               string `json:"LOW24HOUR,omitempty"`
			LastMarket              string `json:"LASTMARKET,omitempty"`
			VolumeHour              string `json:"VOLUMEHOUR,omitempty"`
			VolumeHourTo            string `json:"VOLUMEHOURTO,omitempty"`
			OpenHour                string `json:"OPENHOUR,omitempty"`
			HighHour                string `json:"HIGHHOUR,omitempty"`
			LowHour                 string `json:"LOWHOUR,omitempty"`
			TopTierVolume24Hour     string `json:"TOPTIERVOLUME24HOUR,omitempty"`
			TopTierVolume24HourTo   string `json:"TOPTIERVOLUME24HOURTO,omitempty"`
			Change24Hour            string `json:"CHANGE24HOUR,omitempty"`
			ChangePct24Hour         string `json:"CHANGEPCT24HOUR,omitempty"`
			ChangeDay               string `json:"CHANGEDAY,omitempty"`
			ChangePctDay            string `json:"CHANGEPCTDAY,omitempty"`
			ChangeHour              string `json:"CHANGEHOUR,omitempty"`
			ChangePctHour           string `json:"CHANGEPCTHOUR,omitempty"`
			ConversionType          string `json:"CONVERSIONTYPE,omitempty"`
			ConversionSymbol        string `json:"CONVERSIONSYMBOL,omitempty"`
			Supply                  string `json:"SUPPLY,omitempty"`
			MktCap                  string `json:"MKTCAP,omitempty"`
			MktCapPenalty           string `json:"MKTCAPPENALTY,omitempty"`
			TotalVolume24H          string `json:"TOTALVOLUME24H,omitempty"`
			TotalVolume24Hto        string `json:"TOTALVOLUME24HTO,omitempty"`
			TotalTopTierVolume24H   string `json:"TOTALTOPTIERVOLUME24H,omitempty"`
			TotalTopTierVolume24Hto string `json:"TOTALTOPTIERVOLUME24HTO,omitempty"`
			ImageUrl                string `json:"IMAGEURL,omitempty"`
		} `json:"DISPLAY,omitempty"`
	} `json:"Data,omitempty"`
	RateLimit struct {
	} `json:"RateLimit,omitempty"`
	HasWarning bool `json:"HasWarning,omitempty"`
}

type TopByMarketCapRequest struct {
	Limit int32
}

type Api struct {
	URL string
	Key string
	BaseCurrency string
}

func (api Api) GetTopByMarketCap(request TopByMarketCapRequest) (TopByMarketCapResponse, error) {

	client := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, api.URL+"/data/top/mktcapfull", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authorization", api.Key)

	q := req.URL.Query()
	q.Add("limit", fmt.Sprintf("%v", request.Limit))
	q.Add("tsym", api.BaseCurrency)
	req.URL.RawQuery = q.Encode()

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	topByMarketCapResponse := TopByMarketCapResponse{}

	jsonErr := json.Unmarshal(body, &topByMarketCapResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr.Error())
	}

	return topByMarketCapResponse, nil
}
