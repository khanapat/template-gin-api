package account

import (
	"fmt"
	"template-gin-api/utils"
	"unicode/utf8"

	"github.com/pkg/errors"
)

type InquiryAccountRequest struct {
	Id     *string `json:"id" example:"1233" form:"id"`
	Email  *string `json:"email" example:"email@gmail.com" form:"email"`
	RoleId *int    `json:"roleId" example:"1" form:"role_id"`
}

type CreateAccountRequest struct {
	FirstName string  `json:"firstName" example:"first name"`
	LastName  string  `json:"lastName" example:"last name"`
	Email     string  `json:"email" example:"email@gmail.com"`
	Balance   float64 `json:"balance" example:"1.0"`
	RoleId    int     `json:"roleId" example:"1"`
}

func (req *CreateAccountRequest) validate() error {
	if utf8.RuneCountInString(req.FirstName) == 0 {
		return errors.Wrapf(errors.New("'firstName' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.LastName) == 0 {
		return errors.Wrapf(errors.New("'lastName' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Email) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'email' must be REQUIRED field but the input is '%v'.", req.Email)), utils.ValidateFieldError)
	}
	if req.Balance < 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'balance' must be REQUIRED field but the input is '%v'.", req.Balance)), utils.ValidateFieldError)
	}
	if req.RoleId < 1 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'roleId' must be REQUIRED field but the input is '%v'.", req.RoleId)), utils.ValidateFieldError)
	}
	return nil
}

type UpdateAccountRequest struct {
	Id      string  `json:"id" example:"1233"`
	Balance float64 `json:"balance" example:"1.0"`
	RoleId  int     `json:"roleId" example:"1"`
}

func (req *UpdateAccountRequest) validate() error {
	if utf8.RuneCountInString(req.Id) == 0 {
		return errors.Wrapf(errors.New("'id' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if req.Balance < 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'balance' must be REQUIRED field but the input is '%v'.", req.Balance)), utils.ValidateFieldError)
	}
	if req.RoleId < 1 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'roleId' must be REQUIRED field but the input is '%v'.", req.RoleId)), utils.ValidateFieldError)
	}
	return nil
}

type UpsertAccountRequest struct {
	Id        string  `json:"id" example:"1233"`
	FirstName string  `json:"firstName" example:"first name"`
	LastName  string  `json:"lastName" example:"last name"`
	Email     string  `json:"email" example:"email@gmail.com"`
	Balance   float64 `json:"balance" example:"1.0"`
	RoleId    int     `json:"roleId" example:"1"`
}

func (req *UpsertAccountRequest) validate() error {
	if utf8.RuneCountInString(req.Id) == 0 {
		return errors.Wrapf(errors.New("'id' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.FirstName) == 0 {
		return errors.Wrapf(errors.New("'firstName' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.LastName) == 0 {
		return errors.Wrapf(errors.New("'lastName' must be REQUIRED field"), utils.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Email) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'email' must be REQUIRED field but the input is '%v'.", req.Email)), utils.ValidateFieldError)
	}
	if req.Balance < 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'balance' must be REQUIRED field but the input is '%v'.", req.Balance)), utils.ValidateFieldError)
	}
	if req.RoleId < 1 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'roleId' must be REQUIRED field but the input is '%v'.", req.RoleId)), utils.ValidateFieldError)
	}
	return nil
}
