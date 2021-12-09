package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Location struct {
	Type        string    `json:"type"`
	Query       []string  `json:"query"`
	Attribution string    `json:"attribution"`
	Features    []Feature `json:"features"`
}

type Feature struct {
	PlaceName string    `json:"place_name"`
	Center    []float64 `json:"center"`
}

func FetchLocation(searchText string) (latitude float64, longitude float64, placeName string, err error) {
	encode := url.QueryEscape(searchText)
	params := url.Values{}
	params.Add("access_token", os.Getenv("ACCESS_TOKEN"))

	locationURL :=
		fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?%s", encode, params.Encode())

	res, err := http.Get(locationURL)
	if err != nil {
		return 0, 0, "", fmt.Errorf("cannot get location %v", err)
	}
	defer res.Body.Close()

	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return -1, -1, "", fmt.Errorf("cannot read body %v", err)
	}
	var location Location
	unmarshalErr := json.Unmarshal(data, &location)
	if unmarshalErr != nil {
		return -1, -1, "", fmt.Errorf("error in unmarshalling %v", err)
	}
	if len(location.Features) == 0 {
		return -1, -1, "", errors.New("Unable to find location. Try another search")
	}

	latitude = location.Features[0].Center[1]
	longitude = location.Features[0].Center[0]
	placeName = location.Features[0].PlaceName

	err = nil
	return
}
