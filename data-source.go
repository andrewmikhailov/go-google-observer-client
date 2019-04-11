package google_observer_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DataSource struct {
	Uri string
}

func (this DataSource) Load() ([]Event, error) {
	response, error := http.Get(this.Uri)
	if error != nil {
		return nil, error
	}

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var items []Event
	json.Unmarshal(body, &items)
	return items, nil
}

func NewDataSource(uri string) DataSource {
	return DataSource{Uri: uri}
}
