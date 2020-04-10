package entity

import (
	"strconv"
)

// LocalbitcoinRateInformation is the result struct with the resume of the information
type LocalbitcoinRateInformation struct {
	CurrencyBuy    string
	BankNameBuy    string
	AmountBuy      float64
	PublicViewBuy  string
	CurrencySell   string
	BankNameSell   string
	AmountSell     float64
	PublicViewSell string
	Rate           float64
}

// AdvertisementInformation struct that holds de information from ads
type AdvertisementInformation struct {
	RequireFeedbackScore       int     `json:"require_feedback_score"`
	HiddenByOpeningHours       bool    `json:"hidden_by_opening_hours"`
	Currency                   string  `json:"currency"`
	RequireIdentification      bool    `json:"require_identification"`
	IsLocalOffice              bool    `json:"is_local_office"`
	FirstTimeLimitBtc          string  `json:"fitst_time_limit_btc"`
	City                       string  `json:"city"`
	LocationString             string  `json:"location_string"`
	CountryCode                string  `json:"countrycode"`
	MaxAmount                  string  `json:"max_amount"`
	Longitud                   float64 `json:"lon"`
	SmsVerificationRequired    bool    `json:"sms_verification_required"`
	RequireTradeVolume         float64 `json:"require_trade_volume"`
	OnlineProvider             string  `json:"online_provider"`
	MaxAmountAvailable         string  `json:"max_amount_available"`
	Msg                        string  `json:"msg"`
	VolumeCoefficientBtc       string  `json:"volume_coefficient_btc"`
	AdsProfile                 Profile `json:"profile"`
	BankName                   string  `json:"bank_name"`
	TradeType                  string  `json:"trade_type"`
	AdID                       int     `json:"ad_id"`
	TempPrice                  string  `json:"temp_price"`
	PaymentWindowMinutes       int     `json:"payment_window_minutes"`
	MinAmount                  string  `json:"min_amount"`
	LimitToFiatAmounts         string  `json:"limit_to_fiat_amounts"`
	RequireTrustedByAdvertiser bool    `json:"require_trusted_by_advertiser"`
	TempPriceUsd               string  `json:"temp_price_usd"`
	Latitud                    float64 `json:"lat"`
	Visible                    bool    `json:"visible"`
	CreatedAt                  string  `json:"created_at"`
	AtmModel                   string  `json:"atm_model"`
}

// Profile struct have the info from the user that create the ad
type Profile struct {
	Username      string `json:"username"`
	FeedbackScore int    `json:"feedback_score"`
	TradeCount    string `json:"trade_count"`
	LastOnline    string `json:"last_online"`
	Name          string `json:"name"`
}

// GetLocalbitcoinResume ...
func GetLocalbitcoinResume(amount float64, buy Advertisement, sell Advertisement) (r LocalbitcoinRateInformation, e error) {
	r.CurrencyBuy = buy.AdInfo.Currency
	r.BankNameBuy = buy.AdInfo.BankName
	r.AmountBuy = amount
	r.PublicViewBuy = buy.Actions.PublicView
	r.CurrencySell = sell.AdInfo.Currency
	r.BankNameSell = sell.AdInfo.BankName
	sellTempPrice, e := strconv.ParseFloat(sell.AdInfo.TempPrice, 64)
	if e != nil {
		return
	}
	buyTempPrice, e := strconv.ParseFloat(buy.AdInfo.TempPrice, 64)
	if e != nil {
		return
	}
	r.AmountSell = sellTempPrice * (amount / buyTempPrice)
	r.PublicViewSell = sell.Actions.PublicView
	r.Rate = r.AmountSell / r.AmountBuy
	return
}
