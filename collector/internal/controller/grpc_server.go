package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/artem-smola/GitHubRepoInfoTool/collector/internal/domain"
	"github.com/artem-smola/GitHubRepoInfoTool/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoInfoProvider interface {
	Execute(owner, repoName string) (*domain.RepoInfo, error)
}

type GRPCServer struct {
	gen.UnimplementedCollectorServer
	usecase RepoInfoProvider
}

func NewGRPCServer(useCase RepoInfoProvider) *GRPCServer {
	return &GRPCServer{usecase: useCase}
}

func (s *GRPCServer) GetRepoInfo(_ context.Context, req *gen.GetRepoInfoRequest) (*gen.GetRepoInfoResponse, error) {
	repoInfo, err := s.usecase.Execute(req.GetOwner(), req.GetRepoName())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get repo info: %s", err.Error()))
	}

	return &gen.GetRepoInfoResponse{
		Name:        repoInfo.Name,
		Description: repoInfo.Description,
		StarsCount:  int64(repoInfo.StarsCount),
		ForksCount:  int64(repoInfo.ForksCount),
		CreatedAt:   repoInfo.CreatedAt.Format(time.RFC3339),
	}, nil
}
