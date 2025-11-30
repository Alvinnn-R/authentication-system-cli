package handler

import (
	"authentication-system-cli/service"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserHandler struct {
	service *service.UserService
	scanner *bufio.Scanner
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		service: svc,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (h *UserHandler) ShowMenu() {
	fmt.Println("\n=== SIMPLE LOGIN SYSTEM ===")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Print("Pilih menu: ")
}

func (h *UserHandler) HandleRegister() {
	fmt.Println("\n--- REGISTER ---")

	fmt.Print("Full Name\t: ")
	h.scanner.Scan()
	fullName := strings.TrimSpace(h.scanner.Text())

	fmt.Print("Email\t\t: ")
	h.scanner.Scan()
	email := strings.TrimSpace(h.scanner.Text())

	fmt.Print("Phone\t\t: ")
	h.scanner.Scan()
	phoneNumber := strings.TrimSpace(h.scanner.Text())

	fmt.Print("Password\t: ")
	h.scanner.Scan()
	password := strings.TrimSpace(h.scanner.Text())

	err := h.service.Register(fullName, email, phoneNumber, password)
	if err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", err.Error())
		return
	}

	fmt.Println("Registrasi berhasil! Data tersimpan di users.json")
}

func (h *UserHandler) HandleLogin() {
	fmt.Println("\n--- LOGIN ---")

	fmt.Print("Email\t\t: ")
	h.scanner.Scan()
	email := strings.TrimSpace(h.scanner.Text())

	fmt.Print("Password\t: ")
	h.scanner.Scan()
	password := strings.TrimSpace(h.scanner.Text())

	user, err := h.service.Login(email, password)
	if err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", err.Error())
		return
	}

	fmt.Printf("\033[32mLogin berhasil, selamat datang %s\033[0m\n", user.FullName)
}
