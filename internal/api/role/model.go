package role

import (
	"template-gin-api/utils"
	"unicode/utf8"

	"github.com/pkg/errors"
)

type CreateRoleRequest struct {
	Title       string `json:"title" example:"admin"`
	Description string `json:"description" example:"admin role"`
}

func (req *CreateRoleRequest) validate() error {
	if utf8.RuneCountInString(req.Title) == 0 {
		return errors.Wrapf(errors.New("'title' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Description) == 0 {
		return errors.Wrapf(errors.New("'description' must be REQUIRED field"), utils.ValidateFieldError)
	}
	return nil
}
