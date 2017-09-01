package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type fact struct {
	Category interface{} `json:"category"`
	IconURL  string      `json:"icon_url"`
	ID       string      `json:"id"`
	Value    string      `json:"value"`
}

var factsChannel = make(chan fact)

func main() {
	numberOfFacts := flag.Int("n", 1, "number of facts to print")
	for i := 0; i < *numberOfFacts; i++ {
		go fetchFact()
	}

	for i := 0; i < *numberOfFacts; i++ {
		fact, _ := <-factsChannel
		fmt.Println(fact.Value)
	}
	close(factsChannel)
}

func fetchFact() {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fact := fact{}
	err = json.Unmarshal(body, &fact)
	if err != nil {
		log.Fatal(err)
	}
	factsChannel <- fact
}
