package main

import "time"

type config struct {
	GithubToken    string `env:"GH_TOKEN,required"`
	RepoOwner string `env:"REPO_OWNER,required"`
	RepoName  string `env:"REPO_NAME,required"`

	BranchesPath string `env:"BRANCHES_PATH" envDefault:"/branches"`

	Delay time.Duration `env:"DELAY" envDefault:"12h"`
}
