package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobBuild(args sdk.Arguments) (err error) {
	log.Println("Start: Build")
	defer log.Println("Done: Build")

	if skip(args) {
		log.Println("Skipping: Build")
		return
	}

	log.Println("Building..")
	time.Sleep(1 * time.Second)
	return
}
