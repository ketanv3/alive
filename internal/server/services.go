package server

import (
	"sync"

	"codezest.in/alive/internal/healthcheck"
	"codezest.in/alive/internal/registry"
)

type hcResultPair struct {
	name   string
	result healthcheck.Result
}

func runHealthCheck() map[string]healthcheck.Result {
	var ch = make(chan hcResultPair, 0)
	var wg sync.WaitGroup
	var rg sync.WaitGroup

	// Perform all the healthchecks concurrently.
	for _, hc := range registry.Get() {
		wg.Add(1)
		go func(hc *healthcheck.HealthCheck) {
			defer wg.Done()
			ch <- hcResultPair{name: hc.Definition.Name, result: hc.Result()}
		}(hc)
	}

	// Collect the results until the channel closes.
	results := make(map[string]healthcheck.Result)
	rg.Add(1)
	go func() {
		defer rg.Done()
		for elm := range ch {
			results[elm.name] = elm.result
		}
	}()

	// Wait for all healthchecks to finish and close the results channel.
	wg.Wait()
	close(ch)
	rg.Wait()

	return results
}
