package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HealthCheck struct {
	Health string `bson:"health"`
}

func (envy *Environment) Healthcheck(c *gin.Context) {
	mainDB := envy.DB.Client.Database("main")
	healthColl := mainDB.Collection("healthcheck")
	health := healthColl.FindOne(*envy.CTX, bson.D{{Key: "health", Value: "ok"}})

	if health.Err() == mongo.ErrNoDocuments {
		c.IndentedJSON(500, gin.H{
			"health": "unhealthy",
		})
		return
	}

	var healthStatus *HealthCheck
	err := health.Decode(&healthStatus)
	if err != nil {
		envy.Logger.Error("decoding health status failed")
		c.IndentedJSON(500, gin.H{
			"health": "unhealthy",
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"health": healthStatus.Health,
	})
}
