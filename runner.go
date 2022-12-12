package runner

import "context"

// Runner is a lifecycle manager that run something.
type Runner interface {
	// Name returns the name of the runner.
	Name() string
	// Run starts running this runner.
	// It returns an error if any error occurred during running.
	// Note the running may block current goroutine,
	// and you have better to run this in a new goroutine.
	Run(ctx context.Context) error
	// Stop stops running this runner.
	Stop(ctx context.Context) error
}
