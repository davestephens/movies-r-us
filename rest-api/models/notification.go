package models

type Notification struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}
