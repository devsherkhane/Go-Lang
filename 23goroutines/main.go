package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mt sync.Mutex

var signals = []string{
	"test",
}

var wg sync.WaitGroup

func main() {
	websites := []string{
		"https://youtube.com",
		"https://github.com",
		"https://google.com",
		"https://go.dev",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs There is an issue in endpoint!")
	} else {
		mt.Lock()
		signals = append(signals, endpoint)
		mt.Unlock()
		fmt.Printf(" %d status code of website %s \n ", res.StatusCode, endpoint)
	}
}
