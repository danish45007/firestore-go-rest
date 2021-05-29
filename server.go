package main

import (
	"encoding/json"
	"net/http"

	"github.com/danish45007/go-rest/controller"
	router "github.com/danish45007/go-rest/http"
	"github.com/danish45007/go-rest/service"
)

var (
	// postRepo repository.PostRespositoy = repository.NewFireStoreRepo()
	// services service.PostService       = service.NewPostService(postRepo)

	//muxRouter router.Router = router.NewMuxRouter()
	// control   controller.PostController = controller.NewPostController(services) //using DI to inject servies dependency into controller method
	services  service.CarDetailsService       = service.NewCarDetailsController()
	control   controller.CarDetailsController = controller.NewCarDetailsController(services)
	chiRouter router.Router                   = router.NewChiRouter()
)

func main() {
	//init middleware
	chiRouter.MIDDLEWARE()
	//test route
	chiRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode("service up and runinng")
	})
	/*
		//get post from firestore
		chiRouter.GET("/get-post", control.GetPosts)
		//create post to firestore
		chiRouter.POST("/create-post", control.CreatePost)
		//api-mashup route
	*/
	chiRouter.GET("/car-details", control.GetCarDetails)
	chiRouter.SERVE("127.0.0.1:8080")
}
