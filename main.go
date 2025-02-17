package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mikegetz/godrudge"
	"github.com/mmcdole/gofeed"
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
		case "-w":
			if len(os.Args) > 2 {
				refreshTime, _ := strconv.Atoi(os.Args[2])
				startWatch(refreshTime)
			} else {
				startWatch(30)
			}
			select {}
		}
	}
	runAndPrint(stringArg)
}

func runAndPrint(stringArg bool) {
	client := godrudge.NewClient()
	err := client.ParseRSS()
	if err != nil {
		fmt.Println("Error parsing", err)
	}
	client.PrintDrudge(stringArg)
}

func startWatch(refreshInterval int) {
	const rssURL = "http://feeds.feedburner.com/DrudgeReportFeed"
	var latestPublishedTime *time.Time
	go func() {
		for {
			fp := gofeed.NewParser()
			feed, err := fp.ParseURL(rssURL)
			if err != nil {
				log.Println("Error fetching RSS feed:", err)
			} else {
				if len(feed.Items) > 0 {
					// Only print new items
					for _, item := range feed.Items {
						if item.PublishedParsed != nil && (latestPublishedTime == nil || item.PublishedParsed.After(*latestPublishedTime)) {
							runAndPrint(false)
							if latestPublishedTime == nil {
								fmt.Println("Waiting for update from Drudge...")
							} else {
								currentTime := time.Now().Format("Mon, 02 Jan 2006 03:04:05 PM MST")
								fmt.Println(currentTime)
							}
							latestPublishedTime = item.PublishedParsed
							break
						}
					}
				}
			}
			if refreshInterval > 30 {
				time.Sleep(time.Duration(refreshInterval) * time.Second)
			} else {
				time.Sleep(30 * time.Second)
			}
		}
	}()
}

func printHelp() {
	fmt.Println("Usage: drudge [options]")
	fmt.Println("Options:")
	fmt.Println("  -v        Print the version and exit")
	fmt.Println("  -s        Print the output as a string")
	fmt.Println("  -h        Print this help menu and exit")
	fmt.Println("  -w [n]    Watch for updates every n seconds (default is 30 seconds)")
}
