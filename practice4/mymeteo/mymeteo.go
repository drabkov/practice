package mymeteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Meteorologist struct {
}

type Weather struct {
	temp        float64
	tempMin     float64
	tempMax     float64
	description string
	humidity    int
	speed       int
	gust        int
	direction   string
}

func (m *Meteorologist) WeatherForecast(city string) *Weather {

	url := "http://api.openweathermap.org/data/2.5/weather?" + "q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	weather := &Weather{}
	json.Unmarshal(body, weather)

	return weather
}

func (w *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.temp, w.tempMin, w.tempMax
}

func (w *Weather) GetCloudiness() (description string) {
	return w.description
}

func (w *Weather) GetHumidity() (humidity int) {
	return w.humidity
}

func (w *Weather) GetWind() (speed int, gust int, direction string) {
	return w.speed, w.gust, w.direction
}
