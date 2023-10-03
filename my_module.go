package go_for_all

import (
	"fmt"

	"github.com/expx-tech/go_for_all/app/service"
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
}
