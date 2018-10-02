package main

import (
	sdk "github.com/gaia-pipeline/gosdk"
)

func main() {
	jobs := sdk.Jobs{
		{
			Handler:     jobTag,
			Title:       "Tag",
			Description: "This tags all the github repos",
			Args: sdk.Arguments{
				{
					Description: "Skip step 1: Tag",
					Type:        sdk.BoolInp,
					Key:         "skip",
				},
			},
		},
		{
			Handler:     jobEnv,
			Title:       "Env",
			Description: "This sets up an environment",
			DependsOn:   []string{"Tag"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 2: Env",
					Type:        sdk.BoolInp,
					Key:         "skip",
				},
			},
		},
		{
			Handler:     jobBuild,
			Title:       "Build",
			Description: "This builds everything",
			DependsOn:   []string{"Env"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 3: Build",
					Type:        sdk.BoolInp,
					Key:         "skip",
				},
			},
		},
		{
			Handler:     jobDeploy,
			Title:       "Deploy",
			Description: "This will deploy everything",
			DependsOn:   []string{"Build"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 4: Deploy",
					Type:        sdk.BoolInp,
					Key:         "skip",
				},
			},
		},
		{
			Handler:     jobDestroy,
			Title:       "Destroy",
			Description: "This will destroy the environment",
			DependsOn:   []string{"Deploy"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 5: Destroy",
					Type:        sdk.BoolInp,
					Key:         "skip",
				},
			},
		},
		{
			Handler:     jobPR,
			Title:       "PR",
			Description: "This will make the PR",
			DependsOn:   []string{"Destroy"},
			Args: sdk.Arguments{
				{
					Description: "Skip step 6: PR",
					Type:        sdk.BoolInp,
					Key:         "skip",
				},
			},
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
