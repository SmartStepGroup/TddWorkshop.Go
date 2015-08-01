package casino

import "errors"

// Player model for casino players
type Player struct {
	inTheGame bool
	coins     int
}

// CanJoinGame check if player can join to game
func (player *Player) CanJoinGame() bool {
	return !player.inTheGame
}

// CanLeaveGame check if player can join to game
func (player *Player) CanLeaveGame() bool {
	return !player.CanJoinGame()
}

func (player *Player) Leave() error {
	if !player.CanLeaveGame() {
		return errors.New("You cannot leave from the game")
	}
	player.inTheGame = false
	return nil
}

func (player *Player) Join(game *Game) error {
	return game.Add(player)
}

func (player *Player) BuyCoin(i int) error {
	player.coins += i
	return nil
}

func (player *Player) Coins() int {
	return player.coins
}

func (player *Player) Bet(bet Bet) error {
	if bet.Coins > player.Coins() {
		return errors.New("player have not enouch coins for this bet")
	}
	player.coins -= bet.Coins
	return nil
}
