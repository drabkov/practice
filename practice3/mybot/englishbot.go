package mybot

import (
	"fmt"
	"time"
)

type EnglishBot struct {
	Name     string
	Language string
}

func (b *EnglishBot) GetLanguage() string {
	return b.Language
}

func (b *EnglishBot) SayHello() {
	fmt.Printf("%s: Hello, I am %s\n", b.Name, b.Name)
}

func (b *EnglishBot) SayCurrentTime() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Printf("%s: Now is %s\n", b.Name, t.Format("15:04:02"))
}

func (b *EnglishBot) SayCurrentDate() {
	t := time.Now()
	year, month, day := t.Date()
	fmt.Printf("%s: Today is %v %v, %v\n", b.Name, month, day, year)
}

func (b *EnglishBot) SayCurrentDayOfWeek() {
	fmt.Printf("%s: Today is %s\n", b.Name, time.Now().Weekday())
}

func (b *EnglishBot) SayBye() {
	fmt.Printf("%s: Bye\n", b.Name)
}
