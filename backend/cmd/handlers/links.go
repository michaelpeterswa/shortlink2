package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/shortid"
	"github.com/michaelpeterswa/shortlink2/backend/cmd/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (envy *Environment) CreateID(c *gin.Context) {
	var url structs.LinkPostBody
	err := json.NewDecoder(c.Request.Body).Decode(&url)
	if err != nil {
		envy.Logger.Error("body decode failed", zap.Error(err))
		c.IndentedJSON(500, gin.H{
			"error": true,
		})
	}
	exists, returnedSL, err := getLinkByURLFromDB(url, envy)
	if exists {
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": true,
			})
			return
		}
		c.IndentedJSON(200, returnedSL)
		return
	}
	sl := createShortLink(url, envy)
	err = insertLinkInCache(sl, envy)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": true,
		})
		return
	}
	err = insertLinkInDB(sl, envy)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": true,
		})
		return
	}
	c.IndentedJSON(200, sl)
}

func (envy *Environment) Redirect(c *gin.Context) {
	lpb := structs.LinkPostBody{URL: c.Param("id")}
	exists, sl, err := getLinkFromCache(lpb, envy)
	if exists {
		envy.Logger.Info("cache hit in redirect", zap.Any("sl", sl))
		c.Redirect(302, sl.URL)
		return
	}
	exists, sl, err = getLinkByIDFromDB(lpb, envy)
	if !exists {
		c.IndentedJSON(200, gin.H{
			"found": false,
		})
		return
	}
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": true,
		})
		return
	}
	err = insertLinkInCache(*sl, envy)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": true,
		})
		return
	}
	c.Redirect(302, sl.URL)
}

func (envy *Environment) CheckID(c *gin.Context) {
	_, sl, err := getLinkByIDFromDB(structs.LinkPostBody{URL: c.Param("id")}, envy)
	if err != nil {
		c.IndentedJSON(200, structs.ShortLink{})
		return
	}
	c.IndentedJSON(200, sl)
}

func createShortLink(url structs.LinkPostBody, envy *Environment) structs.ShortLink {
	id := shortid.GenerateShortID()

	return structs.ShortLink{
		ID:   id,
		Link: fmt.Sprintf("%s/%s", envy.Settings.ExposedURL, id),
		URL:  url.URL,
	}
}

func insertLinkInDB(sl structs.ShortLink, envy *Environment) error {
	_, err := envy.DB.Client.Database("main").Collection("links").InsertOne(*envy.CTX, sl)
	if err != nil {
		return err
	}
	return nil
}

func insertLinkInCache(sl structs.ShortLink, envy *Environment) error {
	res, err := envy.KV.Client.Get(*envy.CTX, sl.ID).Result()
	if err == redis.Nil {
		envy.Logger.Info("Cache Miss...")
		slBytes, err := json.Marshal(sl)
		if err != nil {
			return err
		}
		envy.KV.Client.Set(*envy.CTX, sl.ID, slBytes, 6*time.Hour)
	} else if err != nil {
		return err
	} else {
		envy.Logger.Info("Cache Hit...", zap.String("key", sl.ID), zap.String("value", res))
		return nil
	}
	return nil
}

func getLinkFromCache(lpb structs.LinkPostBody, envy *Environment) (bool, *structs.ShortLink, error) {
	res, err := envy.KV.Client.Get(*envy.CTX, lpb.URL).Result()
	if err == redis.Nil {
		return false, nil, nil
	}
	var sl structs.ShortLink
	err = json.Unmarshal([]byte(res), &sl)
	if err != nil {
		return false, nil, err
	}
	return true, &sl, nil
}

func getLinkByURLFromDB(lpb structs.LinkPostBody, envy *Environment) (bool, *structs.ShortLink, error) {
	res := envy.DB.Client.Database("main").Collection("links").FindOne(*envy.CTX, bson.D{{Key: "url", Value: lpb.URL}})
	if res.Err() != nil {
		return false, nil, res.Err()
	}
	envy.Logger.Info("link found in database")
	var sl structs.ShortLink

	err := res.Decode(&sl)
	if err != nil {
		return false, nil, err
	}
	return true, &sl, nil
}

func getLinkByIDFromDB(lpb structs.LinkPostBody, envy *Environment) (bool, *structs.ShortLink, error) {
	res := envy.DB.Client.Database("main").Collection("links").FindOne(*envy.CTX, bson.D{{Key: "shortid", Value: lpb.URL}})
	if res.Err() != nil {
		return false, nil, res.Err()
	}
	envy.Logger.Info("link found in database")
	var sl structs.ShortLink

	err := res.Decode(&sl)
	if err != nil {
		return false, nil, err
	}
	return true, &sl, nil
}
