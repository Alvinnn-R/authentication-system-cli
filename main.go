package main

import (
	"authentication-system-cli/handler"
	"authentication-system-cli/repository"
	"authentication-system-cli/service"
	"fmt"
)

func main() {
	repo := repository.NewUserRepository("data/users.json")
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	for {
		h.ShowMenu()

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			h.HandleRegister()
		case "2":
			h.HandleLogin()
		case "3":
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
