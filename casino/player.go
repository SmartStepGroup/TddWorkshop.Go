package casino

import (
	"errors"
)

type Player struct {
	currentGame *Game
	balance     Chips
	currentBet  Bet
}

func (player *Player) Join(game *Game) error {
	if game == nil {
		return errors.New("Joining nil game is not allowed")
	}
	if game.HasPlayer(player) {
		return errors.New("Player is already in game")
	}
	if player.currentGame != nil {
		return errors.New("Player is already in game")
	}

	game.addPlayer(player)
	player.currentGame = game
	return nil
}

func (player *Player) Leave() error {
	if player.currentGame == nil {
		return errors.New("Please join the game before leaving")
	}
	if !player.currentGame.HasPlayer(player) {
		return errors.New("Please join the game before leaving")
	}

	player.currentGame.removePlayer(player)
	player.currentGame = nil
	return nil
}

func (player *Player) Balance() Chips {
	return player.balance
}

func (player *Player) Buy(chips Chips) error {
	if chips < 0 {
		return errors.New("Buying negative chips is not allowed")
	}
	player.balance += chips
	return nil
}

func (self *Player) Bet(chips Chips, score Score) error {
	if score < 1 || 6 < score {
		return errors.New("Bet only to numbers 1-6")
	}

	self.currentBet = Bet{Chips: chips, Score: score}
	self.balance -= chips
	return nil
}

func (self *Player) CurrentBet() Bet {
	return self.currentBet
}
