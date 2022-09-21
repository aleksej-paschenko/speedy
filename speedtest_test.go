package speedy

import (
	"context"
	"errors"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockClient struct {
	OnFetchUserInfoContext   func(ctx context.Context) (*speedtest.User, error)
	OnFetchServerListContext func(ctx context.Context, user *speedtest.User) (speedtest.Servers, error)
}

func (m mockClient) FetchUserInfoContext(ctx context.Context) (*speedtest.User, error) {
	return m.OnFetchUserInfoContext(ctx)
}

func (m mockClient) FetchServerListContext(ctx context.Context, user *speedtest.User) (speedtest.Servers, error) {
	return m.OnFetchServerListContext(ctx, user)
}

func Test_measureWithSpeedTestNetWithClient(t *testing.T) {
	tests := []struct {
		name    string
		client  *mockClient
		options options
		want    *SpeedTestResults
		wantErr string
	}{
		{
			name: "FetchUserInfo fails",
			client: &mockClient{
				OnFetchUserInfoContext: func(ctx context.Context) (*speedtest.User, error) {
					return nil, errors.New("no-user")
				},
			},
			wantErr: "no-user",
		},
		{
			name: "FetchServerList fails",
			client: &mockClient{
				OnFetchUserInfoContext: func(ctx context.Context) (*speedtest.User, error) {
					return &speedtest.User{}, nil
				},
				OnFetchServerListContext: func(ctx context.Context, user *speedtest.User) (speedtest.Servers, error) {
					return nil, errors.New("fetch-servers-failure")
				},
			},
			wantErr: "fetch-servers-failure",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := measureWithSpeedTestNetWithClient(context.Background(), tt.options, tt.client)
			if len(tt.wantErr) > 0 {
				require.NotNil(t, err)
				require.Nil(t, results)
				require.Equal(t, tt.wantErr, err.Error())
				return
			}
			require.Nil(t, err)
			require.NotNil(t, results)
		})
	}
}
