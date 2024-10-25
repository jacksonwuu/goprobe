package prober

type ProbeType string

const (
	HTTP   ProbeType = "HTTP"
	TCP    ProbeType = "TCP"
	UDP    ProbeType = "UDP"
	TLS    ProbeType = "TLS"
	GRPC   ProbeType = "GRPC"
	THRIFT ProbeType = "THRIFT"
)
