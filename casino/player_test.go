package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingNewPlayer(t *testing.T) {
	player := NewPlayer()

	assert.NotNil(t, player)

	assert.Nil(t, player.currentGame)
	assert.Equal(t, 0, player.availableChips)
	assert.Equal(t, map[int]int{}, player.bets)
}

func TestJoinGame(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame()

	err := player.Join(game)

	assert.Nil(t, err)
	assert.Equal(t, game, player.currentGame)

	_, exists := game.players[player]
	assert.Equal(t, true, exists)
}

func TestJoinGameFail(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame()

	err := player.Join(game)
	assert.Nil(t, err)

	err = player.Join(game)
	assert.NotNil(t, err)
	assert.Equal(t, "Unable to join another game", err.Error())
}

func TestLeaveGame(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame()

	err := player.Join(game)
	assert.Nil(t, err)

	err = player.Leave()
	assert.Nil(t, err)
	assert.Nil(t, player.currentGame)

	_, exists := game.players[player]
	assert.Equal(t, false, exists)
}

func TestLeaveGameFail(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame()

	err := player.Join(game)
	assert.Nil(t, err)

	err = player.Leave()
	assert.Nil(t, err)

	err = player.Leave()
	assert.NotNil(t, err)
	assert.Equal(t, "Unable to leave the game before joining", err.Error())
}
