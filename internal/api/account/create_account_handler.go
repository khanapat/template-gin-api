package account

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

type registerAccount struct {
	repo registerAccountRepo
}

func NewRegisterAccount(repo registerAccountRepo) *registerAccount {
	return &registerAccount{repo}
}

func (s *registerAccount) Handler(c *handler.Ctx) {
	var req CreateAccountRequest
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

	if err := s.repo.CreateAccount(c.Request.Context(), uuid.NewString(), req.FirstName, req.LastName, req.Email, req.Balance, req.RoleId); err != nil {
		if err, ok := err.(*pgconn.PgError); ok {
			switch err.Code {
			case "23503":
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errors.Wrap(err, "role_id not found").Error(),
				})
				return
			case "23505":
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errors.Wrap(err, "email already exists").Error(),
				})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
