package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	sdk "github.com/gaia-pipeline/gosdk"
)

const (
	keySkip = "skip"
)

func main() {
	jobs := sdk.Jobs{
		sdk.Job{
			Handler:     middleware(jobTag, "Tag"),
			Title:       "Tag",
			Description: "This tags all the github repos",
			Args: sdk.Arguments{
				{
					Description: "Skip step 1: Tag",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		},
		sdk.Job{
			Handler:     middleware(jobEnv, "Env"),
			Title:       "Env",
			Description: "This sets up an environment",
			DependsOn:   []string{"Tag"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 2: Env",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		},
		sdk.Job{
			Handler:     middleware(jobBuild, "Build"),
			Title:       "Build",
			Description: "This builds everything",
			DependsOn:   []string{"Env"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 3: Build",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		},
		sdk.Job{
			Handler:     middleware(jobDeploy, "Deploy"),
			Title:       "Deploy",
			Description: "This will deploy everything",
			DependsOn:   []string{"Build"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 4: Deploy",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		},
		sdk.Job{
			Handler:     middleware(jobDestroy, "Destroy"),
			Title:       "Destroy",
			Description: "This will destroy the environment",
			DependsOn:   []string{"Deploy"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 5: Destroy",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		},
		sdk.Job{
			Handler:     middleware(jobPR, "PR"),
			Title:       "PR",
			Description: "This will make the PR",
			DependsOn:   []string{"Destroy"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 6: PR",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}

func middleware(fn func(sdk.Arguments) error, title string) func(sdk.Arguments) error {
	return func(args sdk.Arguments) (err error) {
		if skip(args) {
			log.Println("Skipping:", title)
			return
		}

		log.Println("Starting:", title)
		return fn(args)
	}
}

func skip(args sdk.Arguments) (skip bool) {
	for _, arg := range args {
		spew.Dump(arg)
		if arg.Key == keySkip && arg.Value == "true" {
			skip = true
			return
		}
	}

	return
}
