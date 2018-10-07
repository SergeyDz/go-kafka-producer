package model

type Metrics struct {
	Timestamp string `json:"timespamp"`
	Container string `json:"container"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
}
