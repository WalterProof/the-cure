package controllers

import "pezos/views"

// NewStatic creates a new Static controller.
func NewStatic() *Static {
	return &Static{
		Contact: views.NewView("base", "static/contact"),
	}
}

// Static controller.
type Static struct {
	Contact *views.View
}
