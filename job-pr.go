package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobPR(args sdk.Arguments) error {
	log.Println("PR Start")
	defer log.Println("PR Done")

	log.Println("Creating PR..")
	return nil
}
