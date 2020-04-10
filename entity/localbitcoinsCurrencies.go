package entity

// LocalbitcoinsCurrencieResponse is the main struct of response from localbitcoin api
type LocalbitcoinsCurrencieResponse struct {
	Data CurrencieData `json:"data"`
}

// CurrencieData ...
type CurrencieData struct {
	Currencies map[string]CurrencieInfo `json:"currencies"`
}

// CurrencieInfo hold the info from the currencie
type CurrencieInfo struct {
	Name    string `json:"name"`
	Altcoin bool   `json:"altcoin"`
}
