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
	GetLanguage() string
}

type EnglishBot struct {
	Name     string
	Language string
}

type RussianBot struct {
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

func (b *RussianBot) GetLanguage() string {
	return b.Language
}

func (b *RussianBot) SayHello() {
	fmt.Printf("%s: Привет, я %s\n", b.Name, b.Name)
}

func (b *RussianBot) SayCurrentTime() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Printf("%s: Сейчас %s\n", b.Name, t.Format("15:04:02"))
}

func (b *RussianBot) SayCurrentDate() {
	var locmonth string

	t := time.Now()
	year, month, day := t.Date()

	switch month {
	case 1:
		locmonth = "января"
	case 2:
		locmonth = "февраля"
	case 3:
		locmonth = "марта"
	case 4:
		locmonth = "апреля"
	case 5:
		locmonth = "мая"
	case 6:
		locmonth = "июня"
	case 7:
		locmonth = "июня"
	case 8:
		locmonth = "июля"
	case 9:
		locmonth = "сентября"
	case 10:
		locmonth = "октября"
	case 11:
		locmonth = "ноября"
	case 12:
		locmonth = "декабря"
	}

	fmt.Printf("%s: Сегодня %v %v %v года\n", b.Name, day, locmonth, year)
}

func (b *RussianBot) SayCurrentDayOfWeek() {
	var locWeekday string

	t := time.Now()

	switch t.Weekday() {
	case 1:
		locWeekday = "Понедельник"
	case 2:
		locWeekday = "Вторник"
	case 3:
		locWeekday = "Среда"
	case 4:
		locWeekday = "Четверг"
	case 5:
		locWeekday = "Пятница"
	case 6:
		locWeekday = "Суббота"
	case 7:
		locWeekday = "Воскресенье"
	}

	fmt.Printf("%s: Сегодня %s\n", b.Name, locWeekday)
}

func (b *RussianBot) SayBye() {
	fmt.Printf("%s: Пока\n", b.Name)
}
