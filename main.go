package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Event struct {
	Hash string `json:"hash"`
	Text string `json:"text"`
	Uri  string `json:"uri"`
}

func main() {

	url := "http://google-observer-1.herokuapp.com/api/event/list?kernelIdentifier=593a842d7c52901100c8815c"
	response, error := http.Get(url)
	if error != nil {
		panic(error.Error())
	}

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error.Error())
	}

	var items [] Event
	json.Unmarshal(body, &items)

	for i := 1; i < len(items); i++ {

		var item = items[i]
		fmt.Printf("News event parsed: %s %s %s\n", item.Hash, item.Uri, item.Text)
	}
}
