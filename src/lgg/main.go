package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"lgg/domain"
)

func main() {
	app := iris.New()

	app.Get("/", note)

	hero.Register(new(domain.LifeGame))
	app.Get("/lifegame", hero.Handler(lifeGame))

	_ = app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func note(ctx iris.Context) {
	_, _ = ctx.HTML("<h1>Life Game</h1>")
	_, _ = ctx.Writef("<a href='%s'>Start Game</a>", "http://localhost:8080/lifegame")
}

func lifeGame(service domain.LifeGameService) string {
	cnt, cellsString := service.RunGame()
	return fmt.Sprintf("<h1 style=\"text-align:center\">Current Generation: %v</h1><br><p style=\"text-align:center\">%v</p>", cnt, cellsString)
}
