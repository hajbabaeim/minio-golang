package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
)

var (
	certFile string
	mode     string

	rootCmd = &cobra.Command{
		Use:   "parse",
		Short: "Start yaml parser",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&certFile, "certificate", "public.cert", "Certificate to establish a secure connection.")
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "secure", "The mode of application to run, secure or http")
}

func main() {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadminpass"
	useSSL := false
	var customTransport *http.Transport

	if mode == "secure" {
		useSSL = true
		caCert, err := os.ReadFile(certFile)
		if err != nil {
			log.Fatalf("Failed to load CA certificate: %v", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}

		customTransport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	} else {
		customTransport = &http.Transport{
			TLSClientConfig:    nil,
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: true,
		}
	}

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure:    useSSL,
		Transport: customTransport,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "testbucket"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	objectName := "testdata"
	filePath := "./test.txt"
	contentType := "application/octet-stream"

	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
