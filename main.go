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
		if flag == "-v" {
			fmt.Println("Version:", Version)
			os.Exit(0)
		}
		if flag == "-s" {
			stringArg = true
		}
	}

	client := godrudge.NewClient()
	err := client.Parse()
	if err != nil {
		fmt.Println("Error parsing", err)
	}
	client.PrintDrudge(stringArg)
}
