package casino_new

import (
	"math/rand"
	"time"
)

type RollDiceGame struct {
	players map[*Player]struct{}
}

func NewRollDiceGame() *RollDiceGame {
	return &RollDiceGame {
		players: make(map[*Player]struct{}),
	}
}

func (self *RollDiceGame) Play() {
	rand.Seed(time.Now().UTC().UnixNano())
	var winningScore = rand.Intn(6) + 1

	for player, _ := range self.players {
		player.Win(player.GetBetOn(winningScore)*6)
		player.Lose()
	}
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}
