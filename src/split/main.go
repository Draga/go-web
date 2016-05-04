package split

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/", home)
	iris.Listen(":8080")
}
