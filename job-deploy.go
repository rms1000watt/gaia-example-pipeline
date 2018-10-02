package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobDeploy(args sdk.Arguments) error {
	log.Println("Deploy Start")
	defer log.Println("Deploy Done")

	log.Println("Deploying..")
	return nil
}
