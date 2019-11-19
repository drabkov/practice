package main

import (
	"fmt"
	"practice4/mymeteo"
)

func main() {
	var city string

	fmt.Print("City: ")
	_, err := fmt.Scanln(&city)
	if err != nil {
		fmt.Println(err)
	}

	m := &mymeteo.Meterologist{}
	w := m.WeatherForecast(city)

	fmt.Printf("Сегодня в городе %s %s,", city, w.GetCloudiness())
	/*"температура воздуха %+-%%temperature%°С,",
	"ветер %windDirection% %windSpeed%м/с с порывами до %windGust%м/с.",
	" Влажность воздуха %humidity%%.",
	" Восход солнца %sunrise%, заход солнца %sunset%.", )*/
}
