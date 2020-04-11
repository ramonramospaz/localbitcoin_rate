package cli

import (
	"fmt"
	"os"
	"strings"
)

const version = "0.0.5"

func getVersion() {
	fmt.Printf("Cli app version %v\n", version)
}

func getExample() {
	appPath, _ := os.Executable()
	appPathFolders := strings.Split(appPath, "/")
	appName := appPathFolders[len(appPathFolders)-1]
	fmt.Printf("Example 1:\n%v search -b PAB -s VES -a 100\nThis example search the rate from ", appName)
	fmt.Printf("PAB (Panama Currency) to VES (Venezuelan Currency) from the amount 100$\n")
	fmt.Printf("Example 2:\n%v search -b PAB -1 Banesco -s VES -2 BOD -a 100\nThis example search the rate from ", appName)
	fmt.Printf("PAB (Panama Currency) of the Bank Banesco to VES (Venezuelan Currency) of the Bank BOD from the amount 100$\n")
}
