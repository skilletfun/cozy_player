package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, result any, err error) {
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"detail": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, result)
}

func NotFound(c *gin.Context, err error) {
	c.JSON(
		http.StatusNotFound,
		gin.H{"detail": err.Error()},
	)
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{"detail": err.Error()},
	)
}

func ErrorOrNoContent(c *gin.Context, err error) {
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"detail": err.Error()},
		)
		return
	}

	c.Status(http.StatusNoContent)
}

func File(c *gin.Context, file string) {
	c.File(file)
}

func CachedFile(c *gin.Context, file string) {
	c.Header("Cache-Control", "public, max-age=31536000")
	c.File(file)
}

func Data(c *gin.Context, mimeType string, data []byte) {
	c.Data(200, mimeType, data)
}
