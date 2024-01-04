package role

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type registerRole struct {
	repo registerRoleRepo
}

func NewRegisterRole(repo registerRoleRepo) *registerRole {
	return &registerRole{
		repo: repo,
	}
}

func (s *registerRole) Handler(c *handler.Ctx) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := req.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := s.repo.CreateRole(c.Context, req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, id)
}
