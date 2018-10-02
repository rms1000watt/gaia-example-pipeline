package main

import (
	sdk "github.com/gaia-pipeline/gosdk"
)

func main() {
	skipArgs := sdk.Arguments{
		{
			Description: "Skip this step",
			Type:        sdk.BoolInp,
			Key:         "skip",
		},
	}

	jobs := sdk.Jobs{
		{
			Handler:     jobTag,
			Title:       "Tag",
			Description: "This tags all the github repos",
			Args:        skipArgs,
		},
		{
			Handler:     jobEnv,
			Title:       "Env",
			Description: "This sets up an environment",
			DependsOn:   []string{"Tag"},
			Args:        skipArgs,
		},
		{
			Handler:     jobBuild,
			Title:       "Build",
			Description: "This builds everything",
			DependsOn:   []string{"Env"},
			Args:        skipArgs,
		},
		{
			Handler:     jobDeploy,
			Title:       "Deploy",
			Description: "This will deploy everything",
			DependsOn:   []string{"Build"},
			Args:        skipArgs,
		},
		{
			Handler:     jobDestroy,
			Title:       "Destroy",
			Description: "This will destroy the environment",
			DependsOn:   []string{"Deploy"},
			Args:        skipArgs,
		},
		{
			Handler:     jobPR,
			Title:       "PR",
			Description: "This will make the PR",
			DependsOn:   []string{"Destroy"},
			Args:        skipArgs,
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
