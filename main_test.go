package main

import (
	"fmt"
	"sync"
	"testing"
	// Your path can be different
	"go-google-observer-client/src/vitche.com/google-observer-client"
)

var (
	uri                     = "http://google-observer-1.herokuapp.com/api/event/list?kernelIdentifier="
	feeds                   = make(chan EventFeed)
	items                   []google_observer_client.Event
	kernelIdentifiersLength = len(kernelIdentifiers)
	wg                      sync.WaitGroup
	kernelIdentifiers       = []string{
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
)

func getItems(kernelIdentifier string) {
	go func() {
		items := google_observer_client.NewDataSource(uri + kernelIdentifier).Load()
		feeds <- EventFeed{key: uri, value: items}
		fmt.Println("Kernel done")
		wg.Done()
	}()
}

func runItems() {
	fmt.Println("Start process kernels: ", kernelIdentifiersLength)
	wg.Add(kernelIdentifiersLength)

	for index := 0; index < kernelIdentifiersLength; index++ {
		fmt.Println("Start process kernel: ", index)
		getItems(kernelIdentifiers[index])
	}
	for index := 0; index < kernelIdentifiersLength; index++ {
		select {
		case eventfeed := <-feeds:
			items = append(items, eventfeed.value...)
		}
	}
}

func TestRun(t *testing.T) {
	fmt.Println("Start tests")

	runItems()
	wg.Wait()

	fmt.Println("Run tests")

	t.Run("TestResultCount", func(t *testing.T) {
		itemsCount := len(items)
		fmt.Println(itemsCount)
		if 0 == itemsCount {
			t.Error("result is empty, count of items:", itemsCount)
		}
	})
	t.Run("TestResultType", func(t *testing.T) {
		resultType := fmt.Sprintf("%T", items)
		if "[]google_observer_client.Event" != resultType {
			t.Error("Wxpected result type: []google_observer_client.Event, result type is:" + resultType)
		}
	})
	t.Run("TestResultItems", func(t *testing.T) {
		for i := 1; i < len(items); i++ {
			var item = items[i]
			if item.Hash == "" {
				t.Error("Hash is empty of item:", item)
			} else if item.Text == "" {
				t.Error("Text is empty of item:", item)
			} else if item.Uri == "" {
				t.Error("Uri is empty of item:", item)
			}
		}
	})
}
