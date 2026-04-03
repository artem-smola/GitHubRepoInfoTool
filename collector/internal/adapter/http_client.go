package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/artem-smola/GitHubRepoInfoTool/collector/internal/domain"
)

type RepoInfo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StarsCount  int       `json:"stargazers_count"`
	ForksCount  int       `json:"forks_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient(client *http.Client) *HTTPClient {
	return &HTTPClient{client: client}
}

func (c *HTTPClient) GetRepoInfo(owner, repoName string) (*domain.RepoInfo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gh-info-tool")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	var repoInfo RepoInfo
	if err := json.NewDecoder(resp.Body).Decode(&repoInfo); err != nil {
		return nil, err
	}

	return &domain.RepoInfo{
		Name:        repoInfo.Name,
		Description: repoInfo.Description,
		StarsCount:  repoInfo.StarsCount,
		ForksCount:  repoInfo.ForksCount,
		CreatedAt:   repoInfo.CreatedAt,
	}, nil
}
