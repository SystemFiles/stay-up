package types

const (
	UP ServiceStatus = "UP"
	DOWN ServiceStatus = "DOWN"
	SLOW ServiceStatus = "SLOW"
)

type ServiceStatus string

func (s ServiceStatus) String() string {
	switch s {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case SLOW:
		return "SLOW"
	}
	return "UKNOWN"
}

const (
	PROTO_TCP ServiceProtocol = "TCP"
	PROTO_UDP ServiceProtocol = "UDP"
)
type ServiceProtocol string

func (s ServiceProtocol) String() string {
	switch s {
	case PROTO_TCP:
		return "tcp"
	case PROTO_UDP:
		return "udp"
	}
	return "tcp"
}