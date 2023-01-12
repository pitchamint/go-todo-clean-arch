package models

// Todo ...
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type LoginUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TableName use to specific table
func (b *Todo) TableName() string {
	return "todo"
}

func (b *LoginUser) TableName() string {
	return "login_users"
}
