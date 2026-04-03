package adapter

import (
	"context"
	"time"

	"github.com/artem-smola/GitHubRepoInfoTool/gateway/internal/domain"
	"github.com/artem-smola/GitHubRepoInfoTool/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	client gen.CollectorClient
}

func NewGRPCClient(addr string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := gen.NewCollectorClient(conn)
	return &GRPCClient{client: client}, nil
}

func (c *GRPCClient) GetRepoInfo(owner, repoName string) (*domain.RepoInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.client.GetRepoInfo(ctx, &gen.GetRepoInfoRequest{
		Owner:    owner,
		RepoName: repoName,
	})
	if err != nil {
		return nil, err
	}
	return &domain.RepoInfo{
		Name:        resp.GetName(),
		Description: resp.GetDescription(),
		StarsCount:  int(resp.GetStarsCount()),
		ForksCount:  int(resp.GetForksCount()),
		CreatedAt:   resp.GetCreatedAt(),
	}, nil
}
