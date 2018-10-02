package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobDestroy(args sdk.Arguments) (err error) {
	defer log.Println("Done: Destroy")

	log.Println("Destroying..")
	return
}
