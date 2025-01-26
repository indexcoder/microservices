package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Msg:   "Hi the broker service!",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}
