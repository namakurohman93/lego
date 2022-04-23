package main

import (
	"log"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/tkapi"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	driver, err := tkapi.NewDriver(&c)
	if err != nil {
		log.Fatal(err)
	}
	// vs, err := driver.RequestOwnVillage()
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(vs)
	m, err := driver.RequestMap()
	if err != nil {
		log.Fatal(err)
	}
	// temp := m.Cells.GetRobberHideouts()
	// log.Println(temp)
	// if player, ok := m.Player.GetByName("mightbenotfound"); !ok {
	// log.Println("player not found")
	// } else {
	// log.Println("player found")
	// log.Println(player)
	// }
	if kingdom, ok := m.Kingdom.GetByName("mightbenotfound"); !ok {
		log.Println("kingdom not found")
	} else {
		log.Println("kingdom found")
		log.Println(kingdom)
	}
}
