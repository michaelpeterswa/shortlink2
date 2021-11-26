package handlers

import (
	"github.com/gin-gonic/gin"
)

type Host struct {
	Name string `bson:"hostname"`
}

func (envy *Environment) Hostname(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"hostname": envy.Container,
	})
}
