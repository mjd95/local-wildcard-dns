package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func main() {
	// Choose a random URL under the ".bar" subdomain
	url := fmt.Sprintf("http://foo-%d.bar:8080", rand.Intn(100))
	//url := "http://foo.bar:8080"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
