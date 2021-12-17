package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"
)

const (
	accountSID        = "accc2c661a0-3a41-360e-8417-3b4fdec4f980"
	accountToken      = "aut079a61c0-406d-39cf-9b1e-9321128f8f73"
	accountTokenWrong = "aut00000000-0000-0000-0000-000000000000"

	routineCount = 20
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	do(&wg, accountSID, accountToken, 1)

	for i := 0; i < routineCount; i++ {
		wg.Add(1)
		go do(&wg, accountSID, accountTokenWrong, 50)
	}

	wg.Wait()
}

func do(wg *sync.WaitGroup, accSID string, accToken string, reqCount int) {
	defer wg.Done()

	url := "https://api-voicebot.apifonica.com/v1/campaign"

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)

	if err != nil {
		log.Printf("http.NewRequest err = %s\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(accSID, accToken)

	for i := 0; i < reqCount; i++ {
		request(req, accToken)
	}
}

func request(req *http.Request, accToken string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err = %s\n", err)
		return
	}

	if accToken == accountToken {
		if resp.StatusCode != 206 {
			log.Printf("%s - %+v wrong\n", accToken, resp)
		}
	} else if resp.StatusCode != 403 {
		log.Printf("%d - %s - %+v wrong\n", resp.StatusCode, accToken, resp)
	}

	log.Print("+")

	defer func() {
		_ = resp.Body.Close()
	}()
}
