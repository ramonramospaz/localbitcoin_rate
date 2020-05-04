package conexion

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"localbitcoin_rate/entity"
	"os"
	"strings"
	"testing"
)

func getJSONFromFileDummy(url string) (r entity.LocalbitcoinsResponse, e error) {
	if strings.Contains(url, "https://localbitcoins.com/es/buy-bitcoins-online/PAB/.json") {
		r, _ = getJSONFromFile("localbitcoinsPABBuyTest.json")
	} else if strings.Contains(url, "https://localbitcoins.com/es/sell-bitcoins-online/VES/.json") {
		r, _ = getJSONFromFile("localbitcoinsVESSellTest.json")
	} else {
		e = errors.New("The dummy files cant be loaded")
	}

	return
}

func getJSONFromFile(filename string) (r entity.LocalbitcoinsResponse, e error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		e = errors.New("The file localbitcoinsResposeTest.json could not be loaded")
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
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

func TestGetLocalbitcoinResponse(t *testing.T) {
	url := "https://localbitcoins.com/es/buy-bitcoins-online/pab/.json"
	_, err := getLocalbitcoinResponse(url)

	if err != nil {
		t.Error(err)
	}

	//Descoment if you want to see the output
	//output, _ := json.MarshalIndent(&response, "", "\t\t")
	//t.Log(string(output))
}

func BenchmarkGetLocalbitcoinResponse(b *testing.B) {
	url := "https://localbitcoins.com/es/buy-bitcoins-online/pab/.json"
	_, err := getLocalbitcoinResponse(url)

	if err != nil {
		b.Error(err)
	}

}

func TestGetLocalbitcoinCurrencieResponse(t *testing.T) {
	response, err := getLocalbitcoinCurrencieResponse(urlCoins)

	if err != nil {
		t.Error(err)
	}

	if len(response.Data.Currencies) == 0 {
		t.Error("There are not currency code active, please check localbitcoin")
	}

	//Descoment if you want to see the output
	//output, _ := json.MarshalIndent(&response, "", "\t\t")
	//t.Log(string(output))
}

func TestGetLocalbitcoinResume(t *testing.T) {
	response, err := getLocalbitcoinRate("PAB", "Banesco", "VES", "Banesco", 105, getJSONFromFileDummy)

	if err != nil {
		t.Error(err)
	}

	output, _ := json.MarshalIndent(&response, "", "\t\t")
	t.Log(string(output))

	response, err = getLocalbitcoinRate("PAB", "Banesco", "VES", "BOD", 50, getJSONFromFileDummy)

	if err != nil {
		t.Error(err)
	}

	output, _ = json.MarshalIndent(&response, "", "\t\t")
	t.Log(string(output))

	response, err = getLocalbitcoinRate("PAX", "Banesco", "VES", "BOD", 50, getJSONFromFileDummy)

	if err == nil {
		t.Error("The currency PAX dont exist")
	}

	response, err = getLocalbitcoinRate("PAB", "Banesco", "VES", "PUPIS", 50, getJSONFromFileDummy)

	if err == nil {
		t.Error("There is not bank PUPIS")
	}

	response, err = getLocalbitcoinRate("PAB", "", "VES", "", 50, getJSONFromFileDummy)

	if err != nil {
		t.Error(err)
	}

	output, _ = json.MarshalIndent(&response, "", "\t\t")
	t.Log(string(output))

	response, err = getLocalbitcoinRate("PAB", "", "VES", "", -50, getJSONFromFileDummy)

	if err == nil {
		t.Error("There arent adds with negative numbers")
	}

}

func Test1(t *testing.T) {
	response, err := getLocalbitcoinRate("PAB", "", "VES", "", 50, getLocalbitcoinResponse)

	if err != nil {
		t.Error(err)
	}

	output, _ := json.MarshalIndent(&response, "", "\t\t")
	t.Log(string(output))
}
