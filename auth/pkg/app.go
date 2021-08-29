package pkg

import (
	log "github.com/vine-io/vine/lib/logger"

	"github.com/vine-io/services/auth/pkg/handler"
)

func Run() {
	app := handler.New()

	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
