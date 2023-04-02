package main

import (
	"fmt"

	"github.com/nabilahgs/dts-golang/user"
)

func main() {
	var user1 user.User
	user1 = user.Teacher{
		NIP:    1234456,
		Nama:   "Nabilah",
		Email:  "nabilah@mail.com",
		Subjek: "Seventeen",
	}

	fmt.Printf("Ini yah %+v", user1)
}
