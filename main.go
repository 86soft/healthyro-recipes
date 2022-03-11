package main

import (
	"github.com/rs/zerolog"
	"os"
)

func main() {
	log := zerolog.New(os.Stderr)
	svc, err := Setup(log)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	if err = svc.Run(); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	if err = svc.Clear(); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
