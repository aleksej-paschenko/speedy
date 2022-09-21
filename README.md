
# Speedy 
A small GO library that tests the download and upload speeds by using Ookla’s https://www.speedtest.net/ and Netflix’s
https://fast.com/.

# Prerequisite 

Speedy requires Chrome browser to be installed locally for testing with fast.com.

# Usage 

Here is an example of how you can use speedy:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aleksej-paschenko/speedy"
)

func main() {
	// by default, speedy measures both the download and upload speed.
	results, err := speedy.Measure(context.Background(), speedy.WithProvider(speedy.FastCom))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fast.com download speed: %v, upload speed: %v\n", results.DownloadSpeed, results.UploadSpeed)

	// but there is also a way to measure the download speed only
	results, err = speedy.Measure(context.Background(),
		speedy.WithProvider(speedy.SpeedtestNet),
		speedy.WithDownloadSpeedOnly())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("speedtest.net download speed: %v\n", results.DownloadSpeed)
}
``` 