package casino

type Game struct{}

func (self *Game) Add(player Player) error {
	return nil
}

func (self *Game) Remove(player Player) {

}

func (self *Game) Has(player Player) bool {
	return false
}
