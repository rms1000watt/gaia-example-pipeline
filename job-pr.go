package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobPR(args sdk.Arguments) (err error) {
	defer log.Println("Done: PR")

	log.Println("Creating PR..")
	return
}
