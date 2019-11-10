package main

import (
	"fmt"
	"practice3/mybot"
	"strings"
)

func main() {
	var language string

	for !(strings.ToLower(language) == "english" || strings.ToLower(language) == "russian") {
		fmt.Print("Language : ")
		_, err := fmt.Scanln(&language)
		fmt.Println(err)
	}

	myBot := createBot(language)

	listener(myBot)
}

func listener(myBot mybot.Bot) {
	var req string

	for strings.ToLower(req) != "5" {
		fmt.Print("You : ")
		_, err := fmt.Scanln(&req)
		fmt.Println(err)

		switch {
		case strings.ToLower(req) == "1":
			myBot.SayHello()
		case strings.ToLower(req) == "2":
			myBot.SayCurrentTime()
		case strings.ToLower(req) == "3":
			myBot.SayCurrentDate()
		case strings.ToLower(req) == "4":
			myBot.SayCurrentDayOfWeek()
		case strings.ToLower(req) == "5":
			myBot.SayBye()
		}
	}
}

func createBot(language string) mybot.Bot {
	var myBot mybot.Bot

	switch strings.ToLower(language) {
	case "english":
		myBot = &mybot.EnglishBot{"Jhon"}
	case "russian":
		myBot = &mybot.RussianBot{"Иван"}
	}
	return myBot
}
