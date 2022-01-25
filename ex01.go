package main

import (
	"fmt"
	"time"
)


func crawlData(data int) {
	fmt.Println("Data crawled: ", data)
}


func main() {
	fmt.Print("Crawling data..\n")	
	urlCrawl := 1000
	buffChan := make(chan int, urlCrawl)
	for i := 1; i <= urlCrawl; i++ {
		buffChan <- i
	}
	for i := 1; i <= 5; i++ {
		go func() {
			for data := range buffChan {
				crawlData(data)
			}
		}()
	}
	time.Sleep(time.Second * 2)
}
