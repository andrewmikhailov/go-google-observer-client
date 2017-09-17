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

	var kernelIdentifiers = []string{
		"59822e1c5f0c8f1100bfd7ef",
		"593fd40f981c98110027ac06",
		"593a842d7c52901100c8815c",
		"593839d4a81c151100ec9d13",
		"593839a5a81c151100ec9d10",
		"5938395ea81c151100ec9d0d",
		"593838b7a81c151100ec9d09",
		"59383864a81c151100ec9cfa",
		"593837f4a81c151100ec9cf5",
		"59383794a81c151100ec9cf2",
		"5925a5b82a37d71100415e5f",
		"55a7d9d7c66f3111008e442a",
		"55a7d9eac66f3111008e442b",
		"55a7d9f8c66f3111008e442c",
		"55a7da01c66f3111008e442d",
		"55a8cbe55945261100e9354d",
		"55ccc39a6675e91100163eb8",
		"55ccc3d76675e91100163ebb",
		"55ccc43c6675e91100163ebe",
		"55ccc47d6675e91100163ec1",
		"55ccc4d86675e91100163ec4",
		"55ccc5376675e91100163ec7",
	}

	feeds := make(chan EventFeed)

	for i := 0; i < len(kernelIdentifiers); i++ {
		var kernelIdentifier = kernelIdentifiers[i]
		go func() {
			uri := "http://google-observer-1.herokuapp.com/api/event/list?kernelIdentifier=" + kernelIdentifier
			items := google_observer_client.NewDataSource(uri).Load()
			feeds <- EventFeed{key: uri, value: items}
		}()
	}

	for i := 0; i < len(kernelIdentifiers); i++ {
		select {
		case eventFeed := <-feeds:
			for i := 1; i < len(eventFeed.value); i++ {
				var item = eventFeed.value[i]
				fmt.Printf("News event parsed: %s %s %s\n", item.Hash, item.Uri, item.Text)
			}
		}
	}
}
