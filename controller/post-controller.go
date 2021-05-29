package controller

import (
	"encoding/json"
	"net/http"

	"github.com/danish45007/go-rest/entity"
	"github.com/danish45007/go-rest/error"
	"github.com/danish45007/go-rest/service"
)

var (
	services service.PostService
	posts    []entity.Post
)

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	CreatePost(res http.ResponseWriter, req *http.Request)
}

// type controller struct{}

func NewPostController(service service.PostService) PostController {
	services = service
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
