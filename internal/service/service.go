package service

import (
	"github.com/google/wire"
	"os"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewProxyService, NewNotifyService)

var DebugGroupId int64

func init() {
	id := os.Getenv("DEBUG_GROUP_ID")
	if id == "" {
		panic("can't find \"DEBUG_GROUP_ID\", plz set env")
	}
}
