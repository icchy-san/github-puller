package usecase

import (
	"context"
	"os"

	"github.com/google/go-github/v28/github"
)

// GH ... struct
type GH struct {
	Client *github.Client
}

// GetPrivateRepositoryNames ... returns repository name list
func (gh *GH) GetPrivateRepositoryNames(ctx context.Context) []string {
	var repoNames []string

	opt := &github.RepositoryListByOrgOptions{Type: "private"}
	repos, _, _ := gh.Client.Repositories.ListByOrg(ctx, os.Getenv("ORGANIZATION_ACCOUNT"), opt)
	for _, repo := range repos {
		repoNames = append(repoNames, *repo.Name)
	}

	return repoNames
}
