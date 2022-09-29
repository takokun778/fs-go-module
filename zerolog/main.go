package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().Msg("debug")

	log.Info().Msg("info")

	log.Warn().Msg("warn")

	log.Error().Msg("error")
}
