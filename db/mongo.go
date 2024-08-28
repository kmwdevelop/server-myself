package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

func InitMongoDB(uri string) {
	var err error
	ctx := context.Background()

	// MongoDB 클라이언트 연결 설정
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	pingErr := client.Ping(ctx, readpref.Primary())
	if pingErr != nil {
		panic(pingErr)
	}

	fmt.Println("Connected to MongoDB")
}

func DisconnectMongoDB() {
	ctx := context.Background()
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Disconnected MongoDB")
}

func CreateCollection(collectionName string) {

	// 데이터베이스 선택
	database := client.Database("cluster0")

	// 명시적으로 컬렉션 생성
	collectionErr := database.CreateCollection(context.Background(), collectionName)
	if collectionErr != nil {
		panic(collectionErr)
	}

	fmt.Println("Created collection: " + collectionName)
}

func InsertOne(collectionName string) {
	// 데이터베이스 선택
	database := client.Database("cluster0")
	collection := database.Collection(collectionName)

	// 샘플 데이터 생성
	sampleData := bson.D{
		{"name", "John Doe"},
		{"age", 29},
		{"email", "johndoe@example.com"},
		{"address", bson.D{
			{"street", "123 Main St"},
			{"city", "Anytown"},
			{"state", "CA"},
			{"zip", "12345"},
		}},
		{"created_at", time.Now()},
	}

	insertResult, err := collection.InsertOne(context.Background(), sampleData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
