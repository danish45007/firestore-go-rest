package service

import (
	"errors"
	"math/rand"

	"github.com/danish45007/go-rest/entity"
	"github.com/danish45007/go-rest/repository"
)

var (
	repo repository.PostRespositoy
)

type PostService interface {
	ValidatePost(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService(repos repository.PostRespositoy) PostService {
	repo = repos
	return &service{}
}

func (*service) ValidatePost(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindALL()
}
