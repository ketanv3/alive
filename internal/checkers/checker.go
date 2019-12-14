package checkers

import "errors"

// HealthChecker is an interface for concrete implementation of health checkers.
// It only contains one function 'Check' that returns if the health check was successful or not.
type HealthChecker interface {
	Check() (bool, error)
}

// Get is a factory function to return the appropriate instance of a health checker.
func Get(factory string, parameters map[string]interface{}) (HealthChecker, error) {
	if factory == "url" {
		return NewURLHealthChecker(parameters)
	}

	return nil, errors.New("invalid health checker type")
}
