package google_observer_client

type Event struct {
	Hash string `json:"hash"`
	Text string `json:"text"`
	Uri  string `json:"uri"`
}
