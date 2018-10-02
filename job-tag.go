package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobTag(args sdk.Arguments) (err error) {
	defer log.Println("Done: Tag")

	log.Println("Tagging all images..")
	return
}
