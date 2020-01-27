package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// CallAPI : call api at url and returns json result
func CallAPI(url string) (string, error) {
	// Create HTTP client with timeout
	client := &http.Client{
		// value given in https://www.mediawiki.org/wiki/Wikidata_Query_Service/User_Manual#Wikimedia_service
		Timeout: 60 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "wdjqs https://github.com/JVillafruela/wdjqs")
	request.Header.Set("Accept", "application/json")

	// Make request
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	dataInBytes, err := ioutil.ReadAll(res.Body)
	return string(dataInBytes), err
}
