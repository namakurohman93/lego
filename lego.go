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
	m, err := driver.RequestMap()
	if err != nil {
		log.Fatal(err)
	}
	temp := m.GetRobberHideouts()
	log.Println(temp)
}
