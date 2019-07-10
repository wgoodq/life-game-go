package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/websocket"
	"lgg/domain"
	"strings"
	"time"
)

func main() {
	app := iris.New()

	// 加载静态资源
	html := iris.HTML("./templates", ".html")
	html.Reload(true)
	app.RegisterView(html)

	// 一般模式
	app.Get("/", note)

	// 缓存模式
	hero.Register(new(domain.LifeGame))
	app.Get("/lifegame", hero.Handler(lifeGame))

	// WebSocket模式
	app.Get("/ws", func(ctx context.Context) {
		ctx.View("websockets.html")
	})
	setupWebsocket(app)

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

func setupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		// These are low-level optionally fields,
		// user/client can't see those values.
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// only javascript client-side code has the same rule,
		// which you serve using the ws.ClientSource (see below).
		EvtMessagePrefix: []byte("my-custom-prefix:"),
	})

	ws.OnConnection(handleConnection)

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/lifegamews", ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(ws.ClientSource)
	})
}

func handleConnection(conn websocket.Connection) {

	lg := new(domain.LifeGame)

	for i := 0; i >= 0; i++ {
		cnt, chessCells := lg.RunGame()
		conn.Emit("cnt", fmt.Sprintf("%v", cnt))
		conn.Emit("chessBoard", strings.Join(chessCells, " "))

		time.Sleep(time.Second)
	}

}
