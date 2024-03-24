package controllers

import (
	"github.com/gin-gonic/gin"
)

func (idb *InDB) store(c *gin.Context) {
	c.JSON(200, "fuck")
}
