package controller

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

// LoadGame  GET /LoadGame/:id
func LoadGame(pool *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
}
