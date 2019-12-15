package healthcheck

// Definition is the basic definition of a healthcheck.
type Definition struct {
	Name        string      `yaml:"name" json:"name"`
	Strategy    string      `yaml:"strategy" json:"strategy"`
	Interval    float64     `yaml:"interval" json:"interval"`
	Checker     Checker     `yaml:"checker" json:"checker"`
	RetryPolicy RetryPolicy `yaml:"retryPolicy" json:"retryPolicy"`
}

// Checker defines which health check to perform and how.
type Checker struct {
	Type       string                 `yaml:"type" json:"type"`
	Parameters map[string]interface{} `yaml:"parameters" json:"parameters"`
}

// RetryPolicy defines how to handle retries in case of failing health checks.
type RetryPolicy struct {
	InitialDelay      float64 `yaml:"initialDelay" json:"initialDelay"`
	BackoffMultiplier float64 `yaml:"backoffMultiplier" json:"backoffMultiplier"`
	MaxRetries        uint32  `yaml:"maxRetries" json:"maxRetries"`
}
