package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go sendData(ch, &wg)
	go receiveData(ch, &wg)

	wg.Wait()
}

func sendData(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
	ch <- "7"
}

func receiveData(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for input := range ch {
		fmt.Println(input)
	}
}
