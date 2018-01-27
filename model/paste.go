package model

import "time"

// Paste represents a paste containing some text.
type Paste struct {
	Slug        string    `json:"slug" dynamodbav:"Slug"`
	Title       string    `json:"title,omitempty" dynamodbav:"Title"`
	TimeCreated time.Time `json:"time_created,omitempty" dynamodbav:"TimeCreated"`
	Content     string    `json:"content,omitempty" dynamodbav:"Content"`
	Language    string    `json:"language,omitempty" dynamodbav:"Language"`
}
