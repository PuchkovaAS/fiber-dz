package users

type User struct {
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
