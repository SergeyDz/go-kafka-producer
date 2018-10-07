package model

type Metrics struct {
	Timestamp string `json:"@timestamp"`
	Container string `json:"container"`
	Version   string `json:"version"`
}
