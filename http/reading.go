package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/essce/tempapi/public"
)

// ListReadings returns all the readings of the sensor
func (h *Handler) ListReadings(w http.ResponseWriter, r *http.Request) {
	return
}

// InsertReading inserts the reading from the request to the data store
func (h *Handler) InsertReading(w http.ResponseWriter, r *http.Request) {
	var req public.ReadingRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeJSONData(ctx, w, struct {
			Message string `json:"message"`
		}{
			Message: fmt.Sprintf("invalid request: %s", err.Error()),
		}, 400)
		return
	}

	id, err := h.ReadingStore.InsertReading(ctx, req.Temperature, req.Humidity)
	if err != nil {
		h.writeJSONData(ctx, w, struct {
			Message string `json:"message"`
		}{
			Message: fmt.Sprintf("internal server error: %s", err.Error()),
		}, 500)
		return
	}

	h.writeJSONData(ctx, w, struct {
		ID string `json:"id"`
	}{
		ID: id,
	}, 200)
}
