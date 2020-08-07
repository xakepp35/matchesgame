package controller

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TopScore GET
func TopScore(pool *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
}
