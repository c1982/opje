package main

import (
	"fmt"
	"opje"
)

type AuthService interface {
	Login(username, password string) (bool, error)
	Register(username, password string) (bool, error)
}

type authentication struct{}

func newAuthService() AuthService {
	return &authentication{}
}

func (a *authentication) Login(username, password string) (bool, error) {
	return true, nil
}

func (a *authentication) Register(username, password string) (bool, error) {
	return true, nil
}

func init() {
	locator.Register(newAuthService())
}

func main() {
	auth, err := locator.Resolve[AuthService]()
	if err != nil {
		panic("authService not registered")
	}

	loggedin, err := auth.Login("admin", "admin")
	if err != nil {
		panic("login failed")
	}

	if loggedin {
		fmt.Println("Logged in ✌️")
	}
}
