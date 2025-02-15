package main

import (
	"fmt"
	"os"

	"github.com/MGuitar24/godrudge"
)

// Version variable (will be set via -ldflags)
var Version = "dev" // Default value for local builds

func main() {
	stringArg := false
	if len(os.Args) > 1 {
		flag := os.Args[1]
		switch flag {
		case "-v":
			fmt.Println("Version:", Version)
			os.Exit(0)
		case "-s":
			stringArg = true
		case "-h":
			printHelp()
			os.Exit(0)
		}
	}

	client := godrudge.NewClient()
	err := client.Parse()
	if err != nil {
		fmt.Println("Error parsing", err)
	}
	client.PrintDrudge(stringArg)
}

func printHelp() {
	fmt.Println("Usage: drudge [options]")
	fmt.Println("Options:")
	fmt.Println("  -v        Print the version and exit")
	fmt.Println("  -s        Print the output as a string")
	fmt.Println("  -h        Print this help menu and exit")
}