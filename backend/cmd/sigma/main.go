package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/michaelpeterswa/sigma/backend/cmd/mongo"
	"github.com/michaelpeterswa/sigma/backend/cmd/redis"
	"github.com/michaelpeterswa/sigma/backend/cmd/structs"
	"github.com/michaelpeterswa/sigma/backend/cmd/util"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v2"
)

var settings structs.Settings
var redisConn *redis.RedisConn
var mongoConn *mongo.MongoConn

var ctx = context.Background()

func main() {
	fmt.Println("sigma")
	util.Info()

	fileSettings, err := os.ReadFile("settings.yaml")
	if err != nil {
		log.Println("Error loading settings.yaml file")
	}

	err = yaml.Unmarshal(fileSettings, &settings)
	if err != nil {
		log.Println("Error unmarshalling settings")
	}

	redisConn = redis.InitRedis(settings)
	mongoConn = mongo.InitMongo(ctx, settings)
	collection := mongoConn.Client.Database("main").Collection("versions")
	fmt.Println(collection.Name())

	redisConn.Client.Set(ctx, "count", 0, 0)
	value := redisConn.Client.Get(ctx, "count")
	fmt.Printf("Value of 'lmao': %s\n", value.Val())

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/count", func(c *gin.Context) {
		redisConn.Client.Incr(ctx, "count")
		c.JSON(200, gin.H{
			"count": redisConn.Client.Get(ctx, "count").Val(),
		})
	})
	r.GET("/versions", func(c *gin.Context) {
		cur, currErr := collection.Find(ctx, bson.D{})

		if currErr != nil {
			panic(currErr)
		}
		defer cur.Close(ctx)

		var apps []structs.Version
		if err = cur.All(ctx, &apps); err != nil {
			panic(err)
		}

		c.JSON(200, apps)
	})

	runerr := r.Run()
	if runerr != nil {
		panic(runerr)
	}
}
