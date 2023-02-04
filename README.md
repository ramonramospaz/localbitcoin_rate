# LOCALBITCOIN RATE

Localbitcoin rate is a command-line app that searches the best rate from a specific amount from one currency to another.

## Build

If you want to build the program from the source you need to install [golang](https://golang.org/dl/) version 1.16 first. Then you need to run the follow commands:

```bash
$make clean
$make build
```

## Help

Using the flag help shows all the commands that can do the app.

eg.
```bash
$localbitoin_rate --help

Usage: localbitoin_rate <command>

A shell-like app that search for rates in localbitcoin

Flags:
  --help    Show context-sensitive help.

Commands:
  currencies    Show the list of currencies permitted.
  example       Show some examples of how to fill the params.
  search        Search the rate of a specific exchange.
  version       Version of the app.

Run "localbitoin_rate <command> --help" for more information on a command
```

## Example 1:
```bash
localbitcoin_rate search -b PAB -s VED -a 100
```
This example search the rate from PAB (Panama Currency) to VED (Venezuelan Currency) from the amount 100$
## Example 2:
```bash
localbitcoin_rate search -b PAB -1 Mercantil -s VED -2 Mercantil -a 100
```
This example search the rate from PAB (Panama Currency) of the Bank Banesco to VED (Venezuelan Currency) of the Bank BOD from the amount 100$
