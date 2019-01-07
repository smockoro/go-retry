package retry

import "time"

type config struct {
	attempts      int
	retryAlgoritm func() time.Duration
	circuitBreak  func() bool
}

type Option func(c *config) error

func InitConfig() *config {
	return &config{
		attempts: 10,
		retryAlgoritm: func() time.Duration {
			return 10 * time.Second
		},
		circuitBreak: func() bool {
			return false
		},
	}
}

// SetAttempts :
func SetAttempts(c *config) {
}
