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
		if err != nil {
			fmt.Println(err)
		}
	}

	listener(createBot(language))
}

func listener(myBot mybot.Bot) {
	var req string

	for !(strings.ToLower(req) == "5" || strings.ToLower(req) == "пока" || strings.ToLower(req) == "bye") {
		fmt.Print("You: ")
		_, err := fmt.Scanln(&req)
		if err != nil {
			fmt.Println(err)
		}

		switch strings.ToLower(req) {
		case "1", "привет", "hello":
			myBot.SayHello()
		case "2", "время", "time":
			myBot.SayCurrentTime()
		case "3", "дата", "date":
			myBot.SayCurrentDate()
		case "4", "день", "day":
			myBot.SayCurrentDayOfWeek()
		case "5", "пока", "bye":
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
