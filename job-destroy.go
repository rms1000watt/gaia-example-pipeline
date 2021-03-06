package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobDestroy(args sdk.Arguments) (err error) {
	log.Println("Start: Destroy")
	time.Sleep(1 * time.Second)
	defer log.Println("Done: Destroy")

	if skip(args) {
		log.Println("Skipping: Destroy")
		return
	}

	log.Println("Destroying..")
	return
}
