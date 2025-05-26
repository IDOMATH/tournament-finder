package types

type User struct {
	Name         string
	Email        string
	Id           int
	PasswordHash string
}

type LoginFormUser struct {
	Email    string
	Password string
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
