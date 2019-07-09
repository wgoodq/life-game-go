package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"log"
)

func main() {

	app := iris.Default()

	hero.Register(new(shareMemo))
	app.Get("/doCnt", hero.Handler(doCnt))

	if err := app.Run(iris.Addr(":8080")); err != nil {
		log.Print(err.Error())
	}
}

type CountService interface {
	CountVisit() int
}

type shareMemo struct {
	cnt int
}

func (sm *shareMemo) CountVisit() int {
	sm.cnt++
	return sm.cnt
}

func doCnt(service CountService) string {
	return fmt.Sprintf("<h1>Visit count: %v</h1>", service.CountVisit())
}
