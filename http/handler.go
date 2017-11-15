package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/essce/tempapi"
)

// Handler does stuff
type Handler struct {
	ReadingStore tempapi.ReadingStore
}

func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.writeJSONData(ctx, w,
		struct {
			Version string `json:"version"`
		}{
			Version: "1.0.0",
		}, 200)
}

type jsonRes struct {
	Data interface{} `json:"data"`
}

func (h *Handler) writeJSONData(ctx context.Context, w http.ResponseWriter, data interface{}, code int) {
	res := jsonRes{Data: data}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	enc := json.NewEncoder(w)
	enc.Encode(res)
}
