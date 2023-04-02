package user

import (
	"fmt"

	"github.com/segmentio/fasthash/fnv1a"
)

type Student struct {
	NIM   int
	Nama  string
	Email string
	Kelas string
}

func (s Student) GeneratePassword() (password string) {
	password = fmt.Sprintf(
		"%v",
		fnv1a.HashString64(s.Email+s.Kelas))
	return
}

func (s Student) GetUsername() (username string) {
	username = s.Email
	return
}
