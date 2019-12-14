package healthcheck

import "time"

// Possible set of values for 'status' field.
const (
	statusPending   = "pending"
	statusChecking  = "checking"
	statusHealthy   = "healthy"
	statusUnhealthy = "unhealthy"
)

// Result is the last healthcheck result.
type Result struct {
	Status    string    `json:"status"`
	Error     error     `json:"error"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
