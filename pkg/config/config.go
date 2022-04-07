package config

import (
	"LoginFish/pkg/model"
	"io/ioutil"
	"strings"
)

var AdminPath = "admin"

func GetPassword() string {
	pwd, err := ioutil.ReadFile("password.txt")
	if err == nil {
		return strings.TrimSpace(string(pwd))
	}
	p := model.RandStringRunes(16)
	ioutil.WriteFile("password.txt", []byte(p), 0755)
	return strings.TrimSpace(p)
}
