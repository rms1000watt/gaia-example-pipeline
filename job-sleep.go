package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func jobSleep(args sdk.Arguments) (err error) {
	log.Println("Start: Sleep")
	defer log.Println("Done: Sleep")

	log.Println("Sleeping..")
	time.Sleep(2 * time.Second)
	return
}
