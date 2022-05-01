package main

import (
	"portfolio_spider/input_reader"
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
	mastersEducation := userdata.Education{
		SchoolName: "Northwest Missouri State University",
		Degree:     "Masters in Applied Computer science",
		GradDate:   "2015",
	}
	bachelorsEducation := userdata.Education{
		SchoolName: "Gandhi Institute of Technology",
		Degree:     "Bachelors in Computer science",
		GradDate:   "2014",
	}

	education := []userdata.Education{mastersEducation, bachelorsEducation}

	description := []string{}

	description = append(description, "Build and maintain a data platform by utilizing Go, docker, and MongoDB for the department of pathology. Rendered analytical support in maintaining and upgrading existing department of pathology website using React and PHP.")
	description = append(description, "Designed and deployed scalable web applications by using React SSR for the University of Michigan and vendors")
	description = append(description, "Supported team members in pathology specimen tracking tool built by using react, Node, Express, web Sockets, and	docker")

	work1 := userdata.WorkExperience{
		CompanyName: "Michigan Medicine",
		JobTitle:    "Application Programmer",
		StartDate:   "October 2018",
		EndDate:     "Present",
		Location:    "Ann Arbor, MI",
		Description: description,
	}

	user := userdata.User{
		FirstName: input_reader.ConsoleInput("Enter your first name"),
		LastName:  input_reader.ConsoleInput("Enter your last name"),
		Github:    "https://github.com/kilarusravankumar",
		Twitter:   "https://twitter.com/sravan2331993",
		Education: education,
		Work:      []userdata.WorkExperience{work1},
	}

	ctx.ViewData("user", user)
	ctx.View("index.html")
}
