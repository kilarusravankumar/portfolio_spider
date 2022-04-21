package main

import (
	"portfolio_spider/userdata"
	"regexp"

	"github.com/kataras/iris/v12"
)

var opts = iris.DirOptions{
	IndexName: "index.html",
	PushTargetsRegexp: map[string]*regexp.Regexp{
		// Match all js, css and ico files
		// from all files (recursively).
		// "/": regexp.MustCompile("((.*).js|(.*).css|(.*).ico)$"),
		// OR:
		"/":              iris.MatchCommonAssets,
		"/app2/app2app3": iris.MatchCommonAssets,
	},
	Compress: true,
	ShowList: true,
}

func main() {
	app := newApp()

	app.Listen(":8080")
}

func newApp() *iris.Application {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.HandleDir("/public", iris.Dir("./public"), opts)
	app.Get("/", index)

	return app
}

func index(ctx iris.Context) {
	user := userdata.User{
		FirstName: "Sravan",
		LastName:  "Kilaru",
		Github:    "https://github.com/kilarusravankumar",
		Twitter:   "https://twitter.com/sravan2331993",
	}
	ctx.ViewData("user", user)
	ctx.View("index.html")
}
