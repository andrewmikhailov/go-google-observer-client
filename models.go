package go_google_observer_client

import "time"

type Event struct {
	Hash string    `json:"hash"`
	Text string    `json:"text"`
	Uri  string    `json:"uri"`
	Date time.Time `json:date`
}
