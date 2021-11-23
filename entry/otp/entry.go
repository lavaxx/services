package otp

import (
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/entry/grpcEntry"

	"github.com/lavaxx/services/entry/otp/handler"
)

var Name = "otp"

func GetEntry() entry.Entry {
	ent := grpcEntry.New(Name)
	ent.Description("One Time Password")
	ent.Register(new(handler.Otp))
	return ent
}
