package cli

import (
	"encoding/json"
	"localbitcoin_rate/conexion"
	"localbitcoin_rate/entity"
	"strconv"

	"golang.org/x/text/message"

	"gopkg.in/validator.v2"

	"fmt"

	"github.com/alecthomas/kong"
)

//CLI is the client interface
type cliParm struct {
	Currencies struct{} `cmd help:"Show the list of currencies permitted."`
	Example    struct{} `cmd help:"Show some examples of how to fill the params."`
	Search     struct {
		CoinBuy  string `check:"len=3,regexp=[A-Z]"  required short:"b" help:"Currency code for Buying. (REQUIRED)" `
		BankBuy  string `check:"max=20,regexp=[a-zA-Z]*" optional short:"1" help:"Bank name for Buying. " `
		CoinSell string `check:"len=3,regexp=[A-Z]" required short:"s" help:"Currency code for Selling. (REQUIRED)" `
		BankSell string `check:"max=20,regexp=[a-zA-Z]*" optional short:"2" help:"Bank name for Selling." `
		Amount   string `required short:"a" help:"Amount to exchange. (REQUIRED)"`
	} `cmd help:"Search the rate of a specific exchange."`
	Version struct{} `cmd help:"Version of the app."`
}

// Cli function that check the param that need the cli interface
func Cli() {
	c := cliParm{}
	ctx := kong.Parse(&c,
		kong.Name("localbitoin_rate"),
		kong.Description("A shell-like app that search for rates in localbitcoin"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	switch ctx.Command() {
	case "currencies":
		getListCurrencies()
	case "example":
		getExample()
	case "search":
		getLocalbitcoinRate(c)
	case "version":
		getVersion()
	default:
		panic(ctx.Command())
	}

}

func getLocalbitcoinRate(c cliParm) (errs error) {

	if errs = validator.WithTag("check").Validate(c); errs != nil {
		fmt.Printf("There are some errors in the input %v\n", errs)
		return
	}

	amount, errFloat := strconv.ParseFloat(c.Search.Amount, 64)
	if c.Search.CoinBuy != "" && c.Search.CoinSell != "" && errFloat == nil {
		fmt.Printf("Starting to search ads and rate from currency %v to currency %v with the amount %v\nThis could take some time.....\n",
			c.Search.CoinBuy, c.Search.CoinSell, amount)
		response, err := conexion.GetLocalbitcoinRate(c.Search.CoinBuy, c.Search.BankBuy, c.Search.CoinSell, c.Search.BankSell, amount)
		if err == nil {
			showInformation(response)
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

func showInformation(info entity.LocalbitcoinRateInformation) {

	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println("")
	fmt.Printf("%v    %v\n", greenText("CurrencyBuy:"), info.CurrencyBuy)
	fmt.Printf("%v    %v\n", greenText("BankNameBuy:"), info.BankNameBuy)
	p.Printf("%v      %.2f\n", greenText("AmountBuy:"), info.AmountBuy)
	fmt.Printf("%v  %v\n", greenText("PublicViewBuy:"), blueText(info.PublicViewBuy))
	fmt.Printf("%v   %v\n", greenText("CurrencySell:"), info.CurrencySell)
	fmt.Printf("%v   %v\n", greenText("BankNameSell:"), info.BankNameSell)
	p.Printf("%v     %.2f\n", greenText("AmountSell:"), info.AmountSell)
	fmt.Printf("%v %v\n", greenText("PublicViewSell:"), blueText(info.PublicViewSell))
	p.Printf("%v           %.2f\n", greenText("Rate:"), info.Rate)
	fmt.Println("")
}

func greenText(text string) string {
	return fmt.Sprintf("\033[1;32m%v\033[0m", text)
}

func blueText(text string) string {
	return fmt.Sprintf("\033[1;34m%v\033[0m", text)
}

func getListCurrencies() {
	listCoins := conexion.ShowListCoins()
	output, _ := json.MarshalIndent(&listCoins.Data.Currencies, "", "\t\t")
	fmt.Printf("%v\n", string(output))
}
