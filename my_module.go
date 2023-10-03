package go_for_all

import (
	"fmt"

<<<<<<< HEAD
func Version() {
	fmt.Println("Version v1.1.0")
=======
	"github.com/expx-tech/go_for_all/v2/app/service"
)

func Request(token, method string) {
	if token != "супер_секрет" {
		fmt.Println("Пользователь не авторизован")
		return
	}

	switch method {
	case "GetUsers":
		service.GetUsers()
	default:
		fmt.Println("Нет такого метода")
	}
>>>>>>> c4f7e6e (сервис)
}
