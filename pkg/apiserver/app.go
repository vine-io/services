package apiserver

import (
	log "github.com/vine-io/vine/lib/logger"

	"github.com/vine-io/services/pkg/apiserver/server"
)

func Run() {
	s := server.New()

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}