package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobPR(args sdk.Arguments) (err error) {
	log.Println("Start: PR")
	defer log.Println("Done: PR")

	if skip(args) {
		log.Println("Skipping: PR")
		return
	}

	log.Println("Creating PR..")
	time.Sleep(1 * time.Second)
	return
}
