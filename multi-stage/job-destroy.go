package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobDestroy(args sdk.Arguments) error {
	log.Println("Destroying..")
	return nil
}
