package main

import (
	"fmt"
	"net/http"
)

type Cl struct {
	client *http.Client
}

func (c *Cl) SomeRequest() {
	req, err := http.NewRequest(http.MethodGet, "https://ya.ru", nil)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	res, err := c.client.Do(req)
	defer res.Body.Close()

	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	fmt.Printf("%+v\n", res)
}
