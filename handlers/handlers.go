package handlers

import (
	"net/http"

	"github.com/nassimelhormi/adserver/models"
)

// AdServerHandler func
func AdServerHandler(data []models.Campaign) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
