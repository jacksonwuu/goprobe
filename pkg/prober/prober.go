package prober

type Prober interface {
	Probe() error
}
