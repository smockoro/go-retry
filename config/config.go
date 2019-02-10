package retry

import "time"

type Config struct {
	Attempts      int
	RetryAlgoritm func() time.Duration
}

type Option func(c *Config) error

func InitConfig() *Config {
	return &Config{
		Attempts: 10,
		RetryAlgoritm: func() time.Duration {
			return 10 * time.Second
		},
	}
}

// SetAttempts :
func SetAttempts(c *Config) {
}
