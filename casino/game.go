package casino

import "errors"

type Game struct {
	playersCount int
}

func (g *Game) Add() error {
	if g.playersCount >= 6 {
		return errors.New("Game is full")
	}
	g.playersCount++

	return nil
}
