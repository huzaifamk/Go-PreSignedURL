package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetPresignedURL(sess *session.Session, bucket, key *string) (string, error) {
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func main() {

	bucket := flag.String("b", "testbucket5u3920", "The bucket")
	key := flag.String("k", "audio.mp3", "The object key")
	flag.Parse()

	if *bucket == "" || *key == "" {
		fmt.Println("You must supply a bucket name (-b BUCKET) and object key (-k KEY)")
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println("Could not access the region!")
		return
	}

	urlStr, err := GetPresignedURL(sess, bucket, key)
	if err != nil {
		fmt.Println("Got an error retrieving a presigned URL:")
		fmt.Println(err)
		return
	}

	fmt.Println("The presigned URL: " + urlStr + " is valid for 15 minutes")
}
