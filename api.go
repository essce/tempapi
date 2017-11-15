package tempapi

import "context"

// Reading is a type that stores temperature and humidity readings.
type Reading struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	AddedAt     string  `json:"added_at"`
}

// ReadingStore is where the readings are going to be stored.
type ReadingStore interface {
	InsertReading(ctx context.Context, temp, humidity float64) (string, error)
	ListReadings(context.Context) ([]Reading, error)
}
