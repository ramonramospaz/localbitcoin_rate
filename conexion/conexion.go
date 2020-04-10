package conexion

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"localbitcoin_rate/entity"
	"net/http"
	"os"
	"strconv"
	"time"
)

const urlBuy = "https://localbitcoins.com/es/buy-bitcoins-online/%v/.json"
const urlSell = "https://localbitcoins.com/es/sell-bitcoins-online/%v/.json"
const urlCoins = "https://localbitcoins.com/api/currencies/"

var listCoins entity.LocalbitcoinsCurrencieResponse

func init() {
	fmt.Println("Searching Currencies allowed in localbitcoin")
	r, err := getLocalbitcoinCurrencieResponse(urlCoins)
	if err != nil {
		fmt.Printf("There was a problem with the internet conection: %v\n", err)
		os.Exit(1)
	}
	listCoins = r
}

func getLocalbitcoinResponse(url string) (r entity.LocalbitcoinsResponse, e error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//Makr HTTP GET request
	response, e := client.Get(url)
	if e != nil {
		return
	}
	defer response.Body.Close()

	dataInBytes, e := ioutil.ReadAll(response.Body)

	e = json.Unmarshal(dataInBytes, &r)
	if e != nil {
		return
	}
	return
}

func getLocalbitcoinCurrencieResponse(url string) (r entity.LocalbitcoinsCurrencieResponse, e error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//Makr HTTP GET request
	response, e := client.Get(url)
	if e != nil {
		return
	}
	defer response.Body.Close()

	dataInBytes, e := ioutil.ReadAll(response.Body)

	e = json.Unmarshal(dataInBytes, &r)
	if e != nil {
		return
	}
	return
}

func checkLocalbitcoinCoins(coin string) (e error) {
	if _, ok := listCoins.Data.Currencies[coin]; !ok {
		e = errors.New("The Currency dont exist")
	}
	return
}

//GetLocalbitcoinRate ...
func GetLocalbitcoinRate(coinIn string, bankNameIn string, coinOut string, bankNameOut string, amount float64) (response entity.LocalbitcoinRateInformation, err error) {
	return getLocalbitcoinRate(coinIn, bankNameIn, coinOut, bankNameOut, amount, getLocalbitcoinResponse)
}

func getLocalbitcoinRate(coinIn string, bankNameIn string, coinOut string, bankNameOut string, amount float64, getResponse func(url string) (r entity.LocalbitcoinsResponse, e error)) (response entity.LocalbitcoinRateInformation, err error) {

	if amount <= 0 {
		err = errors.New("The amount cant be 0 or negative")
		return
	}

	if errIn, errOut := checkLocalbitcoinCoins(coinIn), checkLocalbitcoinCoins(coinOut); errIn != nil || errOut != nil {
		err = errors.New("The Currency dont exist")
		return
	}

	completeURLBuy := fmt.Sprintf(urlBuy, coinIn)
	completeURLSell := fmt.Sprintf(urlSell, coinOut)
	localbitcoinsResponseBuy, err := getResponse(completeURLBuy)
	if err != nil {
		err = errors.New("The buy Ads cant be loaded, please check your internet conection")
		return
	}

	findBuyer := false
	var buyAdvertisement entity.Advertisement
	for findBuyer == false {
		buyAdvertisement, err = localbitcoinsResponseBuy.Data.SearchByAmountAndBankFirstMatch(amount, bankNameIn)
		if err == nil {
			findBuyer = true

		}

		if (findBuyer == false && localbitcoinsResponseBuy.Pages.Next == "") || (findBuyer == true) {
			break
		} else {
			localbitcoinsResponseBuy, _ = getResponse(localbitcoinsResponseBuy.Pages.Next)
		}
	}

	if findBuyer == true {
		buyTempPrice, errFloat := strconv.ParseFloat(buyAdvertisement.AdInfo.TempPrice, 64)

		if errFloat == nil {
			BTC := amount / buyTempPrice
			localbitcoinsResponseSell, errURL := getResponse(completeURLSell)
			if errURL == nil {
				findSeller := false
				var sellAdvertisement entity.Advertisement
				for findSeller == false {
					sellAdvertisement, err = localbitcoinsResponseSell.Data.SearchByBTCAndBankFirstMatch(BTC, bankNameOut)
					if err == nil {
						findSeller = true
					}

					if (findSeller == false && localbitcoinsResponseSell.Pages.Next == "") || (findSeller == true) {
						break
					} else {
						localbitcoinsResponseSell, _ = getResponse(localbitcoinsResponseSell.Pages.Next)
					}
				}

				if findBuyer == true && findSeller == true {
					response, err = entity.GetLocalbitcoinResume(amount, buyAdvertisement, sellAdvertisement)
				} else {
					err = errors.New("Could not find a match")
				}
			}

		}

	}

	return

}
