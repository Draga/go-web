package main

import (
	"encoding/xml"
	"github.com/kataras/iris"
)

type ExampleXml struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

type Model struct{
	Name string
	Age uint32
}

func main() {
	iris.Get("/data", func(ctx *iris.Context) {
		ctx.Data(iris.StatusOK, []byte("Some binary data here."))
	})

	iris.Get("/text", func(ctx *iris.Context) {
		ctx.Text(iris.StatusOK, "Plain text here")
	})

	iris.Get("/json", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, map[string]string{"hello": "json"})
	})

	iris.Get("/jsonp", func(ctx *iris.Context) {
		ctx.JSONP(iris.StatusOK, "callbackName", map[string]string{"hello": "jsonp"})
	})

	iris.Get("/xml", func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, ExampleXml{One: "hello", Two: "xml"})
	})

	iris.Get("/html", func(ctx *iris.Context) {
		// Assumes you have a template in ./templates called "example.html".
		// $ mkdir -p templates && echo "<h1>Hello HTML world.</h1>" > templates/example.html
		ctx.HTML(iris.StatusOK, "example", Model{Name:"anem", Age: 21})
	})

	// ctx.Render is the same as ctx.HTML but with default 200 status OK
	iris.Get("/html2", func(ctx *iris.Context) {
		// Assumes you have a template in ./templates called "example.html".
		// $ mkdir -p templates && echo "<h1>Hello HTML world.</h1>" > templates/example.html
		ctx.Render("example", nil)
	})

	iris.Listen(":8080")
}