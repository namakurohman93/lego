package main

import (
	"log"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/tkapi"
)

func main() {
	// t1, err := time.Parse(time.RFC1123, "Sun, 24 Apr 2022 08:04:39 GMT")
	// t1, err := time.Parse(time.RFC1123, strings.ReplaceAll("Sun, 24-Apr-2022 08:04:39 GMT", "-", " "))
	// if err != nil {
	// log.Fatal(err)
	// }
	// st := t1.Format(time.RFC3339)
	// t2, err := time.Parse(time.RFC3339, st)
	// if err != nil {
	// log.Fatal(err)
	// }
	// yesterday := t2.AddDate(0, 0, -1)
	// n := time.Now()
	// log.Println(yesterday.Before(n))
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	gs, err := tkapi.GetGameSession(&c)
	if err != nil {
		log.Fatal(err)
	}
	villages, err := tkapi.RequestOwnVillage(c.Gameworld,
		gs.GameworldSession, gs.GetGameCookie())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(villages)
}
