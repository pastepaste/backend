package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/pastepaste/backend/service"
)

// CreatePaste puts a paste into the Pastes table.
func (db *DynamoDB) CreatePaste(pst *service.Paste) error {
	item, err := dynamodbattribute.MarshalMap(pst)
	if err != nil {
		return err
	}

	in := &dynamodb.PutItemInput{
		Item:      item,
		TableName: db.cfg.PastesTableName,
	}
	req := db.db.PutItemRequest(in)
	_, err := req.Send()
	if err != nil {
		return err
	}

	return nil
}

// GetPaste gets a paste with the given slug from the Pastes table.
func (db *DynamoDB) GetPaste(slug string) (*service.Paste, error) {
	in := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"Slug": {
				S: aws.String(slug),
			},
		},
	}
	req := db.db.GetItemRequest(in)
	out, err := req.Send()
	if err != nil {
		if awserr, ok := err.(awserr.Error); ok {
			switch awserr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return nil, NotFoundError{awserr}
			}
		}
		return nil, err
	}

	var pst *service.Paste
	if err := dynamodbattribute.UnmarshalMap(out.Item, &pst); err != nil {
		return err
	}

	return out.Item, nil
}
