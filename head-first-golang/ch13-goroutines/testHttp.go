package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Url  string
	Size int
}

func responseSize(url string, ch chan Page) {
	fmt.Println("Getting:", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	ch <- Page{Url: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	resources := []string{
		"https://qa.myservicecloud.net",
		"https://dev.myservicecloud.net",
		"https://myservicecloud.net",
		"https://qa-new.myservicecloud.net",
		"https://ya.ru",
		"https://google.com",
	}

	for _, url := range resources {
		go responseSize(url, pages)
	}

	for i := 0; i < len(resources); i++ {
		fmt.Printf("%#v\n", <-pages)
	}
}
