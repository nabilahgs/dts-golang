package user

type Student struct {
	NIM   int
	Nama  string
	Email string
	Kelas string
}

func (s Student) GeneratePassword() (password string) {
	return
}

func (s Student) GetUsername() (username string) {
	username = s.Email
	return
}
