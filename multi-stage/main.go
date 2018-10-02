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
		},
		{
			Handler:     jobEnv,
			Title:       "Env",
			Description: "This sets up an environment",
			DependsOn:   []string{"Tag"},
		},
		{
			Handler:     jobBuild,
			Title:       "Build",
			Description: "This builds everything",
			DependsOn:   []string{"Env"},
		},
		{
			Handler:     jobDeploy,
			Title:       "Deploy",
			Description: "This will deploy everything",
			DependsOn:   []string{"Build"},
		},
		{
			Handler:     jobDestroy,
			Title:       "Destroy",
			Description: "This will destroy the environment",
			DependsOn:   []string{"Deploy"},
		},
		{
			Handler:     jobPR,
			Title:       "PR",
			Description: "This will make the PR",
			DependsOn:   []string{"Destroy"},
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
