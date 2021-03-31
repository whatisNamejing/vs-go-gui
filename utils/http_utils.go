package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJson(url string, jsonBody []byte) string {
	reader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		panic(err)
	}
	request.Header.Set("content-Type", "application/json")
	client := http.Client{Timeout: 5 * time.Second}
	resp, erro := client.Do(request)
	if erro != nil {
		panic(erro)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)

}

func PostJsonWithToken(url string, jsonBody []byte, token string) string {
	reader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		panic(err)
	}
	request.Header.Set("content-Type", "application/json")
	request.Header.Set("WDGAuthToken", token)
	client := http.Client{Timeout: 5 * time.Second}
	resp, erro := client.Do(request)
	if erro != nil {
		panic(erro)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)

}
