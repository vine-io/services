package auth

import (
	log "github.com/vine-io/vine/lib/logger"

	"github.com/vine-io/services/pkg/auth/server"
)

func Run() {
	app := server.New()

	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}