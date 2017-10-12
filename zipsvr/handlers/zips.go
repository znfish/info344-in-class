package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/znfish/info344-in-class/zipsvr/models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):] // start and go through the end
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest) // DO NOT use log.fatal
		return
	}
	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips) // encode into json
}
