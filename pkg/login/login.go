package login

import (
	"log"
)

func Login() {
	e, p := "email@email.com", "password"
	c, s, m := loginToLobby(e, p)
	log.Println(c)
	log.Println(s)
  log.Println(m)
}
