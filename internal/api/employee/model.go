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

// swagger:response inquiryEmployeeResponse
type InquiryEmployeeResponse struct {
	// in:body
	Username  string   `json:"username" example:"username"`
	Email     string   `json:"email" example:"email@email.com"`
	Metadata  Metadata `json:"metadata"`
	Job       []string `json:"job" example:"[]"`
	CreatedAt int64    `json:"createdAt" example:"1707649444"`
	UpdatedAt int64    `json:"updatedAt" example:"1707649444"`
}

type Metadata struct {
	Customer string `json:"customer" example:"John Doe"`
	Items    Item   `json:"items"`
}

type Item struct {
	Product  string `json:"product" example:"Beer"`
	Quantity int    `json:"qty" example:"5"`
}

type UpdateQuantityProductRequest struct {
	Product  string `json:"product" example:"Beer"`
	Quantity int    `json:"qty" example:"5"`
}

func (req *UpdateQuantityProductRequest) validate() error {
	if utf8.RuneCountInString(req.Product) == 0 {
		return errors.New("'product' must be REQUIRED field")
	}
	if req.Quantity == 0 {
		return errors.New("'qty' must be REQUIRED field")
	}
	return nil
}
