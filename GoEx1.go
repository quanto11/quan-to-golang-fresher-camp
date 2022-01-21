package main

import (
	"fmt"
	"time"
)

func main() {
	bfChan := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		bfChan <- i
	}
	close(bfChan)
	for i := 1; i <= 5; i++ {
		go func(j int) {
			for i := range bfChan {
				time.Sleep(time.Millisecond)
				fmt.Println(j, "crawl", i)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
