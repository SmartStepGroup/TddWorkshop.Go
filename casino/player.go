package casino

// Player model for casino players
type Player struct{}

// CanJoinGame check if player can join to game
func (player *Player) CanJoinGame() bool {
	return true
}
