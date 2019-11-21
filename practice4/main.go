package main

import (
	"fmt"
	"practice4/mymeteo"
)

func main() {
	var city string = "Mahilyow"

	fmt.Print("City: ")
	_, err := fmt.Scanln(&city)
	if err != nil {
		fmt.Println(err)
	}

	m := &mymeteo.Meteorologist{}
	w := m.WeatherForecast(city)

	temp, _, _ := w.GetTemperature()
	speed, gust, direction := w.GetWind()

	fmt.Printf("Сегодня в городе %s %s,\n", city, w.GetCloudiness())
	fmt.Printf("температура воздуха %v °С,\n", temp)
	fmt.Printf("ветер %v %v м/с с порывами до %v м/с.\n", direction, speed, gust) 
	fmt.Printf("Влажность воздуха %s.\n", string(w.GetHumidity()))
	fmt.Printf("Восход солнца %v, заход солнца %v.\n", w.Sunrise, w.Sunset)
}
