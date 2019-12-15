package healthcheck

import (
	"fmt"
	"io/ioutil"
	"time"

	"codezest.in/alive/internal/checkers"
	"gopkg.in/yaml.v2"
)

// HealthCheck wraps the healthcheck definition and it's corresponding checker.
type HealthCheck struct {
	Definition         Definition
	HealthChecker      checkers.HealthChecker
	LastResult         Result
	RunBackgroundCheck bool
}

// New creates a new health check from the given YAML definition file.
func New(definitionFilePath string) (*HealthCheck, error) {
	var definition Definition
	var healthChecker checkers.HealthChecker
	var lastResult = Result{Status: statusPending}

	definitionFileContents, err := ioutil.ReadFile(definitionFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(definitionFileContents, &definition)
	if err != nil {
		return nil, err
	}

	healthChecker, err = checkers.Get(definition.Checker.Type, definition.Checker.Parameters)
	if err != nil {
		return nil, err
	}

	return &HealthCheck{
		Definition:    definition,
		HealthChecker: healthChecker,
		LastResult:    lastResult,
	}, nil
}

// Check wraps the call to the healthcheck's 'Check()' and adds some additional functionality:
// 1. Retry mechanism
// 2. Saving the last result
// 3. In case of async checks, picks up the last result.
func (hc *HealthCheck) Check() Result {
	var healthy bool
	var err error

	result := Result{
		Status:    statusChecking,
		StartTime: time.Now(),
	}

	var retriesLeft = hc.Definition.RetryPolicy.MaxRetries
	var sleepTime = hc.Definition.RetryPolicy.InitialDelay

	for retriesLeft > 0 {
		retriesLeft--

		// Break the loop if the health check succeeded or if we are at the last retry.
		if healthy, err = hc.HealthChecker.Check(); healthy || retriesLeft == 0 {
			break
		}

		// Otherwise backoff for some time and retry the health check.
		time.Sleep(time.Second * time.Duration(sleepTime))
		sleepTime = sleepTime * hc.Definition.RetryPolicy.BackoffMultiplier
	}

	// Update the healthcheck result.
	if healthy {
		result.Status = statusHealthy
	} else {
		result.Status = statusUnhealthy
	}
	result.Error = err
	result.EndTime = time.Now()
	hc.LastResult = result

	return result
}

// StartBackgroundCheck starts the healthcheck in background (async strategy).
func (hc *HealthCheck) StartBackgroundCheck() {
	hc.RunBackgroundCheck = true

	go func() {
		for hc.RunBackgroundCheck {
			r := hc.Check()
			fmt.Println("[background]", hc.Definition.Name, r.Status)
			time.Sleep(time.Second * time.Duration(hc.Definition.Interval))
		}
	}()
}

// StopBackgroundCheck stops the healthcheck in background.
func (hc *HealthCheck) StopBackgroundCheck() {
	hc.RunBackgroundCheck = false
}

// Result returns the result for the given healthcheck based on the healthcheck strategy.
// default: picks up what's defined in the healthcheck definition.
// sync: performs the healthcheck and returns the result.
// async: returns the last healthcheck result.
func (hc *HealthCheck) Result(strategy string) Result {
	if strategy == "default" {
		strategy = hc.Definition.Strategy
	}

	if strategy == "sync" {
		return hc.Check()
	}

	return hc.LastResult
}
