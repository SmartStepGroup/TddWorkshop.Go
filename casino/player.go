package casino

import "errors"

type Player struct {
	isInGame bool
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) Join() error {
	if p.isInGame {
		return errors.New("Player already in game")
	}

	p.isInGame = true
	return nil
}

func (p *Player) Leave() error {
	if !p.isInGame {
		return errors.New("Player not in game")
	}
	p.isInGame = false

	return nil
}
