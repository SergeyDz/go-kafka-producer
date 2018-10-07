package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	KafkaBrokers string
	Topic        string
	Timeout      time.Duration
}

// read configuration
func (v *Config) Init() {
	v.KafkaBrokers = os.Getenv("KAFKA_BROKER")
	v.Topic = os.Getenv("TOPIC")

	timeoutStr := os.Getenv("TIMEOUT")
	timeout, err := time.ParseDuration(timeoutStr)

	if err != nil {
		fmt.Println(err)
		timeout = 1 * time.Minute
	}

	v.Timeout = timeout
}

func NewConfig() Config {
	v := Config{}
	v.Init()
	return v
}
