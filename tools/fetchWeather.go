package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Weather struct {
	Current struct {
		Temperature         int      `json:"temperature"`
		WeatherIcons        []string `json:"weather_icons"`
		WeatherDescriptions []string `json:"weather_descriptions"`
		Humidity            int      `json:"humidity"`
		IsDay               string   `json:"is_day"`
	} `json:"current"`
}

func FetchWeather(latitude float64, longitude float64) (temperature int, humidity int, description string, icon string, isDay string, err error) {

	params := url.Values{}
	params.Add("access_key", os.Getenv("ACCESS_KEY"))
	params.Add("query", fmt.Sprintf("%f,%f", latitude, longitude))

	weatherURL := fmt.Sprintf("http://api.weatherstack.com/current?%s", params.Encode())

	res, err := http.Get(weatherURL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return 0, 0, "", "", "", err
	}
	var weather Weather
	err = json.Unmarshal(data, &weather)
	if err != nil {
		return 0, 0, "", "", "", err
	}

	temperature = weather.Current.Temperature
	humidity = weather.Current.Humidity
	description = weather.Current.WeatherDescriptions[0]
	icon = weather.Current.WeatherIcons[0]
	isDay = weather.Current.IsDay

	return
}
