
/*
package service

import (
	"testing"

	"github.com/danish45007/go-rest/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *mockRepository) FindALL() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.ValidatePost(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "B"}

	testService := NewPostService(nil)

	err := testService.ValidatePost(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	post := entity.Post{Id: 1, Title: "A", Text: "B"}
	mockRepo := new(mockRepository)
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestSavePost(t *testing.T) {
	post := entity.Post{Title: "A", Text: "B"}
	mockRepo := new(mockRepository)
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)
	res, err := testService.Create(&post)
	mockRepo.AssertExpectations(t)
	assert.NotNil(t, res.Id)
	assert.Equal(t, "A", res.Title)
	assert.Equal(t, "B", res.Text)
	assert.Nil(t, err)

}
*/