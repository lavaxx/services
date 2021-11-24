package otp

import (
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/entry/grpcEntry"

	"github.com/lavaxx/services/entry/otp/handler"
	"github.com/lavaxx/services/entry/otp/otp_cfg"
)

func GetEntry() entry.Entry {
	ent := grpcEntry.New(otp_cfg.Name)
	ent.Description("One Time Password")
	ent.Register(&handler.Otp{})
	return ent
}
