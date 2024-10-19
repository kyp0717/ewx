package handlers

import (
	"github.com/kyp0717/ewx/internal/middleware"
	"github.com/kyp0717/ewx/internal/store"
	"github.com/kyp0717/ewx/internal/templates"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if !ok {
		c := templates.GuestIndex()

		err := templates.Layout(c, "My website").Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		return
	}

	c := templates.Index(user.Email)
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
