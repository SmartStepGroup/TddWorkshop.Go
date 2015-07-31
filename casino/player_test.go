package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerTest struct {
	suite.Suite
	Player *Player
}

func Test_StartPlayerSuite(t *testing.T) {
	suite.Run(t, &GameTest{})
}

func (self *PlayerTest) SetupTest() {
	self.Player = NewPlayer()
}

func (self *PlayerTest) TestBuyChips_NewPlayer_Success() {
	err := self.Player.BuyChips(1)

	self.Nil(err)
	self.Equal(1, self.Player.AvailableChips())
}

func (self *PlayerTest) TestBuyChips_PlayerHasChips_Success() {
	self.Player.BuyChips(1)

	self.Player.BuyChips(1)

	self.Equal(1+1, self.Player.AvailableChips())
}

func (self *PlayerTest) Test_BuyChips_ZeroChips_NotAllowed() {
	err := self.Player.BuyChips(0)

	self.NotNil(err)
	self.Equal("Please buy positive amount", err.Error())
}

func (self *PlayerTest) TestBuyNegativeChips_NewPlayer_Fail2() {
	err := self.Player.BuyChips(-1)

	self.NotNil(err)
	self.Equal("Please buy positive amount", err.Error())
}

func (self *PlayerTest) TestBet_Success() {
	self.Player.BuyChips(10)

	err := self.Player.Bet(Bet{Score: 1, Amount: 1})

	self.Nil(err)
}

func (self *PlayerTest) TestLeave_MakeBet_ReturnChipsOk() {
	self.Player.BuyChips(10)
	self.Player.Bet(Bet{Score: 1, Amount: 1})
	self.Player.Join(NewRollDiceGame(new(Dice)))

	self.Player.Leave()

	self.Equal(10, self.Player.AvailableChips())
}

func (self *PlayerTest) TestBet_BuyChips_AvailableChipsOk() {
	self.Player.BuyChips(10)

	self.Player.Bet(Bet{Score: 1, Amount: 1})

	self.Equal(9, self.Player.AvailableChips())
}

func (self *PlayerTest) TestBetNotAllowedScore_Fail() {
	self.Player.BuyChips(10)

	err := self.Player.Bet(Bet{Score: 7, Amount: 1})

	self.NotNil(err)
	self.Equal("Bets on 1..6 only are allowed", err.Error())
}

func (self *PlayerTest) TestBetNotAllowedScore_Fail2() {
	self.Player.BuyChips(10)

	err := self.Player.Bet(Bet{Score: 0, Amount: 1})

	self.NotNil(err)
	self.Equal("Bets on 1..6 only are allowed", err.Error())
}

func (self *PlayerTest) TestBetWrongAmount_Fail2() {
	self.Player.BuyChips(10)

	err := self.Player.Bet(Bet{Score: 1, Amount: 11})

	self.NotNil(err)
	self.Equal("Unable to bet chips more than available", err.Error())
}

func (self *PlayerTest) TestBetWrongAmount_Fail3() {
	self.Player.BuyChips(10)

	err := self.Player.Bet(Bet{Score: 1, Amount: -1})

	self.NotNil(err)
}

func TestPlayer_Create_Success(t *testing.T) {
	player := NewPlayer()

	assert.NotNil(t, player)
}

func TestPlayer_Create_NotInGame(t *testing.T) {
	player := NewPlayer()

	assert.Nil(t, player.currentGame)
	assert.Equal(t, false, player.IsInGame())
}

func TestPlayer_Create_HasNoChips(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, 0, player.AvailableChips())
}

func TestPlayer_Create_HasNoBets(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, map[int]int{}, player.bets)
}

func TestPlayer_NotInGame_Join_Success(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame(new(Dice))

	err := player.Join(game)

	assert.Nil(t, err)
	assert.True(t, player.IsInGame())
}

func TestPlayer_InGame_Join_Fail(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame(new(Dice))
	player.Join(game)

	err := player.Join(game)

	if assert.NotNil(t, err) {
		return
	}
	assert.Equal(t, "Unable to join another game", err.Error())
}

func TestGame_NoPlayers_JoinPlayer_Success(t *testing.T) {
	game := NewRollDiceGame(new(Dice))
	player := NewPlayer()

	player.Join(game)

	_, exists := game.players[player]
	assert.Equal(t, true, exists)
}

func TestGame_HasPlayer_Leave_Success(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame(new(Dice))
	player.Join(game)

	player.Leave()

	_, exists := game.players[player]
	assert.Equal(t, false, exists)
}

func TestPlayer_InGame_Leave_Success(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame(new(Dice))
	player.Join(game)

	err := player.Leave()

	assert.Nil(t, err)
	assert.Equal(t, false, player.IsInGame())
}

func TestPlayer_NotInGame_Leave_Fails(t *testing.T) {
	player := NewPlayer()

	err := player.Leave()

	assert.NotNil(t, err)
	assert.Equal(t, "Unable to leave the game before joining", err.Error())
}

func TestPlayer_AvailableChips_Success(t *testing.T) {
	player := NewPlayer()

	player.BuyChips(1)

	assert.Equal(t, 1, player.AvailableChips())
}
