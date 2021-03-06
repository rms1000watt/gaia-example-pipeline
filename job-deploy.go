package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobDeploy(args sdk.Arguments) (err error) {
	log.Println("Start: Deploy")
	time.Sleep(1 * time.Second)
	defer log.Println("Done: Deploy")

	if skip(args) {
		log.Println("Skipping: Deploy")
		return
	}

	log.Println("Deploying..")
	return
}
