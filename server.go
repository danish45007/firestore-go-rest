package main

import (
	"github.com/danish45007/go-rest/controller"
	router "github.com/danish45007/go-rest/http"
)

var (
	route   router.Router             = router.NewMuxRouter()
	control controller.PostController = controller.NewPostController()
)

func main() {
	route.GET("/get-post", control.GetPosts)
	route.POST("/create-post", control.CreatePost)
	route.SERVE("127.0.0.1:8080")
}
