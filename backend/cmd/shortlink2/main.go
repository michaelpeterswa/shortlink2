package main

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/handlers"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/hostname"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/logging"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/mongo"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/redis"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/shortid"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/util"
)

var envy handlers.Environment

var ctx = context.Background()

func main() {

	envy.CTX = &ctx
	envy.Settings = util.GetSettings()
	envy.KV = redis.InitRedis(envy.Settings)
	envy.DB = mongo.InitMongo(ctx, envy.Settings)
	envy.Logger = logging.InitLogger(envy.Settings.ZapLevel)
	envy.Container = hostname.InitHostname()

	shortid.InitShortID(envy.DB, envy.KV)

	gin.SetMode(envy.Settings.GinMode)
	r := gin.New()
	r.Use(cors.Default())
	r.Use(ginzap.Ginzap(envy.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(envy.Logger, true))

	handlers.Handle(r, &envy)

	runerr := r.Run()
	if runerr != nil {
		panic(runerr)
	}
}
