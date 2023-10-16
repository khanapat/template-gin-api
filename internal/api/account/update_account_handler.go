package account

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type updateAccount struct {
	repo editAccountRepo
}

func NewUpdateAccount(repo editAccountRepo) *updateAccount {
	return &updateAccount{
		repo: repo,
	}
}

func (s *updateAccount) Handler(c *handler.Ctx) {
	var req UpdateAccountRequest
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

	if err := s.repo.UpdateAccount(c.Request.Context(), req.Id, req.Balance, req.RoleId); err != nil {
		if err, ok := err.(*pgconn.PgError); ok {
			switch err.Code {
			case "23503":
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "account not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}
