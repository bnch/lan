package dispatcher

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// osu.ppy.sh server

// scoreServer handles everything regarding the score server.
var scoreServer = func() http.Handler {
	c := gin.Default()

	g := c.Group("/web")
	{
		g.GET("/bancho_connect.php", banchoConnect)
		g.POST("/osu-error.php", func(c *gin.Context) {
			c.String(200, "")
		})
		g.GET("/check-updates.php", func(c *gin.Context) {
			c.JSON(200, []struct{}{})
		})
		g.GET("/lastfm.php", func(c *gin.Context) {
			c.String(200, "")
		})
	}

	return c
}()

func banchoConnect(c *gin.Context) {
	c.String(200, "us")
}
