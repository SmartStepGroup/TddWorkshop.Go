package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_CanJoin(t *testing.T) {
	player := Player{}
	game := Game{}

	player.Join(game)

	assert.True(t, player.isInGame)
}

func Test_PlayerInGame_CanLeave(t *testing.T) {
	player := Player{}
	game := Game{}
	player.Join(game)

	player.Leave(game)

	assert.False(t, player.isInGame)
}

func Test_PlayerNotInGame_CanNotLeave(t *testing.T) {
	player := Player{}
	game := Game{}

	err := player.Leave(game)

	assert.NotNil(t, err)
	assert.Equal(t, "You can not leave game", err.Error())
}
