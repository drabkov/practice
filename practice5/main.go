package main

import "practice5/tennis"

func main() {

	p1 := &tennis.Player{"Mike", 7}
	p2 := &tennis.Player{"Phile", 5}

	m := tennis.Match{p1, p2}

	m.Start();
}