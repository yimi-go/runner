package runner

import "context"

type Status string

func (s Status) String() string {
	return string(s)
}

const (
	StatusUnknown   Status = "UNKNOWN"
	StatusDown             = "DOWN"
	StatusPreparing        = "PREPARING"
	StatusUp               = "UP"
	StatusStopping         = "STOPPING"
)

type Stateful interface {
	// Status reports the runner's status.
	Status() Status
}

// Runner is a lifecycle manager that run something.
type Runner interface {
	// Run starts running this runner.
	// It returns an error if any error occurred during running.
	// Note the running may block current goroutine,
	// and you have better to run this in a new goroutine.
	Run(ctx context.Context) error
	// Stop stops running this runner.
	Stop(ctx context.Context) error
}
