package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobTag(args sdk.Arguments) error {
	log.Println("Tag Start")
	defer log.Println("Tag Done")

	log.Println("Tagging all images..")
	return nil
}
