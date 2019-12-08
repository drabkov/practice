package tennis

import (
	"fmt"
	"sync"
)

type Match struct {
	Player1 *Player
	Player2 *Player
}

func (m *Match) Start() {

	var wg sync.WaitGroup
	wg.Add(3)

	ch := make(chan string)
	go func() {
		ch <- "ball"
		wg.Done()
	}()

	fmt.Println("Match start.")

	go func() {
		m.Player1.Play(ch)
		wg.Done()
	}()

	go func() {
		m.Player2.Play(ch)
		wg.Done()
	}()

	wg.Wait()
	//fmt.Printf("Match finish. Winner is %v\n", <-ch)
}
