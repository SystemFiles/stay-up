package util

import (
	"strings"

	"github.com/systemfiles/stay-up/api/types"
)

func GetProtocolFromString(protocol string) types.ServiceProtocol {
	switch strings.ToUpper(protocol) {
	case "UDP":
		return types.PROTO_UDP
	case "TCP":
		return types.PROTO_TCP
	default:
		return types.PROTO_TCP
	}
}