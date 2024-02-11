package employee

import "time"

type Employee struct {
	Id              int       `json:"id" db:"id"`
	Username        string    `json:"username" db:"username"`
	Email           string    `json:"email" db:"email"`
	Metadata        *string   `json:"metadata" db:"metadata"`
	Job             *[]string `json:"job" db:"job"`
	CreatedDateTime time.Time `json:"createdDateTime" db:"created_date_time"`
	UpdatedDateTime time.Time `json:"updatedDateTime" db:"updated_date_time"`
}
