package tennis

import "fmt"

type Match struct {
	Player1 *Player
	Player2 *Player
}

func (m *Match) Start() {

	var winer string

	ch := make(chan string)
	ch <- ""

	fmt.Println("Match start.")

	go func() { 
		m.Player1.Play(ch) 
		winer,_ = <-ch}()

	go func() { 
		m.Player2.Play(ch) 
		winer,_ = <-ch}()

	fmt.Println("Match finish. Winer is " + winer)
}
