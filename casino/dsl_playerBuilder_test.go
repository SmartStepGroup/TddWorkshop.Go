package casino

type PlayerBuilder struct {
	player   Player
	betChips Chips
	betScore Score
}

func (playerBuilder *PlayerBuilder) InGame() *PlayerBuilder {
	playerBuilder.player.currentGame = &Game{}
	return playerBuilder
}

func (playerBuilder *PlayerBuilder) Joined(game IGame) *PlayerBuilder {
	playerBuilder.player.Join(game)
	return playerBuilder
}

func (playerBuilder *PlayerBuilder) With(chips Chips) *PlayerBuilder {
	playerBuilder.player.Buy(chips)
	return playerBuilder
}

func (playerBuilder *PlayerBuilder) Bet(betAmount uint) *PlayerBuilder {
	playerBuilder.betChips = Chips(betAmount)
	return playerBuilder
}

func (playerBuilder *PlayerBuilder) OnScore(score uint) *PlayerBuilder {
	playerBuilder.betScore = Score(score)
	return playerBuilder
}

func (playerBuilder *PlayerBuilder) Please() *Player {
	if playerBuilder.player.Balance() == 0 {
		playerBuilder.With(Chips(1000))
	}
	if playerBuilder.betChips != 0 || playerBuilder.betScore != 0 {
		if playerBuilder.betChips == 0 {
			playerBuilder.betChips = SOME_CHIPS
		}
		if playerBuilder.betScore == 0 {
			playerBuilder.betScore = SOME_SCORE
		}
		playerBuilder.player.Bet(playerBuilder.betChips, playerBuilder.betScore)
	}
	return &playerBuilder.player
}
