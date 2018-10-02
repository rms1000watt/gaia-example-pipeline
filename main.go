package main

import (
	"fmt"
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

const (
	keySkip = "skip"
)

func main() {
	jobs := sdk.Jobs{
		middleware(sdk.Job{
			Handler:     jobTag,
			Title:       "Tag",
			Description: "This tags all the github repos",
			Args: sdk.Arguments{
				{
					Description: "Skip step 1: Tag",
					Type:        sdk.BoolInp,
					Key:         keySkip,
				},
			},
		}),
		middleware(sdk.Job{
			Handler:     jobEnv,
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
		}),
		middleware(sdk.Job{
			Handler:     jobBuild,
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
		}),
		middleware(sdk.Job{
			Handler:     jobDeploy,
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
		}),
		middleware(sdk.Job{
			Handler:     jobDestroy,
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
		}),
		middleware(sdk.Job{
			Handler:     jobPR,
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
		}),
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}

func middleware(in sdk.Job) (out sdk.Job) {
	if skip(in.Args) {
		log.Println("Skipping:", in.Title)
		out = sdk.Job{
			Handler: jobNil,
			Title:   fmt.Sprintf("Skipped: %s", in.Title),
		}
		return
	}

	log.Println("Start:", in.Title)
	out = in
	return
}

func skip(args sdk.Arguments) (skip bool) {
	for _, arg := range args {
		if arg.Key == keySkip && arg.Value == "true" {
			skip = true
			return
		}
	}

	return
}

func jobNil(args sdk.Arguments) (err error) {
	return
}
