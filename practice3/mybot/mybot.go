package mybot

type Bot interface {
	SayHello()
	SayCurrentTime()
	SayCurrentDate()
	SayCurrentDayOfWeek()
	SayBye()
	GetLanguage() string
}
