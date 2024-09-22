package model

import "fmt"

type Auth struct {
	Login          string `json:"login"`
	Password       string `json:"password"`
	IP             string `json:"ip"`
	loginBucket    string
	passwordBucket string
	ipBucket       string
}

func (a Auth) LoginBucket() string {
	if a.loginBucket != "" {
		return a.loginBucket
	}

	return fmt.Sprintf("login:%s", a.Login)
}

func (a Auth) PasswordBucket() string {
	if a.passwordBucket != "" {
		return a.passwordBucket
	}

	return fmt.Sprintf("password:%s", a.Password)
}

func (a Auth) IPBucket() string {
	if a.ipBucket != "" {
		return a.ipBucket
	}

	return fmt.Sprintf("ip:%s", a.IP)
}
