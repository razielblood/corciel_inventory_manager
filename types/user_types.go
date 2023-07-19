package types

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func CreateLoginRequest(username, password string) *LoginRequest {
	return &LoginRequest{Username: username, Password: password}
}

type CreateUserRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Email     string `form:"email" json:"email"`
}

func CreateCreateUserRequest(username, password, firstName, lastName, email string) *CreateUserRequest {
	return &CreateUserRequest{Username: username, Password: password, FirstName: firstName, LastName: lastName, Email: email}
}
