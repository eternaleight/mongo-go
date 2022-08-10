package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type demo struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {
	// コンテキストの作成
	// バックグラウンドで接続する。タイムアウトは5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// 関数を抜けたらクローズするようにする
	defer cancel()

	// .envfile
	err := godotenv.Load(".env")
	message := os.Getenv("MONGO_URL")

	// Connect to MongoDB
	// 指定したURIに接続する
	client, err := mongo.NewClient(options.Client().ApplyURI(message))

	// Close connection
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// DBにPingする
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("connection error:", err)
	} else {
		fmt.Println("connection success!")
	}

	_ = client.Database("grpc").Collection("test")

}
