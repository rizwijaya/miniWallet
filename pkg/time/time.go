package time

import (
	"time"

	log "github.com/rizwijaya/miniWallet/infrastructures/logger"
)

func Location() *time.Location {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Errorf("Error load location: %v", err)
		loc = time.UTC
	}

	return loc
}

func TimeNow() time.Time {
	return time.Now().In(Location())
}
