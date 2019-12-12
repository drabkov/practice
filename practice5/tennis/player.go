package tennis

import (
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Skill int
}

func (p *Player) Play(ch chan string, winner chan string) {
	for {
		_, ok := <-ch

		if ok {
			if p.Skill < rand.Intn(10) {
				close(ch)
			} else {
				ch <- "ball"
			}
		} else {
			winner <- p.Name
		}
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
