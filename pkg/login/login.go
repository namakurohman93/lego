package login

import (
	"log"
)

func Login() {
	e, p := "email@email.com", "password"
	c, s := loginToLobby(e, p)
	log.Println(c)
	log.Println(s)
}
