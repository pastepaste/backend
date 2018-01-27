package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Config contains configuration options for the DynamoDB service.
type Config struct {
	PastesTableName string
}

// DynamoDB represents the DynamoDB service.
type DynamoDB struct {
	db  *dynamodb.DynamoDB
	cfg Config
}

// New creates a new DynamoDB service.
func New(cfg Config) (*DynamoDB, error) {
	awscfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}
	db := dynamodb.New(awscfg)

	return &DynamoDB{
		db:  db,
		cfg: cfg,
	}, nil
}
