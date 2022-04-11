package login

import (
	"log"
)

func Login() {
	e, p, gw := "email@email.com", "password", "comx"
	c, s, m := loginToLobby(e, p)
  loginToGameworld(c, s, m, gw)
	log.Println(c)
	log.Println(s)
  log.Println(m)
}
