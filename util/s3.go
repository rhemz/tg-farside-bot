package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	DOSpacesS3Client struct {
		Client *s3.Client
	}
)

//func (c DOSpacesS3Client)

func NewDOSpacesS3Client(endpoint string, keyId string, secretkey string) (*DOSpacesS3Client, error) {
	const region = "us-east-1" // digitalocean requires us-east-1 region for all their DCs

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyId, secretkey, "")),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == s3.ServiceID {
				return aws.Endpoint{
					PartitionID:       "aws",
					URL:               endpoint,
					SigningRegion:     region,
					HostnameImmutable: true,
				}, nil
			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})),
	)
	if err != nil {
		return nil, err
	}

	return &DOSpacesS3Client{Client: s3.NewFromConfig(cfg)}, nil
}
