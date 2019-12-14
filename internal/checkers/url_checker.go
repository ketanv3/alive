package checkers

import (
	"errors"
	"net/http"
	"time"
)

// URLHealthChecker checks if the given URL is reachable or not.
type URLHealthChecker struct {
	url     string
	timeout float64
}

// NewURLHealthChecker creates an instance of the URLHealthChecker.
func NewURLHealthChecker(parameters map[string]interface{}) (*URLHealthChecker, error) {
	url, ok1 := parameters["url"]
	timeout, ok2 := parameters["timeout"]

	if !(ok1 && ok2) {
		return nil, errors.New("url and timeout fields are required")
	}

	return &URLHealthChecker{
		url:     url.(string),
		timeout: toFloat64(timeout),
	}, nil
}

// Check performs the actual HTTP GET request to check if the url is reachable.
func (hc URLHealthChecker) Check() (bool, error) {
	client := http.Client{Timeout: time.Second * time.Duration(hc.timeout)}
	resp, err := client.Get(hc.url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	return true, nil
}

func toFloat64(val interface{}) float64 {
	switch i := val.(type) {
	case int:
		return float64(i)
	case float32:
		return float64(i)
	case float64:
		return i
	default:
		panic("unknown type")
	}
}
