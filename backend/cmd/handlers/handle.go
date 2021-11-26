package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/mongo"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/redis"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/structs"
	"go.uber.org/zap"
)

type Environment struct {
	CTX       *context.Context
	Settings  *structs.Settings
	KV        *redis.RedisConn
	DB        *mongo.MongoConn
	Logger    *zap.Logger
	Container string
}

func Handle(r *gin.Engine, envy *Environment) {
	r.GET("/healthcheck", envy.Healthcheck)
	r.GET("/hostname", envy.Hostname)
	r.POST("/create", envy.CreateID)
	r.GET("/id/:id", envy.CheckID)
	r.GET("/:id", envy.Redirect)
}
