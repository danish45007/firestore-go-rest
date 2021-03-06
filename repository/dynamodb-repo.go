package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/danish45007/go-rest/entity"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type dynamoDBRepo struct {
	tableName string
}

func NewDynamoDBRepository() PostRespositoy {
	return &dynamoDBRepo{
		tableName: "posts",
	}
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Create AWS Session
	ACCESS_KEY := goDotEnvVariable("access_key_ID")
	SECRET_KEY := goDotEnvVariable("secret_access_key")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
	})
	if err != nil {
		panic("unable to create aws session")
	}

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo *dynamoDBRepo) Save(post *entity.Post) (*entity.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Transforms the post to map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return nil, err
	}

	// Create the Item Input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}

	// Save the Item into DynamoDB
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return post, err
}

func (repo *dynamoDBRepo) FindALL() ([]entity.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}
	var posts []entity.Post = []entity.Post{}
	for _, i := range result.Items {
		post := entity.Post{}

		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *dynamoDBRepo) FindByID(id string) ([]entity.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()
	var posts []entity.Post
	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	post := entity.Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}
	posts = append(posts, post)
	return posts, nil
}
