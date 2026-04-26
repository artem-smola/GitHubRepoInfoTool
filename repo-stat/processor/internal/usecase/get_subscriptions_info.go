package usecase

import "github.com/artem-smola/GitHubRepoInfoTool/processor/internal/domain"

type GetSubscriptionsRepoInfoUsecase struct {
	client Client
}

func NewGetSubscriptionsRepoInfoUsecase(client Client) *GetSubscriptionsRepoInfoUsecase {
	return &GetSubscriptionsRepoInfoUsecase{client: client}
}

func (u *GetSubscriptionsRepoInfoUsecase) Execute() ([]domain.SubscriptionRepoInfo, error) {
	return u.client.GetSubscriptionsRepoInfo()
}
