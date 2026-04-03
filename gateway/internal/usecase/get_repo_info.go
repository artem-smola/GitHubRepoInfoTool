package usecase

import "github.com/artem-smola/GitHubRepoInfoTool/gateway/internal/domain"

type GRPCClient interface {
	GetRepoInfo(owner, repoName string) (*domain.RepoInfo, error)
}

type RepoInfoProvider struct {
	client GRPCClient
}

func NewRepoInfoProvider(client GRPCClient) *RepoInfoProvider {
	return &RepoInfoProvider{client: client}
}

func (p *RepoInfoProvider) Execute(owner, repoName string) (*domain.RepoInfo, error) {
	return p.client.GetRepoInfo(owner, repoName)
}
