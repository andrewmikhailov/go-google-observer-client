package main

import (
	"fmt"
	"testing"

	// Your path can be different
	"go-google-observer-client/src/vitche.com/google-observer-client"
)

var (
	uri   = "http://google-observer-1.herokuapp.com/api/event/list?kernelIdentifier=593a842d7c52901100c8815c"
	items = google_observer_client.NewDataSource(uri).Load()
)

func TestResultCount(t *testing.T) {
	itemsCount := len(items)
	if 0 == itemsCount {
		t.Error("result is empty, count of items: %i", itemsCount)
	}
}

func TestResultType(t *testing.T) {
	resultType := fmt.Sprintf("%T", items)
	if "[]google_observer_client.Event" != resultType {
		t.Error("Wxpected result type: []google_observer_client.Event, result type is: %s", resultType)
	}
}

func TestResultItems(t *testing.T) {
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
}
