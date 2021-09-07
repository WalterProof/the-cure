package controllers

import (
	"net/http"

	"pezos/views"

	"pezos/models"
)

func NewHomepage(tt models.TezTools) *Homepage {
	return &Homepage{
		HomeView: views.NewView("base", "homepage"),
		tt:       tt,
	}
}

// Homepage controller.
type Homepage struct {
	HomeView *views.View
	tt       models.TezTools
}

func (h *Homepage) Index(w http.ResponseWriter, r *http.Request) {
	xtzPrice, err := h.tt.XTZPrice()

	var vd views.Data
	vd.Yield = struct {
		XTZPrice float64
		Error    error
	}{xtzPrice, err}

	h.HomeView.Render(w, r, vd)

	return
}
