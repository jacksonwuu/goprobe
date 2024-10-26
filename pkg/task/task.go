package task

import (
	"strconv"
	"sync"
	"time"

	"github.com/jacksonwuu/goprobe/pkg/prober"
	"github.com/sirupsen/logrus"
)

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

// HTTPTask will use HTTPProber to do it's work.
type HTTPTask struct {
	ipOrDomain string
	port       int
	timeout    time.Duration
	interval   time.Duration

	stopCh  chan struct{}
	running bool
	mu      sync.Mutex

	onSuccess        func(id string) error
	onFail           func(id string) error
	successThreshold int
	failThreshold    int
}

func NewHTTPTask(ipOrDomain string, port int, onSuccess func(id string) error, onFail func(id string) error) *HTTPTask {
	return &HTTPTask{
		ipOrDomain: ipOrDomain,
		port:       port,
		timeout:    2 * time.Second,  // TODO: make it configurable
		interval:   10 * time.Second, // TODO: make it configurable

		stopCh:  make(chan struct{}, 1), // Buffer so the channel can be non-blocking
		running: false,
		mu:      sync.Mutex{},

		onSuccess:        onSuccess,
		onFail:           onFail,
		successThreshold: 1,
		failThreshold:    3,
	}
}

// TODO
func (task *HTTPTask) DeepCopy() *HTTPTask {
	newTask := &HTTPTask{}
	return newTask
}

// Start a infinity loop to do the probe (use prober) util it get a signal from stopCh.
func (task *HTTPTask) run() {

	ticker := time.NewTicker(task.interval)
	defer ticker.Stop()

	for {
		select {
		case <-task.stopCh:
			return
		case <-ticker.C:
			// TODO: do the probe and callback onSucces or onFail
			logrus.Info("Do porbe.")
		}
	}
}

// ID return the unique identifier of a HTTPTask, TaskManager will use this to identify duplicate task.
func (task *HTTPTask) ID() string {
	return task.ipOrDomain + ":" + strconv.Itoa(task.port)
}
