package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobEnv(args sdk.Arguments) (err error) {
	defer log.Println("Done: Environment")

	log.Println("Creating Environment..")
	return
}
