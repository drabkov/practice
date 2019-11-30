package main

import (
	"fmt"
	"practice4/mymeteo"
	"time"
)

func main() {
	var city string

	fmt.Print("City: ")
	_, err := fmt.Scanln(&city)
	if err != nil {
		fmt.Println(err)
	}

	m := &mymeteo.Meteorologist{}
	w := m.WeatherForecast(city)

	t, _, _ := w.GetTemperature()
	speed, gust, direction := w.GetWind()

	fmt.Printf("Сегодня в городе %s %s,", city, w.GetCloudiness())
	fmt.Printf("температура воздуха %v °С,\n", t)
	fmt.Printf("ветер %v %v м/с с порывами до %v м/с. \n", direction, speed, gust) 
	fmt.Printf("Влажность воздуха %d %%.\n", w.GetHumidity())
	fmt.Printf("Восход солнца %v, заход солнца %v.\n", time.Unix(w.Sys.Sunrise, 0).Format("15:04"), time.Unix(w.Sys.Sunset, 0).Format("15:04"))

	dw := m.DailyForecast(city, 7)

	fmt.Println("Изменение температуры в ", city, " на следующие 5 дней.")
	for v := range dw.List {
		fmt.Println(v)
		//fmt.Println(time.Unix(v.Dt, 0).Format("01.01.2001"), " - " +  v.Main.Temp)	
	}
}
