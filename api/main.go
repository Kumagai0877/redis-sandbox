package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var rClient *redis.Client

func main() {
	e :=  gin.Default()
	e.Use(cors.Default())
	e.GET("/", getKey)

	if err :=e.Run(":8081"); err != nil {
		panic(err)
	}
}

func getKey(c *gin.Context) {
	redisConnection()
	ctx := c.Request.Context()
	key, err := c.GetQuery("key")
	if !err {
		c.JSON(http.StatusBadRequest, "must be key")
		return
	}
	res, e := getCachedKey(ctx, key)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	c.JSON(http.StatusOK, res)
}

func redisConnection() {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func getCachedKey(ctx context.Context, key string) (string, error) {
	res, err := rClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

