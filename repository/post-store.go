package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/danish45007/go-rest/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	projectId      string = "go-fire-c8913"
	collectionName string = "posts"
)

type PostRespositoy interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindALl() ([]entity.Post, error)
}

type repo struct{}

// constructor NewPostRepo
func NewPostRepo() PostRespositoy {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:\\Users\\DANISH\\go\\src\\go-rest\\repository\\fire.json")
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Printf("Failed to create Firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	// get the collection instance
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Printf("Failed creating post %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindALl() ([]entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:\\Users\\DANISH\\go\\src\\go-rest\\repository\\fire.json")
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Printf("Failed to create Firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	fireIterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := fireIterator.Next()
		if err != nil {
			log.Printf("Failed to iterate the list of post %v", err)
			return nil, err
		}

		if err == iterator.Done {
			break
		}

		post := entity.Post{
			Id:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil

}