package s3

import (
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/davestephens/movies-r-us/rest-api/models"
	"github.com/davestephens/movies-r-us/rest-api/utils"
)

func DownloadFile(notification models.Notification) *os.File {
	bucket := notification.Bucket
    item := notification.Key

	file, err := os.CreateTemp("", "movie")
    if err != nil {
		utils.Logger.Errorf("Unable to open file %q for writing, %v", item, err)
    }

    defer file.Close()

	// init aws session
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("eu-west-1")},
    )

	// set up s3manager and download
	downloader := s3manager.NewDownloader(sess)
    numBytes, err := downloader.Download(file,
        &s3.GetObjectInput{
            Bucket: aws.String(bucket),
            Key:    aws.String(item),
        })
    if err != nil {
        utils.Logger.Errorf("Unable to download item %q, %v", item, err)
    }

    utils.Logger.Infof("Downloaded %s %d bytes", file.Name(), numBytes)

	// return handle to file
	return file
}