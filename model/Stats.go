package model

type Metrics struct {
	Timestamp string `json:"@timestamp"`
	Container string `json:"container"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
}
