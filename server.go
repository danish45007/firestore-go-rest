package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/danish45007/go-rest/controller"
	router "github.com/danish45007/go-rest/http"
	"github.com/danish45007/go-rest/repository"
	"github.com/danish45007/go-rest/service"
	"github.com/joho/godotenv"
)

var (
	postRepo repository.PostRespositoy = repository.NewFireStoreRepo()
	services service.PostService       = service.NewPostService(postRepo)
	//muxRouter router.Router             = router.NewMuxRouter()
	control   controller.PostController = controller.NewPostController(services) //using DI to inject servies dependency into controller method
	chiRouter router.Router             = router.NewChiRouter()
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	//init middleware
	chiRouter.MIDDLEWARE()
	//test route
	chiRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode("service up and runinng")
	})
	//get post from firestore
	chiRouter.GET("/get-post", control.GetPosts)
	//create post to firestore
	chiRouter.POST("/create-post", control.CreatePost)
	URL := goDotEnvVariable("URL")
	chiRouter.SERVE(URL)
}
