// main.go

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"rinha-basic/cache"
	"rinha-basic/repositories"
	"rinha-basic/usecases"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:password@(127.0.0.1:3306)/rinhadev?charset=utf8mb4,utf8\\u0026readTimeout=30s\\u0026writeTimeout=30s&parseTime=true")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	cacheStorage := cache.NewCacheStorage(rdb)
	repo := repositories.NewPersonRepository(db)
	usecase := usecases.NewPersonUsecase(repo, cacheStorage)

	r := gin.Default()

	SetupPersonRoutes(r, usecase)

	r.Run(":8080")
}
