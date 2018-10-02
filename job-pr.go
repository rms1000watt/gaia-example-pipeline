package main

import (
	"log"

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
	return
}
