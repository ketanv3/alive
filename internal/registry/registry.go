package registry

import "codezest.in/alive/internal/healthcheck"

var registry = make(map[string]*healthcheck.HealthCheck)

// Initialize takes a list of healthcheck definition files and creates instances for them.
func Initialize(configs []string) error {
	for _, config := range configs {
		hc, err := healthcheck.New(config)
		if err != nil {
			return err
		}

		if hc.Definition.Strategy == "async" {
			hc.StartBackgroundCheck()
		}

		registry[hc.Definition.Name] = hc
	}

	return nil
}

// Get simply returns the registry.
func Get() map[string]*healthcheck.HealthCheck {
	return registry
}

// GetList gives back the list of registered healthchecks.
func GetList() []healthcheck.Definition {
	var healthchecks = make([]healthcheck.Definition, 0)

	for _, hc := range registry {
		healthchecks = append(healthchecks, hc.Definition)
	}

	return healthchecks
}
