package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobEnv(args sdk.Arguments) (err error) {
	log.Println("Start: Environment")
	defer log.Println("Done: Environment")

	if skip(args) {
		log.Println("Skipping: Environment")
		return
	}

	log.Println("Creating Environment..")
	return
}
