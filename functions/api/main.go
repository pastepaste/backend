package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

func main() {
	lambda.Start(func() error {
		return errors.New("unimplemented")
	})
}
