package mybot

import (
	"fmt"
	"time"
)

type Bot interface {
	SayHello()
	SayCurrentTime()
	SayCurrentDate()
	SayCurrentDayOfWeek()
	SayBye()
}

type EnglishBot struct {
	Name string
}

type RussianBot struct {
	Name string
}

func (b *EnglishBot) SayHello() {
	fmt.Println("Hello, I am " + b.Name)
}

func (b *EnglishBot) SayCurrentTime() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Printf(b.Name, "%s: Now is", t.Format("15:04:02"))
}

func (b *EnglishBot) SayCurrentDate() {
	fmt.Println("Today is ")
}

func (b *EnglishBot) SayCurrentDayOfWeek() {
	fmt.Println("Today is ")
}

func (b *EnglishBot) SayBye() {
	fmt.Println("Bye")
}

func (b *RussianBot) SayHello() {
	fmt.Println("Привет, я  " + b.Name)
}

func (b *RussianBot) SayCurrentTime() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Сейчас", t.Format("15:04"))
}

func (b *RussianBot) SayCurrentDate() {
	fmt.Println("Сегодня ")
}

func (b *RussianBot) SayCurrentDayOfWeek() {
	fmt.Println("Сегодня ")
}

func (b *RussianBot) SayBye() {
	fmt.Println("Пока")
}
