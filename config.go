package main

type config struct {
	GithubToken string `env:"GH_TOKEN,required"`
	RepoOwner   string `env:"REPO_OWNER,required"`
	RepoName    string `env:"REPO_NAME,required"`

	BranchesPath string `env:"BRANCHES_PATH" envDefault:"/branches"`

	Excluded []string `env:"EXCLUDED"`

	Cron string `env:"CRON" envDefault:"0 */12 * * *"`
}
