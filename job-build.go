package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobBuild(args sdk.Arguments) (err error) {
	defer log.Println("Done: Build")

	log.Println("Building..")
	return
}
