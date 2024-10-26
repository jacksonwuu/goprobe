package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	CustomRegistry *prometheus.Registry

	httpProbeMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_probe",
			Help: "Indicates success (1) or failure (0) of the probe.",
		},
		[]string{
			"ip", "domain", "port",
		},
	)

	tcpProbeMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tcp_probe",
			Help: "Indicates success (1) or failure (0) of the probe.",
		},
		[]string{
			"ip", "port",
		},
	)
)

func InitMetrics() {
	CustomRegistry = prometheus.NewRegistry()
	CustomRegistry.MustRegister(httpProbeMetrics, tcpProbeMetrics)
}
