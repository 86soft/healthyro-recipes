package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := Start()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func Start() error {
	svc, err := setup()
	if err != nil {
		return err
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	r := svc.Run()

	select {
	case err := <-r:
		if err != nil {
			err = fmt.Errorf("run: %w", err)
		}
	case <-exit:
	}

	problems := svc.Stop()
	if problems != nil {
		err = fmt.Errorf("problems: %w", err)
	}

	return err
}
