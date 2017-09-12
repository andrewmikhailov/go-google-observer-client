package main

import (
	"fmt"
	"vitche.com/google-observer-client"
)

func main() {
	uri := "http://google-observer-1.herokuapp.com/api/event/list?kernelIdentifier=593a842d7c52901100c8815c"
	items := google_observer_client.NewDataSource(uri).Load()
	for i := 1; i < len(items); i++ {
		var item = items[i]
		fmt.Printf("News event parsed: %s %s %s\n", item.Hash, item.Uri, item.Text)
	}
}
