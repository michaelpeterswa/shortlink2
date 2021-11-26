package shortid

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/mongo"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/redis"
)

type shortID struct {
	db *mongo.MongoConn
	kv *redis.RedisConn
}

var shortIDConns *shortID

func InitShortID(db *mongo.MongoConn, kv *redis.RedisConn) {
	shortIDConns = &shortID{db, kv}
}

func GenerateShortID() string {
	id, err := gonanoid.New(7)
	if err != nil {
		log.Println("failed to generate shortID")
	}
	return id
}
