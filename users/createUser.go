package users

import (
	"fmt"
	"time"

	"github.com/MALPV/Tutorial-Godesde0/modelos"
)

func CreateAdmin() {

	admin := new(modelos.User)
	admin.AddUser(1, "Admin", time.Now(), true)
	fmt.Println(admin)

}
