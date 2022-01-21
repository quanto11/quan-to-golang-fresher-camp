package main

import (
	"fmt"
	"time"
)

func main() {
	nUrls := 1000
	nCrawler := 5
	bfChan := make(chan int, nUrls)
	done := make(chan bool)
	for i := 1; i <= nUrls; i++ {
		bfChan <- i
	}
	close(bfChan)
	for i := 1; i <= nCrawler; i++ {
		go func(j int) {
			for i := range bfChan {
				time.Sleep(time.Millisecond)
				fmt.Println(j, " crawl ", i)
			}
			done <- true
		}(i)
	}
	for i := 1; i <= nCrawler; i++ {
		<-done
	}
}
