package employee

import (
	"encoding/json"
	"net/http"
	"strconv"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /employees/:id employee inquiryEmployeeById
// responses:
//
// 200: inquiryEmployeeResponse
type inquiryEmployeeById struct {
	repo getEmployeeByIdRepo
}

func NewInquiryEmployeeById(repo getEmployeeByIdRepo) *inquiryEmployeeById {
	return &inquiryEmployeeById{
		repo: repo,
	}
}

func (s *inquiryEmployeeById) Handler(c *handler.Ctx) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee, err := s.repo.InquiryEmployeeById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "employee not found",
		})
		return
	}

	var metadata Metadata
	if employee.Metadata != nil {
		if err := json.Unmarshal([]byte(*employee.Metadata), &metadata); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "invalid json",
			})
			return
		}
	}

	job := make([]string, 0) // var job []string (return nil)
	if employee.Job != nil {
		job = *employee.Job
	}

	resp := InquiryEmployeeResponse{
		Username:  employee.Username,
		Email:     employee.Email,
		Metadata:  metadata,
		Job:       job,
		CreatedAt: employee.CreatedDateTime.Unix(),
		UpdatedAt: employee.CreatedDateTime.Unix(),
	}

	c.JSON(http.StatusOK, resp)
}
