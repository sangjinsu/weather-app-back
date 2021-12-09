package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchWeather(t *testing.T) {

	LoadEnv()

	type args struct {
		latitude  float64
		longitude float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"seoul", args{37.58333, 127}, nil},
		{"incheon", args{37.46389, 126.64861}, nil},
		{"london", args{51.507321899999994, -0.12764739999999997}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			_, _, _, _, _, err := FetchWeather(tt.args.latitude, tt.args.longitude)
			assertions.ErrorIsf(err, nil, "FetchWeather(%v, %v)", tt.args.latitude, tt.args.longitude)
		})
	}
}
