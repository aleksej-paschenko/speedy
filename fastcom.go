package speedy

import (
	"context"
	"strconv"

	"github.com/chromedp/chromedp"
)

func measureWithFastCom(ctx context.Context, options options) (*SpeedTestResults, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var download, upload string

	actions := []chromedp.Action{
		chromedp.Navigate(`https://fast.com`),
		chromedp.WaitVisible(`#show-more-details-link`),
	}
	if options.measureUpload {
		actions = append(actions,
			chromedp.Click(`#show-more-details-link`, chromedp.NodeVisible),
			chromedp.WaitVisible(`#upload-value.succeeded`),
			chromedp.Text(`#upload-value`, &upload))
	}

	actions = append(actions, chromedp.Text(`#speed-value`, &download))
	err := chromedp.Run(ctx, actions...)
	if err != nil {
		return nil, err
	}
	downloadf, err := strconv.ParseFloat(download, 64)
	if err != nil {
		return nil, err
	}
	var uploadf float64
	if options.measureUpload {
		uploadf, err = strconv.ParseFloat(upload, 64)
		if err != nil {
			return nil, err
		}
	}
	return &SpeedTestResults{DownloadSpeed: downloadf, UploadSpeed: uploadf}, nil
}
