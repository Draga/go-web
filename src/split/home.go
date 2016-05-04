package split

import "github.com/kataras/iris"

func home(c *iris.Context) {
	c.Write("From %s", c.PathString())
}
