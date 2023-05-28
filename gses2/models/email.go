package models

type Email struct {
	Address string `json:"email"`
}

type CoinGeckoResponse struct {
	MarketData struct {
		CurrentPrice struct {
			UAH float64 `json:"uah"`
		} `json:"current_price"`
	} `json:"market_data"`
}
