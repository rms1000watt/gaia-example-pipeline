package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

// JobHelloWorld is a hello world job
func JobHelloWorld(args sdk.Arguments) error {
	log.Println("Hello world")
	return nil
}

func main() {
	jobs := sdk.Jobs{
		sdk.Job{
			Handler:     JobHelloWorld,
			Title:       "Hello World",
			Description: "This job says hello world",
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
