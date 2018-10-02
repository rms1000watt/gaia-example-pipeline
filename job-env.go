package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobEnv(args sdk.Arguments) error {
	log.Println("Evironment Start")
	defer log.Println("Environment Done")

	log.Println("Creating Environment..")
	return nil
}
