package client_config

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

func MinioConnect(endpoint string, accessKey string, secretKey string, secure bool) (error, minio.Client) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		log.Fatalln(err)
		return err, *minioClient
	}
	log.Println("Connection to minio established")
	return nil, *minioClient
}

func GetFileObject(bucketName string, objectName string, client minio.Client) ([]byte, error) {
	ctx := context.Background()
	contentBuffer, err := client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	contentBytes := new(bytes.Buffer)
	if _, err := io.Copy(contentBytes, contentBuffer); err != nil {
		return nil, err
	}
	log.Println("File object getting from minio done...")
	return contentBytes.Bytes(), nil
}
