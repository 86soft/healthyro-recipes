package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

var appErrors []error

func main() {
	log := zerolog.New(os.Stdout)
	Start(log)
	if len(appErrors) != 0 {
		for _, err := range appErrors {
			log.Error().Msg(err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func Start(log zerolog.Logger) {
	svc, err := setup(log)
	if err != nil {
		appErrors = append(appErrors, fmt.Errorf("setup: %w", err))
		return
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	result := svc.Run()

	select {
	case err = <-result:
		if err != nil {
			appErrors = append(appErrors, fmt.Errorf("run: %w", err))
		}
	case <-exit:
	}

	problems := svc.Stop()
	if problems != nil {
		appErrors = append(appErrors, fmt.Errorf("problems: %w", problems))
	}
	return
}
