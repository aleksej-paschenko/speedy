package speedy

import "context"

// SpeedTestResults contains measure results of the download and upload speeds.
type SpeedTestResults struct {
	// UploadSpeed in Mbps.
	UploadSpeed float64
	// DownloadSpeed in Mbps.
	DownloadSpeed float64
}

type provider string

const (
	FastCom      provider = "fast.com"
	SpeedtestNet provider = "speedtest.net"
)

// Measure measures the download and upload speeds with a specific provider.
func Measure(ctx context.Context, opts ...option) (*SpeedTestResults, error) {
	options := &options{
		measureFn:     measureWithSpeedTestNet,
		measureUpload: true,
	}
	for _, o := range opts {
		o(options)
	}
	return options.measureFn(ctx, *options)
}

type measureFn func(ctx context.Context, options options) (*SpeedTestResults, error)

type options struct {
	measureFn     measureFn
	measureUpload bool
}

type option func(*options)

// WithProvider configures a specific provider to be used.
func WithProvider(provider provider) option {
	return func(options *options) {
		switch provider {
		case FastCom:
			options.measureFn = measureWithFastCom
		case SpeedtestNet:
			options.measureFn = measureWithSpeedTestNet
		}
	}
}

// WithDownloadSpeedOnly configures to measure the download speed only.
func WithDownloadSpeedOnly() option {
	return func(options *options) {
		options.measureUpload = false
	}
}
