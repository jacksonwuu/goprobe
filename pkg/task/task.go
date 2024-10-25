package task

import "github.com/jacksonwuu/goprobe/pkg/prober"

type Task struct {
	// Get task ID
	ID func() string

	// Which kind of probe, http or tcp or someone else
	ProbeType prober.ProbeType

	// Called when it probe success
	OnSuccess func() error

	// Called when it probe fail
	OnFail func() error
}
