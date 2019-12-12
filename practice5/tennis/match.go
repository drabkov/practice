package tennis

import (
	"fmt"
)

type Match struct {
	Player1, Player2 *Player
}

func (m *Match) Start() {

	//var wg sync.WaitGroup
	//wg.Add(2)

	ch := make(chan string)
	winner := make(chan string)

	fmt.Println("Match start.")

	go func() {
		m.Player1.Play(ch, winner)
		//wg.Done()
	}()

	go func() {
		m.Player2.Play(ch, winner)
		//wg.Done()
	}()

	ch <- "ball"

	//wg.Wait()
	fmt.Printf("Match finish. Winner is %v\n", <-winner)

}
