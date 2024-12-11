package memcache

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	configLib "github.com/rizwijaya/miniWallet/infrastructures/config"
)

var Memcache *memcache.Client

func NewMemcache(config configLib.LoadConfig) *memcache.Client {
	Memcache = memcache.New(fmt.Sprintf("%s:%s", config.Memcache.Host, config.Memcache.Port))

	return Memcache
}
