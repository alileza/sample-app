package config

import "os"

type Config struct {
	DatabaseDSN        string
	QueueDSN           string
	CustomerAppBaseURL string
}

func Retrieve() *Config {
	return &Config{
		DatabaseDSN:        os.Getenv("DATABASE_DSN"),
		QueueDSN:           os.Getenv("QUEUE_DSN"),
		CustomerAppBaseURL: os.Getenv("CUSTOMER_APP_BASE_URL"),
	}
}
