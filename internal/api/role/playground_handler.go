package role

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type playground struct {
	repo playgroundRepo
}

func NewPlayground(repo playgroundRepo) *playground {
	return &playground{
		repo: repo,
	}
}

func (s *playground) Handler(c *handler.Ctx) {
	t, err := s.repo.GetCurrentTimestamp(c.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, t)
}
