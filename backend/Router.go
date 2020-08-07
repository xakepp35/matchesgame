package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/xakepp35/matchesgame/backend/controller"
)

func initHandlers(pool *pgxpool.Pool) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/GitRevision",
		func(w http.ResponseWriter, r *http.Request) {
			controller.GitRevision(pool, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/NewGame",
		func(w http.ResponseWriter, r *http.Request) {
			controller.NewGame(pool, w, r)
		}).Methods("PUT")

	r.HandleFunc("/api/v1/LoadGame/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			controller.LoadGame(pool, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/MakeTurn",
		func(w http.ResponseWriter, r *http.Request) {
			controller.MakeTurn(pool, w, r)
		}).Methods("POST")

	r.HandleFunc("/api/v1/TopScore",
		func(w http.ResponseWriter, r *http.Request) {
			controller.TopScore(pool, w, r)
		}).Methods("GET")

	return r
}

