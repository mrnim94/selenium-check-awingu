package git_impl

import (
	"selenium-check-awingu/helper/git"
)

type ConfigurationGithub struct {
	GithubAccessToken string
}

func NewGitHubConnection(cg *ConfigurationGithub) git.Github {
	return &ConfigurationGithub{
		GithubAccessToken: cg.GithubAccessToken,
	}
}
