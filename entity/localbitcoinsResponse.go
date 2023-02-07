package entity

import (
	"errors"
	"strconv"
	"strings"
)

// LocalbitcoinsResponse is the main struct of response from localbitcoin api
type LocalbitcoinsResponse struct {
	Pages Pagination `json:"pagination"`
	Data  MainData   `json:"data"`
}

// Pagination struct that holds the previous and next page to search all the ads.
type Pagination struct {
	Prev string `json:"prev"`
	Next string `json:"next"`
}

// MainData struct that holds the list of ads and the counts
type MainData struct {
	Adlist  []Advertisement `json:"ad_list"`
	AdCount int             `json:"ad_count"`
}

// SearchByAmountAndBankFirstMatch is a function that search a ad by amount
func (m *MainData) SearchByAmountAndBankFirstMatch(amount float64, bankName string) (r Advertisement, e error) {
	for _, ad := range m.Adlist {
		minAmount, _ := strconv.ParseFloat(ad.AdInfo.MinAmount, 64)
		maxAmount, _ := strconv.ParseFloat(ad.AdInfo.MaxAmountAvailable, 64)

		if amount >= minAmount && amount <= maxAmount {

			if (strings.Contains(strings.ToUpper(ad.AdInfo.BankName), strings.ToUpper(bankName))) || (bankName == "") {
				r = ad
				return
			}

		}

	}
	e = errors.New("the ad was not found")
	return
}

// SearchByBTCAndBankFirstMatch is a function that search a ad by BTC
func (m *MainData) SearchByBTCAndBankFirstMatch(btcAmount float64, bankName string) (r Advertisement, e error) {
	for _, ad := range m.Adlist {
		minAmount, _ := strconv.ParseFloat(ad.AdInfo.MinAmount, 64)
		maxAmount, _ := strconv.ParseFloat(ad.AdInfo.MaxAmountAvailable, 64)
		tempPrice, _ := strconv.ParseFloat(ad.AdInfo.TempPrice, 64)

		amount := tempPrice * btcAmount

		if amount >= minAmount && amount <= maxAmount {

			if (strings.Contains(strings.ToUpper(ad.AdInfo.BankName), strings.ToUpper(bankName))) || (bankName == "") {
				r = ad
				return
			}

		}

	}
	e = errors.New("the ad was not found")
	return
}

// Advertisement struct that holds the ad information and the url
type Advertisement struct {
	AdInfo  AdvertisementInformation `json:"data"`
	Actions AdvertisementAction      `json:"actions"`
}

// AdvertisementAction struct that hold url from the add
type AdvertisementAction struct {
	PublicView string `json:"public_view"`
}
