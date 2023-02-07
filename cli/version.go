package cli

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Change the version in the makefile
var version = "undefined"

func showVersion() {
	fmt.Printf("Cli app version %v\n", version)
}

func showExample() {
	appPath, _ := os.Executable()
	pathSeparator := "/"
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
	}
	appPathFolders := strings.Split(appPath, pathSeparator)
	appName := appPathFolders[len(appPathFolders)-1]
	fmt.Printf("Example 1:\n%v search -b PAB -s VED -a 100\nThis example search the rate from ", appName)
	fmt.Printf("PAB (Panama Currency) to VED (Venezuelan Currency) from the amount 100$\n")
	fmt.Printf("Example 2:\n%v search -b PAB -1 Banesco -s VED -2 BNC -a 100\nThis example search the rate from ", appName)
	fmt.Printf("PAB (Panama Currency) of the Bank Banesco to VED (Venezuelan Currency) of the Bank BNC from the amount 100$\n")
}
