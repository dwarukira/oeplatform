package bucket

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "omarmarketplace"
)

func UploadFile(uploadFile []byte) (*s3.PutObjectOutput, string, error) {
	hasher := sha256.New()
	hasher.Write(uploadFile)

	key := hex.EncodeToString(hasher.Sum(nil))

	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	fmt.Println(http.DetectContentType(uploadFile))
	object, err := s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(AWS_S3_BUCKET),
		Key:           aws.String(key),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(uploadFile),
		ContentLength: aws.Int64(int64(len(uploadFile))),
		ContentType:   aws.String(http.DetectContentType(uploadFile)),
	})
	url := GetUrl(key)
	return object, url, err
}

func GetUrl(key string) string {
	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", AWS_S3_BUCKET, AWS_S3_REGION, key)
}

// func BulkUploadS3
