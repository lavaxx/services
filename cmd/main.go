package main

import (
	"github.com/pubgo/lava"

	"github.com/lavaxx/services/entry/db"
	"github.com/lavaxx/services/entry/otp"
	"github.com/lavaxx/services/entry/user"
)

func main() {
	lava.Run(
		"services",
		otp.GetEntry(),
		db.GetEntry(),
		user.GetEntry(),
	)
}
