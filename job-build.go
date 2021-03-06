package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobBuild(args sdk.Arguments) (err error) {
	log.Println("Start: Build")
	time.Sleep(1 * time.Second)
	defer log.Println("Done: Build")

	if skip(args) {
		log.Println("Skipping: Build")
		return
	}

	log.Println("Building..")
	return
}
