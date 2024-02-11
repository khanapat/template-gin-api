package employee

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type inquiryQuantityByProduct struct {
	repo getEmployeeQuantityByProductNameRepo
}

func NewInquiryQuantityByProduct(repo getEmployeeQuantityByProductNameRepo) *inquiryQuantityByProduct {
	return &inquiryQuantityByProduct{
		repo: repo,
	}
}

func (s *inquiryQuantityByProduct) Handler(c *handler.Ctx) {
	product := c.Param("product")

	quantity, err := s.repo.InquiryEmployeeQuantityByProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, quantity)
}
