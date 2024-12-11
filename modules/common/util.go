package common

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/rizwijaya/miniWallet/infrastructures/logger"
)

func MustMarshal(payload interface{}) string {
	marshalledPayload, err := jsoniter.MarshalToString(payload)
	if err != nil {
		log.Info("Failed to marshal payload: %s, error: %s", payload, err.Error())
	}
	return marshalledPayload
}
