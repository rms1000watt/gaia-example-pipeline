package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobBuild(args sdk.Arguments) error {
	log.Println("Build Start")
	defer log.Println("Build Done")

	log.Println("Building..")
	return nil
}
