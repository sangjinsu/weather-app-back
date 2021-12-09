package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchLocation(t *testing.T) {
	LoadEnv()

	type args struct {
		searchText string
	}
	tests := []struct {
		name          string
		args          args
		wantLatitude  float64
		wantLongitude float64
		wantPlaceName string
		wantErr       error
	}{
		{"seoul", args{searchText: "seoul"},
			37.58333, 127,
			"Seoul, South Korea", nil},
		{"incheon", args{searchText: "incheon"},
			37.46389, 126.64861,
			"Incheon, South Korea", nil},
		{"london", args{searchText: "london"},
			51.507321899999994, -0.12764739999999997,
			"London, Greater London, England, United Kingdom", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			gotLatitude, gotLongitude, gotPlaceName, err := FetchLocation(tt.args.searchText)
			assertions.ErrorIsf(err, nil, "FetchLocation(%v)", tt.args.searchText)
			assertions.Equalf(tt.wantLatitude, gotLatitude, "FetchLocation(%v)", tt.args.searchText)
			assertions.Equalf(tt.wantLongitude, gotLongitude, "FetchLocation(%v)", tt.args.searchText)
			assertions.Equalf(tt.wantPlaceName, gotPlaceName, "FetchLocation(%v)", tt.args.searchText)
		})
	}
}
