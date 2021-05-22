package service

import (
	"testing"

	"github.com/danish45007/go-rest/entity"
	"github.com/stretchr/testify/assert"
)

func testValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.ValidatePost(nil) // passing no posts
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post is empty")
}

func testValidateTitle(t *testing.T) {
	post := entity.Post{Id: 123, Title: "Test Title", Text: "Test Text"}
	testService := NewPostService(nil)
	err := testService.ValidatePost(&post)
	assert.NotNil(t, err)
	assert.Empty(t, err.Error(), "The post title is empty")
}
