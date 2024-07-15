package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qpliu/qrencode-go/qrencode"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", index)

	router.Run(":8080")
}

func index(c *gin.Context) {
	msg, _ := c.GetQuery("msg")

	if msg != "" {
		grid, err := qrencode.Encode(msg, qrencode.ECLevelQ)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"msg": err,
			})
			return
		}

		var buf bytes.Buffer
		err = png.Encode(&buf, grid.Image(8))
		encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"msg": err,
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": msg,
			"qr":  encoded,
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
