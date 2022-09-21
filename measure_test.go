package speedy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMeasure(t *testing.T) {
	tests := []struct {
		name           string
		opts           []option
		wantZeroUpload bool
	}{
		{
			opts: []option{
				WithProvider(FastCom),
			},
		},
		{
			opts: []option{
				WithProvider(FastCom),
				WithDownloadSpeedOnly(),
			},
			wantZeroUpload: true,
		},
		{
			opts: []option{
				WithProvider(SpeedtestNet),
			},
		},
		{
			opts: []option{
				WithProvider(SpeedtestNet),
				WithDownloadSpeedOnly(),
			},
			wantZeroUpload: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Measure(context.Background(), tt.opts...)
			require.Nil(t, err)
			if tt.wantZeroUpload {
				require.True(t, got.UploadSpeed < 1e-9)
			}
			require.True(t, got.DownloadSpeed > 0.1)
		})
	}
}

func BenchmarkSpeedTestNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opts := options{
			measureUpload: true,
		}
		_, err := measureWithSpeedTestNet(context.Background(), opts)
		if err != nil {
			b.Error(err)
		}
	}
}
func BenchmarkFastCom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opts := options{
			measureUpload: true,
		}
		_, err := measureWithFastCom(context.Background(), opts)
		if err != nil {
			b.Error(err)
		}
	}
}
