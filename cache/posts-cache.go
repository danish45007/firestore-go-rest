package cache

import "github.com/danish45007/go-rest/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
