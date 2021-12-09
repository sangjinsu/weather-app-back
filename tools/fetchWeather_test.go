package tools

import "testing"

func TestFetchWeather(t *testing.T) {
	type args struct {
		latitude  float64
		longitude float64
		placeName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FetchWeather(tt.args.latitude, tt.args.longitude, tt.args.placeName)
		})
	}
}
