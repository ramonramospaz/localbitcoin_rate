package conexion

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func getListCoins() {
	r, err := getLocalbitcoinCurrencieResponse(urlCoins)
	if err != nil {
		fmt.Printf("There was a problem with the internet conection: %v\n", err)
		os.Exit(1)
	}
	listCoins = r
	listCoins.Ready = true
}

//ShowListCoins show the list of all currency that localbitcoin can work with.
func ShowListCoins() entity.LocalbitcoinsCurrencieResponse {
	fmt.Println("Searching the list of currency in the web page localbitcoin. This is going to take some time....")
	getListCoins()
	return listCoins
}

func getHttpResponse(url string) (dataInBytes []byte, e error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//Create request
	request, e := http.NewRequest("GET", url, nil)

	if e != nil {
		return
	}
	//Add the header for read compresion

	request.Header.Add("Accept-Encoding", "gzip")

	//Makr HTTP GET request
	response, e := client.Do(request)
	if e != nil {
		return
	}
	defer response.Body.Close()

	//Check if the server send compressd data
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, e = gzip.NewReader(response.Body)
		if e != nil {
			return
		}

		dataInBytes, e = ioutil.ReadAll(reader)
	default:
		dataInBytes, e = ioutil.ReadAll(response.Body)
	}

	return
}

func getLocalbitcoinResponse(url string) (r entity.LocalbitcoinsResponse, e error) {

	var localbitcoinsResponse entity.LocalbitcoinsResponse

	dataInBytes, e := getHttpResponse(url)

	e = json.Unmarshal(dataInBytes, &localbitcoinsResponse)
	if e != nil {
		return
	}
	r = localbitcoinsResponse

	return
}

func getLocalbitcoinCurrencieResponse(url string) (r entity.LocalbitcoinsCurrencieResponse, e error) {
	var localbitcoinsCurrencieResponse entity.LocalbitcoinsCurrencieResponse

	dataInBytes, e := getHttpResponse(url)

	e = json.Unmarshal(dataInBytes, &localbitcoinsCurrencieResponse)
	if e != nil {
		return
	}
	r = localbitcoinsCurrencieResponse

	return
}

func checkLocalbitcoinCoins(coin string) (e error) {

	if listCoins.Ready == false {
		getListCoins()
	}

	if _, ok := listCoins.Data.Currencies[coin]; !ok {
		e = errors.New("The Currency dont exist")
	}
	return
}

//GetLocalbitcoinRate ...
func GetLocalbitcoinRate(coinIn string, bankNameIn string, coinOut string, bankNameOut string, amount float64) (response entity.LocalbitcoinRateInformation, err error) {
	return getLocalbitcoinRate(coinIn, bankNameIn, coinOut, bankNameOut, amount, getLocalbitcoinResponse)
}

func getLocalbitcoinBuyAd(coinIn string, bankNameIn string, amount float64, getResponse func(url string) (r entity.LocalbitcoinsResponse, e error)) (buyAdvertisement entity.Advertisement, err error) {
	completeURLBuy := fmt.Sprintf(urlBuy, coinIn)
	localbitcoinsResponseBuy, err := getResponse(completeURLBuy)
	if err != nil {
		err = errors.New("The buy Ads cant be loaded, please check your internet conection")
		return
	}

	findBuyer := false
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

	if findBuyer == false {
		err = errors.New("The buy Ads cant be Found")
		return
	}

	return
}

func getLocalbitcoinSellAd(coinOut string, bankNameOut string, BTC float64, getResponse func(url string) (r entity.LocalbitcoinsResponse, e error)) (sellAdvertisement entity.Advertisement, err error) {
	completeURLSell := fmt.Sprintf(urlSell, coinOut)
	localbitcoinsResponseSell, errURL := getResponse(completeURLSell)
	if errURL != nil {
		err = errors.New("The Sell Ads cant be loaded, please check your internet conection")
		return
	}

	findSeller := false
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

	if findSeller == false {
		err = errors.New("The Sell Ads cant be Found")
		return
	}

	return
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

	buyAdvertisement, err := getLocalbitcoinBuyAd(coinIn, bankNameIn, amount, getResponse)

	if err != nil {
		return
	}

	buyTempPrice, errFloat := strconv.ParseFloat(buyAdvertisement.AdInfo.TempPrice, 64)

	if errFloat != nil {
		err = errors.New("The prince of the ads cant be parse, invalid amount")
		return
	}

	BTC := amount / buyTempPrice

	sellAdvertisement, err := getLocalbitcoinSellAd(coinOut, bankNameOut, BTC, getResponse)

	if err != nil {
		return
	}

	response, err = entity.GetLocalbitcoinResume(amount, buyAdvertisement, sellAdvertisement)

	return

}
