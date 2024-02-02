package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Player struct {
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Position string `json:"position"`
	Age      int    `json:"age"`
}

var players = []Player{
	{"LeBron James", 23, "Forward", 36},
	{"Anthony Davis", 3, "Forward/Center", 28},
	{"Russell Westbrook", 0, "Guard", 33},
	{"Dwight Howard", 39, "Center", 35},
	{"Carmelo Anthony", 7, "Forward", 37},
	{"Kent Bazemore", 9, "Guard", 32},
	{"Talen Horton-Tucker", 5, "Guard/Forward", 20},
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, player := range players {
		if player.Name == params["name"] {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.NotFound(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Real Madrid App is healthy!"))
}
