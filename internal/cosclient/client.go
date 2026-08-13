package cosclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"

	"cosh/internal/config"
)

func NewClient(cfg *config.Config, bucket string) (*cos.Client, error) {
	if bucket == "" {
		bucket = cfg.Bucket
	}
	if bucket == "" {
		return nil, fmt.Errorf("bucket required: use --bucket or set bucket in config")
	}
	bucketURL, err := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("invalid bucket url: %w", err)
	}
	client := cos.NewClient(
		&cos.BaseURL{BucketURL: bucketURL},
		&http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  cfg.SecretID,
				SecretKey: cfg.SecretKey,
			},
		},
	)
	return client, nil
}

func NewServiceClient(cfg *config.Config) (*cos.Client, error) {
	serviceURL, err := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("invalid service url: %w", err)
	}
	client := cos.NewClient(
		&cos.BaseURL{ServiceURL: serviceURL},
		&http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  cfg.SecretID,
				SecretKey: cfg.SecretKey,
			},
		},
	)
	return client, nil
}
