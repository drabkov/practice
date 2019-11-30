package mymeteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Meteorologist struct {
}

func (m *Meteorologist) WeatherForecast(city string) *Weather {

	url := "http://api.openweathermap.org/data/2.5/weather?" + "q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather := &Weather{}
	if err := json.Unmarshal(body, weather); err != nil {
		fmt.Println(err)
	}

	return weather
}

func (m *Meteorologist) DailyForecast(city string, cnt int) *DailyWeather {

	url := "http://api.openweathermap.org/data/2.5/forecast?" + "q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	dw := &DailyWeather{}
	if err := json.Unmarshal(body, dw); err != nil {
		fmt.Println(err)
	}

	return dw
}
