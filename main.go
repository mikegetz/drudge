package main

import (
	"fmt"
	"os"

	"github.com/MGuitar24/godrudge"
)

// Version variable (will be set via -ldflags)
var Version = "dev" // Default value for local builds

func main() {
	if len(os.Args) > 1 {
		flag := os.Args[1]
		if flag == "-v" {
			fmt.Println("Version:", Version)
			os.Exit(0)
		}
	}

	client := godrudge.NewClient()
	err := client.Parse()
	if err != nil {
		fmt.Println("Error parsing", err)
	}
	client.PrintDrudge()
}
