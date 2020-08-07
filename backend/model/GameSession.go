package model

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// GameSession интерфейс модели игры
type GameSession interface {

	// GetID возвращает уникальный ид сессии
	GetID() uuid.UUID

	// IsMovePossible желаемое количество должно быть положительным, не превышать текущего количества и не превишыть максимума на ход
	IsMovePossible(numMatchesToTake int) bool

	// IsGameFinished игра закончена если на столе нет спичек
	IsGameFinished() bool

	// PerformMove совершает ход в рамках сессии
	PerformMove(numMatches int) (bool, error)

	// CalculateBotTurn вычисляет сколько спичек взять боту
	CalculateBotTurn() int

	// DbSelect грузит из бд в модель
	DbSelect(conn pgx.Conn, sessionID uuid.UUID) error

	// DbInsert создаёт новую запись в бд
	DbInsert(conn pgx.Conn) error

	// обновляет данные после хода
	DbUpdate(conn pgx.Conn) error
}

// NewSesion фабрика сессии
func NewSesion(userName string, maxToTake int, initialAmount int) GameSession {
	generatedUUID, err := uuid.NewRandom()
	if err != nil {
		return nil
	}
	return &gameSessionModel{
		SessionID:     generatedUUID,
		UserName:      userName,
		StartDate:     time.Now(),
		MaxToTake:     maxToTake,
		InitialAmount: initialAmount,
		MatchesLeft:   initialAmount,
	}
}

// LoadSession загрузить игру
func LoadSession(conn pgx.Conn, sessionID uuid.UUID) GameSession {
	var model gameSessionModel
	err := model.DbSelect(conn, sessionID)
	if err != nil {
		return nil
	}
	return &model
}

// gameSessionModel бизнес моделька игровой сессии
type gameSessionModel struct {
	SessionID     uuid.UUID // уникальный id сессии
	UserName      string    // имя игрока для таблицы результатов
	StartDate     time.Time // начало матча
	EndDate       time.Time // окончание матча
	MaxToTake     int       // максимальное кол-во спичек которые игрок может взять за ход
	InitialAmount int       // начальное количество спичек в начале матча (больше для истории, для реплея)
	MatchesLeft   int       // сколько спичек сейчас на столе в рамках текущей игровой сессии
	TurnHistory   []int     // запись истории матча, для потомков :)
}

func (s *gameSessionModel) GetID() uuid.UUID {
	return s.SessionID
}

func (s *gameSessionModel) IsMovePossible(numMatchesToTake int) bool {
	return numMatchesToTake > 0 && numMatchesToTake <= s.MatchesLeft && numMatchesToTake <= s.MaxToTake
}

func (s *gameSessionModel) IsGameFinished() bool {
	return s.MatchesLeft == 0
}

func (s *gameSessionModel) CalculateBotTurn() int {
	// сколько взять чтоб загнать игрока в кольцо вычетов?
	desiredAmount := (s.MatchesLeft - 1) % (s.MaxToTake + 1)
	// в С с троичным оператором конструкция будет короче: return max(1,(left-1)%(take+1))
	// в го делаю для повышения читабельности кода
	if desiredAmount > 0 {
		return desiredAmount
	}

	return 1 // или random.. бот проигрывает
}

func (s *gameSessionModel) PerformMove(numMatches int) (bool, error) {
	if !s.IsMovePossible(numMatches) {
		return false, errors.New("Impossible move")
	}

	s.TurnHistory = append(s.TurnHistory, numMatches)
	s.MatchesLeft -= numMatches

	return s.IsGameFinished(), nil
}

func (s *gameSessionModel) DbSelect(conn pgx.Conn, sessionID uuid.UUID) error {
	return conn.QueryRow(
		context.Background(),
		`SELECT * FROM gamesession
		WHERE sessionid=$1`,
		sessionID,
	).Scan(&s.SessionID, &s.UserName, &s.StartDate, &s.EndDate, &s.MaxToTake, &s.InitialAmount, &s.MatchesLeft, &s.TurnHistory)
}

func (s *gameSessionModel) DbInsert(conn pgx.Conn) error {
	var sessionID uuid.UUID
	return conn.QueryRow(
		context.Background(),
		`INSERT INTO gamesession
		(sessionid, username, startdate, maxtotake, initialamount, matchesleft, turnhistory)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING sessionid`,
		&s.SessionID, &s.UserName, &s.StartDate, &s.MaxToTake, &s.InitialAmount, &s.MatchesLeft, &s.TurnHistory,
	).Scan(&sessionID)
}

func (s *gameSessionModel) DbUpdate(conn pgx.Conn) error {
	var sessionID uuid.UUID
	return conn.QueryRow(
		context.Background(),
		`UPDATE gamesession
		set enddate=$1, matchesleft=$2, turnhistory=$3
		WHERE sessionid=$4
		RETURNING sessionid`,
		&s.EndDate, &s.MatchesLeft, &s.TurnHistory, &s.SessionID,
	).Scan(&sessionID)
}
