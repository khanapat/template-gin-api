package employee

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type upsertEmployee struct {
	repo upsertEmployeeRepo
}

func NewUpsertEmployee(repo upsertEmployeeRepo) *upsertEmployee {
	return &upsertEmployee{
		repo: repo,
	}
}

func (s *upsertEmployee) Handler(c *handler.Ctx) {
	var req UpsertEmployeeRequest
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

	if err := s.repo.UpsertEmployee(c.Request.Context(), req.Id, req.Username, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
