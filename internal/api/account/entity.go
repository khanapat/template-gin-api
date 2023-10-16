package account

import "time"

type Account struct {
	Id              string     `json:"id" example:"1"`
	FirstName       string     `json:"firstName" example:"first name"`
	LastName        string     `json:"lastName" example:"last name"`
	Email           string     `json:"email" example:"a@gmail.com"`
	Balance         float64    `json:"balance" example:"1"`
	RoleId          int        `json:"roleId" example:"1"`
	RoleTitle       string     `json:"roleTitle" example:"admin"`
	RoleDescription string     `json:"roleDescription" example:"admin"`
	CreatedDateTime time.Time  `json:"createdDateTime" example:"2021-01-01T00:00:00+07:00"`
	UpdatedDateTime *time.Time `json:"updatedDateTime" example:"2021-01-01T00:00:00+07:00"`
}
