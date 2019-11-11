package main

import (
	"fmt"
	"practice3/mybot"
	"strings"
)

func main() {
	var language string

	for !(strings.ToLower(language) == "english" || strings.ToLower(language) == "russian") {
		fmt.Println("Type english or russian")
		fmt.Print("Language: ")
		_, err := fmt.Scanln(&language)
		if err != nil {
			fmt.Println(err)
		}
	}

	listener(createBot(strings.ToLower(language)))
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
		default:
			if myBot.GetLanguage() == "english" {
				fmt.Println("Type 1 or 2 or 3 or 4 or 5 or hello or time or date or day or bye")
			} else {
				fmt.Println("Введите 1 или 2 или 3 или 4 или 5 или привет или время или дата или день или пока")
			}
		}
	}
}

func createBot(language string) mybot.Bot {
	var myBot mybot.Bot

	switch strings.ToLower(language) {
	case "english":
		myBot = &mybot.EnglishBot{"Jhon", language}
	case "russian":
		myBot = &mybot.RussianBot{"Иван", language}
	}
	return myBot
}
