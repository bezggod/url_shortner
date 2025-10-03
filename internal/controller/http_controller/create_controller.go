package http_controller

import (
	"encoding/json"
	"net/http"
)

type reqCreate struct {
	URL string `json:"url"`
}

type respCreate struct {
	ShortURL string `json:"short_url"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req reqCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "bad reques", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	code, err := c.uc.Create(ctx, req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(respCreate{ShortURL: code})
}
