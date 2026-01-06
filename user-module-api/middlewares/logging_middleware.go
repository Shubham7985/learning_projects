package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		fmt.Println(
			"METHOD:", c.Request.Method,
			"PATH:", c.Request.URL.Path,
			"TIME:", time.Since(start),
		)
	}
}
