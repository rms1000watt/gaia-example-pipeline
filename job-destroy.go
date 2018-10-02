package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobDestroy(args sdk.Arguments) error {
	log.Println("Destroy Start")
	defer log.Println("Destroy Done")

	log.Println("Destroying..")
	return nil
}
