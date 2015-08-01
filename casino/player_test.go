package casino
import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type GameTest struct {
	suite.Suite
}

func (test *GameTest) SetupTest() {
}

func TestInit(t *testing.T) {
	suite.Run(t, &GameTest{})
}

func (test *GameTest) Test_Player_JoinToGame_Success() {
	player := &Player{}

	player.Join()

	test.True(player.IsInGame())
}
