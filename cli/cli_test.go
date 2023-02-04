package cli

import (
	"testing"
)

func TestCli(t *testing.T) {
	c := cliParm{}
	err := getLocalbitcoinRate(c)
	if err == nil {
		t.Error("The parameters CoinBuy, CoinShell and Amount are obligatories, and the structure is empty")
	}
	c.Search.CoinBuy = "PAB"
	c.Search.CoinSell = "VED"
	c.Search.Amount = "100"
	err = getLocalbitcoinRate(c)
	if err != nil {
		t.Error("In this test the params are right, there cant be any error. Check your internet conection")
	}
	c.Search.CoinBuy = "PAx"
	c.Search.CoinSell = "VED"
	c.Search.Amount = "100"
	err = getLocalbitcoinRate(c)
	if err == nil {
		t.Error("The buy coin is wrong, it can not be a valid search")
	}

}
