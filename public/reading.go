package public

// ReadingRequest is the POST request sent to the API.
type ReadingRequest struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
