package user

import (
	"fmt"

	"github.com/segmentio/fasthash/fnv1a"
)

type Teacher struct {
	NIP    int
	Nama   string
	Email  string
	Subjek string
}

func (t Teacher) GeneratePassword() (password string) {
	password = fmt.Sprintf(
		"%v",
		fnv1a.HashString64(t.Email+t.Subjek))
	return
}

func (t Teacher) GetUsername() (username string) {
	username = t.Email
	return
}
