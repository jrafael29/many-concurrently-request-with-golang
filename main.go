package main

import (
	"io"
	"fmt"
	"net/http"
	_ "time"
)

func makeRequest(url string, channel chan string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    bodyString := string(bodyBytes)
		channel <- bodyString
	}
}

func main(){

	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go makeRequest("https://api.adviceslip.com/advice", ch)
	}

	var items []string
	for i := range ch {
		// items = append(items, i);
		fmt.Println(i, ch)
	}
	// fmt.Println(items)
}