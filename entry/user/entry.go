package user

import (
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/entry/grpcEntry"

	"github.com/lavaxx/services/entry/user/handler"
)

var name = "user"

func GetEntry() entry.Entry {
	ent := grpcEntry.New(name)
	ent.Description("user service")
	ent.Register(&handler.User{})
	return ent
}
