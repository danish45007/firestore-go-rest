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

type repo struct{}

// constructor NewFireStoreRepo
func NewFireStoreRepo() PostRespositoy {
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

func (*repo) FindALL() ([]entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:\\Users\\DANISH\\go\\src\\go-rest\\repository\\fire.json")
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Printf("Failed to create Firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
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

func (*repo) FindByID(id int) ([]entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:\\Users\\DANISH\\go\\src\\go-rest\\repository\\fire.json")
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Printf("Failed to create Firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	data := client.Collection(collectionName).Where("ID", "==", id).Documents(ctx)
	for {
		doc, err := data.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
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
