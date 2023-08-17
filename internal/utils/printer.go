package utils

import (
	"fmt"
	"os"
)

const (
	usage  = `Usage: pentimento encode|decode|fit <image_file> <data_file> [password]`
	banner = `
             _   _               _       
 ___ ___ ___| |_|_|_____ ___ ___| |_ ___ 
| . | -_|   |  _| |     | -_|   |  _| . |
|  _|___|_|_|_| |_|_|_|_|___|_|_|_| |___|
|_|                                      

`
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func ExitWithManual() {
	printWithColor(banner, green)
	printWithColor(usage, white)
	os.Exit(0)
}

func PrintRed(msg string) {
	printWithColor(msg, red)
}

func PrintGreen(msg string) {
	fmt.Printf("\n%s%s%s\n", green, msg, reset)
}

func printWithColor(msg string, color string) {
	fmt.Printf("%s%s%s", color, msg, reset)
}
