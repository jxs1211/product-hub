package main

import (
	"fmt"
	"io" // Import the io package
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials" // Import the credentials package
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Define the function to download a file from S3
func DownloadFromS3Bucket(bucket string, item string, filePath string, displayProgressBar bool) error {
	// Create the file at the specified path
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	accessKey := ""
	secretKey := ""
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	// Create a new session with the AWS configuration
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds, //credentials.AnonymousCredentials,
	})

	// Create a downloader with the session
	downloader := s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 6
	})

	// Initialize a progress bar if required
	var writer io.WriterAt = file
	if displayProgressBar {
		// Initialize the progress bar
		// (details of initializing the progress bar are provided in the source)
	}

	// Download the file from S3
	numBytes, err := downloader.Download(writer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		return err
	}

	// Finish the progress bar if it was displayed
	if displayProgressBar {
		// Finish the progress bar
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return nil
}

func doFunc(){
	// Example usage of the DownloadFromS3Bucket function
	bucket := "wave-loadtest-dev"
	item := "test-project/Tests/01fd6b43-71fc-7463-80c8-eed7767f37c1/client_logs/zipped/logs-i-0715ba3015485d430.zip"
	filePath := "./script.zip"
	displayProgressBar := true

	err := DownloadFromS3Bucket(bucket, item, filePath, displayProgressBar)
	if err != nil {
		fmt.Println("Error downloading file:", err)
	} else {
		fmt.Println("File downloaded successfully")
	}

}

func main() {
	doFunc()
}
