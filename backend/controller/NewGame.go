package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

type NewGameRequest struct {
	PlayerName         string
	MaxMatchesPrerTurn int
	StartMatchesAmount int
}

type NewGameResponse struct {
	SessionId string
}

// NewGame контроллер маршрута POST /NewGame
func NewGame(pool *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	requestBody := NewGameRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		SendErrorResponse(w, 400, "Json can't unmarshal request")
		return
	}
	if requestBody.PlayerName == "" {
		SendErrorResponse(w, 400, "PlayerName cannot be empty")
		return
	}
	if len(requestBody.PlayerName) > 40 {
		SendErrorResponse(w, 400, "PlayerName cannot be more than 40 chars")
		return
	}
	if requestBody.MaxMatchesPrerTurn <= 0 {
		SendErrorResponse(w, 400, "MaxMatchesPrerTurn must be positive")
		return
	}
	if requestBody.MaxMatchesPrerTurn > requestBody.StartMatchesAmount {
		SendErrorResponse(w, 400, "MaxMatchesPrerTurn must be less than StartMatchesAmount")
		return
	}
	if requestBody.StartMatchesAmount <= 0 {
		SendErrorResponse(w, 400, "StartMatchesAmount must be positive")
		return
	}
	if requestBody.StartMatchesAmount > 40 {
		SendErrorResponse(w, 400, "StartMatchesAmount must be sane (max 40 matches allowed)")
		return
	}

	// look for playername, if it alredy has session?

	// create new session in db

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		SendErrorResponse(w, 500, "Unable to acquire a database connection: "+err.Error())
		return
	}
	defer conn.Release()

	newSessionId := "testsess"
	resStruct := NewGameResponse{
		SessionId: newSessionId,
	}
	res, err := json.Marshal(&resStruct)
	if err != nil {
		SendErrorResponse(w, 500, "Json Marshaller failed")
		return
	}
	w.Write(res)
}
