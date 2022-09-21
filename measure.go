package speedy

import "context"

type SpeedTestResults struct {
	UploadSpeed   float64
	DownloadSpeed float64
}

type provider string

const (
	FastCom      provider = "fast.com"
	SpeedtestNet provider = "speedtest.net"
)

func Measure(ctx context.Context, opts ...Option) (*SpeedTestResults, error) {
	options := &Options{
		measureFn:     measureWithSpeedTestNet,
		measureUpload: true,
	}
	for _, o := range opts {
		o(options)
	}
	return options.measureFn(ctx, *options)
}

type measureFn func(ctx context.Context, options Options) (*SpeedTestResults, error)

type Options struct {
	measureFn     measureFn
	measureUpload bool
}

type Option func(*Options)

func WithProvider(provider provider) Option {
	return func(options *Options) {
		switch provider {
		case FastCom:
			options.measureFn = measureWithFastCom
		case SpeedtestNet:
			options.measureFn = measureWithSpeedTestNet
		}
	}
}

func WithDownloadSpeedOnly() Option {
	return func(options *Options) {
		options.measureUpload = false
	}
}
