package users

type UserCreateForm struct {
	Email    string
	Name     string
	Password string
}

type User struct {
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
