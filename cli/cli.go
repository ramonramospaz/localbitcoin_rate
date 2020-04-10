package cli

import (
	"encoding/json"
	"localbitcoin_rate/conexion"
	"strconv"

	"gopkg.in/validator.v2"

	"fmt"

	"github.com/alecthomas/kong"
)

//CLI is the client interface
type cliParm struct {
	Search struct {
		CoinBuy  string `check:"len=3,regexp=[A-Z]"  required short:"b" help:"Currency code for Buying." `
		BankBuy  string `check:"max=10,regexp=[a-zA-Z]*" optional short:"1" help:"Bank name for Buying." `
		CoinSell string `check:"len=3,regexp=[A-Z]" required short:"s" help:"Currency code for Selling." `
		BankSell string `check:"max=10,regexp=[a-zA-Z]*" optional short:"2" help:"Bank name for Selling." `
		Amount   string `required short:"a" help:"Amount to exchange."`
	} `cmd help:"Search the rate of a specific exchange."`
}

// Cli function that check the param that need the cli interface
func Cli() {
	c := cliParm{}
	ctx := kong.Parse(&c)

	switch ctx.Command() {
	case "search":
		getLocalbitcoinRate(c)
	default:
		fmt.Println("Do something")
		//panic(ctx.Command())
	}

}

func getLocalbitcoinRate(c cliParm) (errs error) {

	if errs = validator.WithTag("check").Validate(c); errs != nil {
		fmt.Printf("There are some errors in the input\n")
		return
	}

	amount, errFloat := strconv.ParseFloat(c.Search.Amount, 64)
	if c.Search.CoinBuy != "" && c.Search.CoinSell != "" && errFloat == nil {
		fmt.Printf("Starting to search ads and rate from currency %v to currency %v with the amount %v\nThis could take some time.....\n",
			c.Search.CoinBuy, c.Search.CoinSell, amount)
		response, err := conexion.GetLocalbitcoinRate(c.Search.CoinBuy, c.Search.BankBuy, c.Search.CoinSell, c.Search.BankSell, amount)
		if err == nil {
			output, _ := json.MarshalIndent(&response, "", "\t\t")
			fmt.Printf("%v\n", string(output))
		} else {
			fmt.Printf("%v\n", err)
			errs = err
		}
	} else {
		fmt.Println("The parameters CoinBuy, CoinShell and Amount are obligatories, use --help param")
		errs = errFloat
	}
	return

}
