package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"lgg/domain"
)

func main() {
	app := iris.New()
	html := iris.HTML("./templates", ".html")
	html.Reload(true)
	app.RegisterView(html)

	app.Get("/", note)

	hero.Register(new(domain.LifeGame))
	app.Get("/lifegame", hero.Handler(lifeGame))

	app.Run(iris.Addr(":8080"))
}

func note(ctx iris.Context) {
	ctx.ViewData("msg", "Life Game")
	ctx.View("index.html")
}

func lifeGame(service domain.LifeGameService) hero.View {
	cnt, chessboard := service.RunGame()

	return hero.View{
		Name: "lifegame.html",
		Data: map[string]interface{}{
			"cnt":        cnt,
			"chessboard": chessboard,
		},
	}
}
