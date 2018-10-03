package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobTag(args sdk.Arguments) (err error) {
	log.Println("Start: Tag")
	time.Sleep(1 * time.Second)
	defer log.Println("Done: Tag")

	if skip(args) {
		log.Println("Skipping: Tag")
		return
	}

	log.Println("Tagging all images..")
	return
}
