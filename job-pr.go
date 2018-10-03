package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobPR(args sdk.Arguments) (err error) {
	log.Println("Start: PR")
	time.Sleep(1 * time.Second)
	defer log.Println("Done: PR")

	if skip(args) {
		log.Println("Skipping: PR")
		return
	}

	log.Println("Creating PR..")
	return
}
