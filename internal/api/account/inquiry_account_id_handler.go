package account

import (
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type inquiryAccountById struct {
	repo getAccountByIdRepo
}

func NewInquiryAccountById(repo getAccountByIdRepo) *inquiryAccountById {
	return &inquiryAccountById{repo}
}

func (s *inquiryAccountById) Handler(c *handler.Ctx) {
	id := c.Param("id")

	account, err := s.repo.InquiryAccountById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if account == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "account not found",
		})
		return
	}

	c.JSON(http.StatusOK, account)
}
