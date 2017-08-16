package dispatcher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
		g.GET("/osu-osz2-getscores.php", getScores)
		g.POST("/osu-submit-modular.php", submitModular)
	}

	return c
}()

func banchoConnect(c *gin.Context) {
	c.String(200, "us")
}

func getScores(c *gin.Context) {
	var scores string

	// https://zxq.co/ripple/lets/src/master/objects/beatmap.pyx#L266-L283
	scores += "2|false\n" // 2 = ranked, false = we don't have osz2
	scores += "0\n\n10\n"

	// No personal best
	scores += "\n"

	// https://zxq.co/ripple/lets/src/master/objects/score.pyx#L182-L199
	scores += "1|This is|4003251|382|0|0|241|0|0|44|1|0|511||1502876163|0\n"
	scores += "2|an illusion|3414400|331|0|31|210|0|10|44|0|0|512||1502866163|0\n"
	scores += "3|these scores|3404400|331|0|31|210|0|10|44|0|0|513||1502856163|0\n"
	scores += "4|do not exist|3394400|331|0|31|210|0|10|44|0|0|514||1502846163|0\n"

	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteString(scores)
}

func submitModular(c *gin.Context) {
	// Have you ever wondered why the fuck it's called "modular" when it's
	// not fucking modular?
	// Anyway, this function does not quite work because I fucking hate PHP.
	// https://stackoverflow.com/q/45712450/5328069
	// I'm giving up until somebody smarter than me comes up with a solution
	// on stackoverflow.

	score := decodePost(c.PostForm("score"))
	iv := decodePost(c.PostForm("iv"))
	pass := c.PostForm("pass")

	aesKey := getAESKey(c.PostForm("osuver"))
	baseAES, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	block := cipher.NewCBCDecrypter(baseAES, []byte(iv)[16:])
	block.CryptBlocks(score, score)

	fmt.Println("We have a score, we just don't know what to do with it yet.")
	_ = pass
	c.String(200, "")
}

// https://github.com/osuripple/lets/blob/4322e367705bc199a06861a659df42c3e76fcd9f/handlers/submitModularHandler.pyx#L66-L70
func getAESKey(s string) string {
	if s == "" {
		return "h89f2-890h2h89b34g-h80g134n90133"
	}
	return "osu!-scoreburgr---------" + s
}

func decodePost(data string) []byte {
	decoded, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, strings.NewReader(data)))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return decoded
}
