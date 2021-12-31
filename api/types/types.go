package types

const (
	UP ServiceStatus = iota
	DOWN ServiceStatus = iota
	SLOW ServiceStatus = iota
)

type ServiceStatus int64

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
	PROTO_TCP ServiceProtocol = iota
	PROTO_UDP ServiceProtocol = iota
)
type ServiceProtocol int64

func (s ServiceProtocol) String() string {
	switch s {
	case PROTO_TCP:
		return "tcp"
	case PROTO_UDP:
		return "udp"
	}
	return "tcp"
}