package repository

import "github.com/danish45007/go-rest/entity"

type PostRespositoy interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindALL() ([]entity.Post, error)
	FindByID(id string) ([]entity.Post, error)
}
