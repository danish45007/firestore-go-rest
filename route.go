package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/danish45007/go-rest/entity"
	"github.com/danish45007/go-rest/repository"
)

var (
	repo  repository.PostRespositoy = repository.NewPostRepo()
	posts []entity.Post
)

// func init() {
// 	posts = []entity.Post{{
// 		Id:    1,
// 		Title: "First Post",
// 		Text:  "This is the first post",
// 	}}
// }

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindALl()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error getting the post array"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)

}

func createPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling request data"}`))
	}
	post.Id = rand.Int63()
	repo.Save(&post)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}
