package employee

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type updateQuantityProduct struct {
	repo editEmployeeQuantityByProductNameRepo
}

func NewUpdateQuantityProduct(repo editEmployeeQuantityByProductNameRepo) *updateQuantityProduct {
	return &updateQuantityProduct{
		repo: repo,
	}
}

func (s *updateQuantityProduct) Handler(c *handler.Ctx) {
	var req UpdateQuantityProductRequest
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

	if err := s.repo.UpdateEmployeeQuantityByProduct(c.Request.Context(), req.Product, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
