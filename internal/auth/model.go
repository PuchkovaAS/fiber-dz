package auth

type userCreateForm struct {
	Email    string
	Name     string
	Password string
}

type userLoginForm struct {
	Email    string
	Password string
}
