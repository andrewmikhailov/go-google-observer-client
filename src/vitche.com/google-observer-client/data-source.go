package google_observer_client

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type DataSource struct {
	Uri string
}

func (this DataSource) Load() [] Event {
	response, error := http.Get(this.Uri)
	if error != nil {
		panic(error.Error())
	}

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error.Error())
	}

	var items [] Event
	json.Unmarshal(body, &items)
	return items;
}

func NewDataSource(uri string) DataSource {
	return DataSource{Uri: uri}
}
