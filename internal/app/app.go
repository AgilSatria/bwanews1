package app

import (
	"github.com/rs/zerolog/log"

	"bwanews/config"
)

func RunServer() {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Msgf("Error conection to database : %v", err)
		return
	}
}
