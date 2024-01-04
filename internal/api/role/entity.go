package role

type Role struct {
	Id          int    `json:"id" example:"1"`
	Title       string `json:"title" example:"admin"`
	Description string `json:"description" example:"system owner"`
}
