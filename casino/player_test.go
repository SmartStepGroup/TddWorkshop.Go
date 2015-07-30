package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayerNotInGame_Leave_Fail(t *testing.T) {
	player := NewPlayer()

	err := player.Leave()

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Unable to leave the game before joining", err.Error())
}

func TestNewPlayerInGame_Leave_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.Leave()

	assert.Nil(t, err, "Player error is not null")
}

func TestNewPlayer_IsInGame_Fail(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, false, player.IsInGame())
}
func TestPlayerInGame_Join_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	assert.Equal(t, true, player.IsInGame())
}

func TestPlayerInGame_BuyChipsWithInvalidValue_Fail(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.BuyChips(-1)

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Please buy positive amount", err.Error())
}

func TestPlayerInGame_BuyChipsWithValidValue_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.BuyChips(22)

	assert.Nil(t, err, "Player error is not null")
	assert.Equal(t, 22, player.AvailableChips())
}

func TestNewPlayer_AvailableChipsIsNull(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, 0, player.AvailableChips())
}

func TestPlayerInGame_BetAmountMoreThanAvailable_Fail(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())
	player.BuyChips(1)

	err := player.Bet(Bet{Amount: 2, Score: 2})

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Unable to bet chips more than available", err.Error())

}

func TestPlayerInGame_BetScoreNotValid_Fail(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())
	player.BuyChips(1)

	err := player.Bet(Bet{Amount: 1, Score: 7})

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Bets on 1..6 only are allowed", err.Error())
}

func TestPlayerInGame_BetScoreValid_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())
	player.BuyChips(9)

	_ = player.Bet(Bet{Amount: 2, Score: 5})

	assert.Equal(t, 7, player.AvailableChips())
	assert.Equal(t, 2, player.GetBetOn(5))
}
