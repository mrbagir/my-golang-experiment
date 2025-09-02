package handler

import (
	"backend-supporter/backend/api/app"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/api/app", GetListApp).Methods("GET")
	r.HandleFunc("/api/app/{appname}", GetApp).Methods("GET")
	r.HandleFunc("/api/app/{appname}/settings", GetAppSettings).Methods("GET")
}

func GetListApp(w http.ResponseWriter, r *http.Request) {
	appnames := make([]string, len(app.Apps))
	for i, app_ := range app.Apps {
		appnames[i] = app_.Name
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appnames)
}

func GetApp(w http.ResponseWriter, r *http.Request) {
	appname := mux.Vars(r)["appname"]
	if appname == "" {
		respondWithError(w, http.StatusBadRequest, "App name is required")
		return
	}

	for _, app_ := range app.Apps {
		if app_.Name == appname {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(app_)
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "App not found")
}

func GetAppSettings(w http.ResponseWriter, r *http.Request) {
	appname := mux.Vars(r)["appname"]
	if appname == "" {
		respondWithError(w, http.StatusBadRequest, "App name is required")
		return
	}

	for _, app_ := range app.Apps {
		if app_.Name == appname {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(app_.Settings)
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "App not found")
}
