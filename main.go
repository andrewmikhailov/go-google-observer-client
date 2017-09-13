package main

import (
	"fmt"
	"vitche.com/google-observer-client"
)

type EventFeed struct {
	key   string
	value []google_observer_client.Event
}

func main() {

	feeds := make(chan EventFeed)

	go func() {
		uri := "http://google-observer-1.herokuapp.com/api/event/list?kernelIdentifier=593a842d7c52901100c8815c"
		items := google_observer_client.NewDataSource(uri).Load()
		feeds <- EventFeed{key: uri, value: items}
	}()

	select {
	case eventFeed := <-feeds:
		for i := 1; i < len(eventFeed.value); i++ {
			var item = eventFeed.value[i]
			fmt.Printf("News event parsed: %s %s %s\n", item.Hash, item.Uri, item.Text)
		}
	}
}
