package database

import (
	"context"
	"fmt"
	"gin-boiler/internal/config"
	"log"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client
var MongoDatabase *mongo.Database

func MongoConnect(env config.Env) (*mongo.Client, error) {
	uri := env.MongoConfig.Url
	log.Printf("MongoDB URI: %s", uri)

	// MongoDB 클라이언트 옵션 설정
	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(10 * time.Second).
		SetMaxPoolSize(100). // 최대 연결 풀 크기
		SetMinPoolSize(10)   // 최소 연결 풀 크기

	// MongoDB 연결
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB 연결 실패: %v", err)
	}

	// 연결 테스트
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB 핑 실패: %v", err)
	}

	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("MongoDB URI 파싱 실패: %v", err)
	}
	dbName := strings.TrimPrefix(parsedURI.Path, "/")
	if dbName == "" {
		return nil, fmt.Errorf("MongoDB 데이터베이스 이름이 없습니다")
	}

	MongoDB = client
	MongoDatabase = client.Database(dbName)
	log.Printf("MongoDB 데이터베이스 연결 성공")
	return client, nil
}

// MongoDB 연결 종료
func DisconnectMongo() error {
	if MongoDB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return MongoDB.Disconnect(ctx)
	}
	return nil
}
