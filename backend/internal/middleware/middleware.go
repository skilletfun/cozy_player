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
