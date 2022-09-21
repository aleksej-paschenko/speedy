package speedy

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMeasure(t *testing.T) {
	tests := []struct {
		name    string
		opts    []Option
		want    *SpeedTestResults
		wantErr bool
	}{
		{
			opts: []Option{
				WithProvider(FastCom),
			},
		},
		{
			opts: []Option{
				WithProvider(FastCom),
				WithDownloadSpeedOnly(),
			},
		},
		{
			opts: []Option{
				WithProvider(SpeedtestNet),
			},
		},
		{
			opts: []Option{
				WithProvider(SpeedtestNet),
				WithDownloadSpeedOnly(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Measure(context.Background(), tt.opts...)
			require.Nil(t, err)
			fmt.Printf("got = %#v\n", got)
		})
	}
}
