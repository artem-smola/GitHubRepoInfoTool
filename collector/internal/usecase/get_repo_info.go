package usecase

import "github.com/artem-smola/GitHubRepoInfoTool/collector/internal/domain"

type HTTPClient interface {
	GetRepoInfo(owner, repoName string) (*domain.RepoInfo, error)
}

type RepoInfoProvider struct {
	client HTTPClient
}

func NewRepoInfoProvider(client HTTPClient) *RepoInfoProvider {
	return &RepoInfoProvider{client: client}
}

func (p *RepoInfoProvider) Execute(owner, repoName string) (*domain.RepoInfo, error) {
	return p.client.GetRepoInfo(owner, repoName)
}
