package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	url := fmt.Sprintf("http://foo-%d.bar:8080", rand.Intn(100))

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
