package speedy

import (
	"context"
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

type speedtestClient interface {
	FetchUserInfoContext(ctx context.Context) (*speedtest.User, error)
	FetchServerListContext(ctx context.Context, user *speedtest.User) (speedtest.Servers, error)
}

func measureWithSpeedTestNet(ctx context.Context, options options) (*SpeedTestResults, error) {
	return measureWithSpeedTestNetWithClient(ctx, options, speedtest.New())
}

func measureWithSpeedTestNetWithClient(ctx context.Context, options options, client speedtestClient) (*SpeedTestResults, error) {
	user, err := client.FetchUserInfoContext(ctx)
	if err != nil {
		return nil, err
	}
	serverList, err := client.FetchServerListContext(ctx, user)
	if err != nil {
		return nil, err
	}
	if len(serverList) <= 0 {
		return nil, fmt.Errorf("no servers")
	}
	server := serverList[0]
	if err := server.DownloadTestContext(ctx, false); err != nil {
		return nil, err
	}
	if options.measureUpload {
		if err := server.UploadTestContext(ctx, false); err != nil {
			return nil, err
		}
	}
	return &SpeedTestResults{
		UploadSpeed:   server.ULSpeed,
		DownloadSpeed: server.DLSpeed,
	}, nil
}
