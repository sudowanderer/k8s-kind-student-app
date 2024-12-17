package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

// 初始化 MongoDB 连接
func InitMongoDB() *mongo.Client {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://mongodb:27017").SetTimeout(10 * time.Second))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	return client
}

// 获取集合
func GetCollection(name string) *mongo.Collection {
	return client.Database("student_management").Collection(name)
}
