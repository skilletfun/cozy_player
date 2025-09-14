package middleware

import (
	"errors"
	"gcozy_player/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ValidateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id := c.Param("id"); id != "" {
			if id, err := strconv.Atoi(id); err != nil {
				response.BadRequest(c, errors.New("invalid ID"))
				return
			} else {
				c.Set("id", id)
			}
		}
		c.Next()
	}
}
