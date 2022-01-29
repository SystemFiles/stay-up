package util

import (
	"encoding/json"
	"strings"

	"github.com/systemfiles/stay-up/api/models"
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

func StructToMap(s interface{}, dest *map[string]interface{}) error {
	svcJson, _ := json.Marshal(s)
	return json.Unmarshal(svcJson, dest)
}

func MapToStruct(m map[string]interface{}) (models.Service, error) {
	var svcDest models.Service
	svcJson, _ := json.Marshal(m)
	err := json.Unmarshal(svcJson, &svcDest)
	if err != nil {
		return models.Service{}, err
	}

	return svcDest, nil
}