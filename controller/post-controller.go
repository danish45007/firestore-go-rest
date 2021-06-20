package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/danish45007/go-rest/cache"
	"github.com/danish45007/go-rest/entity"
	"github.com/danish45007/go-rest/error"
	"github.com/danish45007/go-rest/service"
)

var (
	services  service.PostService
	posts     []entity.Post
	postCache cache.PostCache
)

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	CreatePost(res http.ResponseWriter, req *http.Request)
	GetPostByID(res http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	services = service
	postCache = cache
	return &controller{}
}

func (*controller) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	posts, err := services.FindAll()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: "Error while getting the posts"})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)

}

func (*controller) CreatePost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: "Error unmarshalling request data"})

	}
	// validation of post
	validateErorr := services.ValidatePost(&post)
	if validateErorr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: validateErorr.Error()})
	}
	data, err := services.Create(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: "Error saving the posts"})
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]
	var post *entity.Post = postCache.Get(postID)
	// when not stored cache or expired
	if post == nil {
		post, err := services.FindById(postID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(error.ServiceError{Message: "No posts found!"})
			return
		}
		//store post into cache
		p := post[0]
		postCache.Set(postID, &p)
		fmt.Printf("Post %v set into redis", p)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(p)
	} else {
		fmt.Println("Getting from redis")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}

}
