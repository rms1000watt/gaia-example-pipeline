package main

import (
	"log"
	"strings"

	sdk "github.com/gaia-pipeline/gosdk"
)

const (
	keySkip = "skip"
)

func main() {
	jobs := sdk.Jobs{
		sdk.Job{
			Handler:     jobTag,
			Title:       "Tag",
			Description: "This tags all the github repos",
			Args: sdk.Arguments{
				{
					Description: "Skip step 1: Tag",
					Type:        sdk.TextFieldInp,
					Key:         keySkip,
					Value:       "false",
				},
			},
		},
		sdk.Job{
			Handler:     jobEnv,
			Title:       "Env",
			Description: "This sets up an environment",
			DependsOn:   []string{"Tag"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 2: Env",
					Type:        sdk.TextFieldInp,
					Key:         keySkip,
					Value:       "false",
				},
			},
		},
		sdk.Job{
			Handler:     jobBuild,
			Title:       "Build",
			Description: "This builds everything",
			DependsOn:   []string{"Env"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 3: Build",
					Type:        sdk.TextFieldInp,
					Key:         keySkip,
					Value:       "false",
				},
			},
		},
		sdk.Job{
			Handler:     jobDeploy,
			Title:       "Deploy",
			Description: "This will deploy everything",
			DependsOn:   []string{"Build"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 4: Deploy",
					Type:        sdk.TextFieldInp,
					Key:         keySkip,
					Value:       "false",
				},
			},
		},
		sdk.Job{
			Handler:     jobDestroy,
			Title:       "Destroy",
			Description: "This will destroy the environment",
			DependsOn:   []string{"Deploy"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 5: Destroy",
					Type:        sdk.TextFieldInp,
					Key:         keySkip,
					Value:       "false",
				},
			},
		},
		sdk.Job{
			Handler:     jobPR,
			Title:       "PR",
			Description: "This will make the PR",
			DependsOn:   []string{"Destroy"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 6: PR",
					Type:        sdk.TextFieldInp,
					Key:         keySkip,
					Value:       "false",
				},
			},
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}

func skip(args sdk.Arguments) (skip bool) {
	for _, arg := range args {
		log.Println("Key:", arg.Key)
		log.Println("Value:", arg.Value)
		if strings.Contains(arg.Key, keySkip) && arg.Value != "false" {
			skip = true
			return
		}
	}

	return
}
