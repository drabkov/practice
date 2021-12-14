package mymeteo

type Weather struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`
}

func (w *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.Main.Temp, w.Main.TempMin, w.Main.TempMax
}

func (w *Weather) GetCloudiness() (description string) {
	return w.Weather[0].Description
}

func (w *Weather) GetHumidity() (humidity int) {
	return int(w.Main.Humidity)
}

func (w *Weather) GetWind() (speed int, gust int, direction string) {

	derection := ""
	deg := w.Wind.Deg
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

	return int(w.Wind.Speed), gust, derection
}
