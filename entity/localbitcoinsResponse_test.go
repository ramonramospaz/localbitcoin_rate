package entity

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
)

func createDummyDataAdsInfo() (r AdvertisementInformation) {

	r.RequireFeedbackScore = 0
	r.HiddenByOpeningHours = false
	r.Currency = "VED"
	r.RequireIdentification = true
	r.IsLocalOffice = false
	r.FirstTimeLimitBtc = ""
	r.City = ""
	r.LocationString = "Venezuela"
	r.CountryCode = "VE"
	r.MaxAmount = "6100000"
	r.Longitud = 0
	r.SmsVerificationRequired = true
	r.RequireTradeVolume = 0
	r.OnlineProvider = "SPECIFIC_BANK"
	r.MaxAmountAvailable = "6100000.00"
	r.Msg = "Libero apenas este verificado el pago 100% online y Activo, con confianza.\nSi tiene menos de 10 transacciones debe enviar prueba que el nombre de la cuenta bancarias debe ser el mismo en localbitcoin e igualmete enviarme su cedula o pasaporte.\n\nObligatoro eviar capture de la copia de transferencias"
	r.VolumeCoefficientBtc = "1.50"
	r.BankName = "⚡️ Bco. CARIBE - BANCARIBE ⚡️2.400.000 Maximo"
	r.TradeType = "ONLINE_SELL"
	r.AdID = 1086349
	r.TempPrice = "560000000.00"
	r.PaymentWindowMinutes = 90
	r.MinAmount = "1400000"
	r.LimitToFiatAmounts = ""
	r.RequireTrustedByAdvertiser = false
	r.TempPriceUsd = "7594.77"
	r.Latitud = 0
	r.Visible = true
	r.CreatedAt = "2019-10-21T11:05:37+00:00"
	r.AtmModel = ""
	//profile
	r.AdsProfile.Username = "iliotest"
	r.AdsProfile.FeedbackScore = 100
	r.AdsProfile.TradeCount = "3000+"
	r.AdsProfile.LastOnline = "2020-03-28T23:21:39+00:00"
	r.AdsProfile.Name = "iliotest (3000+; 100%)"
	return
}

func getJSONFromFile() (r LocalbitcoinsResponse, e error) {
	jsonFile, err := os.Open("localbitcoinsResponseTest.json")
	if err != nil {
		e = errors.New("The file localbitcoinsResposeTest.json could not be loaded")
		return
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		e = errors.New("The data from the file localbitcoinsResposeTest.json could not be readed")
		return
	}

	err = json.Unmarshal(jsonData, &r)
	if err != nil {
		e = errors.New("The json from the file localbitcoinsResposeTest.json could not be unmarshal")
		return
	}
	return
}

func TestLocalbitcoinsResponse(t *testing.T) {

	localbitcoinsResponse, err := getJSONFromFile()

	if err != nil {
		t.Error(err)
	}

	firstAdd := createDummyDataAdsInfo()

	if reflect.DeepEqual(firstAdd, localbitcoinsResponse.Data.Adlist[0].AdInfo) != true {
		t.Error("The first element of the list is not the same as the dummy data created")
	}

	if localbitcoinsResponse.Data.AdCount != 50 {
		t.Error("The number of elements is not right")
	}

}

func TestSearchFunctionByAmountOfMoney(t *testing.T) {

	const errorText = "The item returned in the search is not the one that need the test. Need %v, returned %v"
	localbitcoinsResponse, err := getJSONFromFile()

	if err != nil {
		t.Error(err)
	}

	//First search
	response, err := localbitcoinsResponse.Data.SearchByAmountAndBankFirstMatch(400000, "BaNeScO")
	if err != nil {
		t.Error(err)
	}

	const adIDtest1 = 1058164

	if response.AdInfo.AdID != adIDtest1 {
		t.Errorf(errorText, adIDtest1, response.AdInfo.AdID)
	}

	//Second search
	response, err = localbitcoinsResponse.Data.SearchByAmountAndBankFirstMatch(100, "")
	if err != nil {
		t.Error(err)
	}

	const adIDtest2 = 1075144

	if response.AdInfo.AdID != adIDtest2 {
		t.Errorf(errorText, adIDtest2, response.AdInfo.AdID)
	}

	//Last search
	_, err = localbitcoinsResponse.Data.SearchByAmountAndBankFirstMatch(100, "BOD")
	if err == nil {
		t.Error("It supossed that the item could be find")
	}

}

func TestSearchFunctionByAmountOfBTC(t *testing.T) {

	localbitcoinsResponse, err := getJSONFromFile()

	if err != nil {
		t.Error(err)
	}

	//First search
	response, err := localbitcoinsResponse.Data.SearchByBTCAndBankFirstMatch(0.0007141103746847872, "BaNeScO")
	if err != nil {
		t.Error(err)
	}

	const adIDtest1 = 1058164
	const errorText = "The item returned in the search is not the one that need the test. Need %v, returned %v"

	if response.AdInfo.AdID != adIDtest1 {
		t.Errorf(errorText, adIDtest1, response.AdInfo.AdID)
	}

	//Second search
	response, err = localbitcoinsResponse.Data.SearchByBTCAndBankFirstMatch(1.7798656166588295e-07, "")
	if err != nil {
		t.Error(err)
	}

	const adIDtest2 = 1075144

	if response.AdInfo.AdID != adIDtest2 {
		t.Errorf(errorText, adIDtest2, response.AdInfo.AdID)
	}
}
