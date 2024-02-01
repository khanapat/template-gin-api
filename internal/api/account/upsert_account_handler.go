package account

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type upsertAccount struct {
	repo upsertAccountRepo
}

func NewUpsertAccount(repo upsertAccountRepo) *upsertAccount {
	return &upsertAccount{
		repo: repo,
	}
}

func (s *upsertAccount) Handler(c *handler.Ctx) {
	var req UpsertAccountRequest
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

	if err := s.repo.UpsertAccount(c.Request.Context(), req.Id, req.FirstName, req.LastName, req.Email, req.Balance, req.RoleId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
