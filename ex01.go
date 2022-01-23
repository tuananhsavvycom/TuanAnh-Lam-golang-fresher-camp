package main

import (
	"fmt"
	"time"
)

// func crawl data
func crawlData(data int) {
	fmt.Println("Data crawled: ", data)
}

// func main
func main() {
	fmt.Print("Crawling data..\n")

	// Create number of URLs
	urlCrawl := 1000
	// Make buffered channel
	buffChan := make(chan int, urlCrawl)

	// Push data in channel
	for i := 1; i <= urlCrawl; i++ {
		buffChan <- i
	}

	//Make go func in a loop to print data
	for i := 1; i <= 5; i++ {
		go func() {
			for data := range buffChan {
				crawlData(data)
			}
		}()
	}
	time.Sleep(time.Second * 2)
}
