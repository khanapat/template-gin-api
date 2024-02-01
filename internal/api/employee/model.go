package employee

import (
	"unicode/utf8"

	"github.com/pkg/errors"
)

type UpsertEmployeeRequest struct {
	Id       int    `json:"id" example:"1"`
	Username string `json:"username" example:"username"`
	Email    string `json:"email" example:"email@email.com"`
}

func (req *UpsertEmployeeRequest) validate() error {
	if req.Id == 0 {
		return errors.New("'id' must be REQUIRED field")
	}
	if utf8.RuneCountInString(req.Username) == 0 {
		return errors.New("'username' must be REQUIRED field")
	}
	if utf8.RuneCountInString(req.Email) == 0 {
		return errors.New("'email' must be REQUIRED field")
	}
	return nil
}

type UpsertEmployeeResponse struct {
	Id int `json:"id" example:"1"`
}
