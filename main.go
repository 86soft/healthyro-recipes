package main

import (
	"log"
	"os"
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
	if err = svc.Run(); err != nil {
		return err
	}
	if err = svc.Clear(); err != nil {
		return err
	}
	return nil
}
