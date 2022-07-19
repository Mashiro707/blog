package bcrypt

import (
	"fmt"
	"log"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	pwd, err := PasswordHash("123456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pwd)
	if bool := PasswordVerify("123456", pwd); bool != true {
		log.Fatal(err)
	}
}
