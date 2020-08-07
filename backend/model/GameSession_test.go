package main

import (
	"fmt"
	"testing"
)

// TestCalculateBotTurn пример теста на валидность ходов
func TestCalculateBotTurn(t *testing.T) {
	maxToTake := 3 // for...
	for matchesLeft := 0; matchesLeft <= 30; matchesLeft++ {
		s := NewSesion("tester", maxToTake, matchesLeft)
		if s.IsGameFinished() {
			fmt.Println("игра закончена")
			continue
		}
		desiredAmount := s.CalculateBotTurn()
		fmt.Println("осталось: ", s.MatchesLeft, ". бот возьмёт: ", desiredAmount) // ez debug
		_, err := s.TakeMatch(desiredAmount)
		if err != nil {
			t.FailNow()
		}
	}
}
