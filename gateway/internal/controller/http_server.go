package controller

import (
	"net/http"

	"github.com/artem-smola/GitHubRepoInfoTool/gateway/internal/domain"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StarsCount  int    `json:"stargazers_count"`
	ForksCount  int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
}

type RepoInfoProvider interface {
	Execute(owner, repoName string) (*domain.RepoInfo, error)
}

type HTTPServer struct {
	usecase RepoInfoProvider
}

func NewHTTPServer(usecase RepoInfoProvider) *HTTPServer {
	return &HTTPServer{usecase: usecase}
}

// GetRepo godoc
// @Summary      Get GitHub-repository info by owner + repoName
// @Description  Get GitHub-repository info by owner + repoName
// @Tags         repo-info
// @Produce      json
// @Param        owner      path      string  true  "Repository owner (for example: Desbordante)"
// @Param        repoName   path      string  true  "Repository name (for example: desbordante-core)"
// @Success      200    {object}  domain.RepoInfo
// @Failure      404    {object}  map[string]string
// @Failure      429    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /gh-tool/{owner}/{repoName} [get]
func (s *HTTPServer) GetRepoInfoByPath(c *gin.Context) {
	owner := c.Param("owner")
	repoName := c.Param("repoName")

	ans, err := s.usecase.Execute(owner, repoName)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		switch st.Code() {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
		case codes.ResourceExhausted:
			c.JSON(http.StatusTooManyRequests, gin.H{"error": st.Message()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		}
		return
	}

	c.JSON(http.StatusOK, RepoInfo{
		Name:        ans.Name,
		Description: ans.Description,
		StarsCount:  ans.StarsCount,
		ForksCount:  ans.ForksCount,
		CreatedAt:   ans.CreatedAt,
	})
}

// GetRepoByURL godoc
// @Summary      Get GitHub-repository info by URL
// @Description  Get GitHub-repository info by URL
// @Tags         repo-info
// @Produce      json
// @Param        url  query     string  true  "GitHub repository URL"
// @Success      200  {object}  domain.RepoInfo
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      429  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /gh-tool/by-url [get]
func (s *HTTPServer) GetRepoInfoByURL(c *gin.Context) {
	rawURL := c.Query("url")
	owner, repoName, err := ParseGitHubRepoURL(rawURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ans, err := s.usecase.Execute(owner, repoName)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
				return
			case codes.ResourceExhausted:
				c.JSON(http.StatusTooManyRequests, gin.H{"error": st.Message()})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RepoInfo{
		Name:        ans.Name,
		Description: ans.Description,
		StarsCount:  ans.StarsCount,
		ForksCount:  ans.ForksCount,
		CreatedAt:   ans.CreatedAt,
	})
}

func (s *HTTPServer) RegisterRoutes(router *gin.Engine) {
	router.GET("/gh-tool/by-url", s.GetRepoInfoByURL)
	router.GET("/gh-tool/:owner/:repoName", s.GetRepoInfoByPath)
}
