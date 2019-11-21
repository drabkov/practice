package mymeteo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Meteorologist struct {
}

func (m *Meteorologist) WeatherForecast(city string) *Weather {

	url := "http://api.openweathermap.org/data/2.5/weather?" + "q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	t := dat["main"].(map[string]interface{})
	w := dat["weather"].([]interface{})
	w0 := w[0].(map[string]interface{})
	wind := dat["wind"].(map[string]interface{})
	sys := dat["sys"].(map[string]interface{})

	gust := 0
	if wind["gust"] != nil {
		gust = int(wind["gust"].(float64))
	}

	derection := ""
	deg := wind["deg"].(float64)
	switch {
	case deg == 0 || deg == 360:
		derection = "северный"
	case deg > 0 || deg < 90:
		derection = "северовосточный"
	case deg == 90:
		derection = "восточный"
	case deg > 90 || deg < 180:
		derection = "юго-восточный"
	case deg == 180:
		derection = "южный"
	case deg > 180 || deg < 270:
		derection = "юго-западный"
	case deg == 270:
		derection = "западный"
	case deg > 270 || deg < 360:
		derection = "северо-западный"
	}

	weather := &Weather{
		Temp:        t["temp"].(float64),
		TempMin:     t["temp_min"].(float64),
		TempMax:     t["temp_max"].(float64),
		Description: w0["description"].(string),
		Humidity:    int(t["humidity"].(float64)),
		Speed:       int(wind["speed"].(float64)),
		Gust:        gust,
		Direction:   derection,
		Sunrise:     int64(sys["sunrise"].(float64)),
		Sunset:      int64(sys["sunset"].(float64)),
	}

	return weather
}

type Weather struct {
	Temp        float64 `json:"temp"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Description string  `json:""`
	Humidity    int     `json:"humidity"`
	Speed       int     `json:"speed"`
	Gust        int     `json:"gust"`
	Direction   string  `json:"direction"`
	Sunrise     int64   `json:"sunrise"`
	Sunset      int64   `json:"sunset"`
}

func (w *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.Temp, w.TempMin, w.TempMax
}

func (w *Weather) GetCloudiness() (description string) {
	return w.Description
}

func (w *Weather) GetHumidity() (humidity int) {
	return w.Humidity
}

func (w *Weather) GetWind() (speed int, gust int, direction string) {
	return w.Speed, w.Gust, w.Direction
}
