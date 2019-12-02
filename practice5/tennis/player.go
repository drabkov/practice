package tennis

import (
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Skill int
}

func (p *Player) Play(ch chan int) string {
	var win string

	ball, ok := <-ch

	if !ok {
		win = "Win player " + p.Name
	}

	if p.Skill < rand.Intn(10) {
		close(ch)
	} else {
		ch <- ball
	}

	return win
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
