package user

type User interface {
	GeneratePassword() (password string)
	GetUsername() (username string)
}
