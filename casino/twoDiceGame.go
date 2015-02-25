package casino

import (
	"errors"
)

type TwoDiceGame struct {
	players []*Player
}

func (game *TwoDiceGame) addPlayer(player *Player) {
	game.players = append(game.players, player)
}

func (game *TwoDiceGame) removePlayer(player *Player) {
	index := 0
	for i, p := range game.players {
		if p == player {
			index = i
		}
	}
	game.players = append(game.players[:index], game.players[index+1:]...)
}

func (game *TwoDiceGame) HasPlayer(player *Player) bool {
	for _, p := range game.players {
		if p == player {
			return true
		}
	}
	return false
}

func (game *TwoDiceGame) Validate(score Score) error {
	if score < 2 || 12 < score {
		return errors.New("Bet only to numbers 2-12")
	}
	return nil
}
