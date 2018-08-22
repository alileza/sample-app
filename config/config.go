package config

type Config struct {
	DatabaseDSN        string
	QueueDSN           string
	CustomerAppBaseURL string
}

func Retrieve() *Config {
	return &Config{
		DatabaseDSN:        "postgres://",
		QueueDSN:           "postgres://",
		CustomerAppBaseURL: "postgres://",
	}
}
