package tools

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestFetchLocation(t *testing.T) {

	loadEnv()

	type args struct {
		searchText string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "seoul", args: args{searchText: "seoul"}},
		{name: "incheon", args: args{searchText: "incheon"}},
		{name: "london", args: args{searchText: "london"}},
	}

	results := []struct {
		latitude  float64
		longitude float64
		placeName string
	}{
		{37.58333, 127, "Seoul, South Korea"},
		{37.46389, 126.64861, "Incheon, South Korea"},
		{51.507321899999994, -0.12764739999999997, "London, Greater London, England, United Kingdom"},
	}

	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			latitude, longitude, placeName, err := FetchLocation(tt.args.searchText)
			if err != nil {
				return
			}
			assertions.Equalf(results[idx].latitude, latitude, "latitude should be %v", results[idx].latitude)
			assertions.Equalf(results[idx].longitude, longitude, "longitude should be %v", results[idx].longitude)
			assertions.Equalf(results[idx].placeName, placeName, "placeName should be %v", results[idx].placeName)
		})
	}
}
