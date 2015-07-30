package casino_new

import (
	"math/rand"
)

type IDice interface {
	Roll() int
}

type RollDiceGame struct {
	dice    IDice
	players map[*Player]struct{}
}

func NewRollDiceGame(dice IDice) *RollDiceGame {
	return &RollDiceGame{
		dice:    dice,
		players: make(map[*Player]struct{}),
	}
}

func (self *RollDiceGame) Play() {
	r := rand.New(rand.NewSource(99))
	var winningScore = r.Int()%6 + 1

	for player, _ := range self.players {
		player.Win(player.GetBetOn(winningScore) * 6)
		player.Lose()
	}
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}
