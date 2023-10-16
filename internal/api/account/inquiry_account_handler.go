package account

import (
	"net/http"
	"strconv"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type inquiryAccount struct {
	repo getAccountRepo
}

func NewInquiryAccount(repo getAccountRepo) *inquiryAccount {
	return &inquiryAccount{repo}
}

func (s *inquiryAccount) Handler(c *handler.Ctx) {
	id := c.Query("id")
	email := c.Query("email")
	roleId := c.Query("role_id")
	if roleId != "" {
		if _, err := strconv.Atoi(roleId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	// or
	// var req InquiryAccountRequest
	// if err := c.ShouldBind(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	m := make(map[string]interface{})
	if id != "" {
		m["id"] = id
	}
	if email != "" {
		m["email"] = email
	}
	if roleId != "" {
		m["role_id"] = roleId
	}

	accounts, err := s.repo.InquiryAccount(c.Request.Context(), m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
