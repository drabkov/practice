package tennis

import "fmt"

type Match struct {
	Player1 *Player
	Player2 *Player
}

func (m *Match) Start() {

	var win string

	ch := make(chan int)

	fmt.Println("Match start.")
	go func() { win = m.Player1.Play(ch) }()

	go func() { win = m.Player1.Play(ch) }()

	fmt.Println(win)
	fmt.Println("Match finish.")

}
