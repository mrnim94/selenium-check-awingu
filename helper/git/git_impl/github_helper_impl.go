package git_impl

import (
	"context"
	b64 "encoding/base64"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"selenium-check-awingu/log"
)

func (cg *ConfigurationGithub) GetContentYaml(owner string, repo string, path string) (contentYaml []byte, err error) {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cg.GithubAccessToken},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	contentNim, _, _, err := client.Repositories.GetContents(context, owner, repo, path, nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err

	}
	sDec, _ := b64.StdEncoding.DecodeString(*contentNim.Content)
	if err != nil {
		log.Error(err.Error())
		return nil, err

	}

	return sDec, nil
}
