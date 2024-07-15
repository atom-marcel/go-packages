package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"
	"strconv"
	"strings"

	"MMProd/maureralphabet"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", index)

	router.Run(":8080")
}

func index(c *gin.Context) {
	text, _ := c.GetQuery("text")
	plot_size_str, _ := c.GetQuery("plot_pixel")
	lw_str, _ := c.GetQuery("line_pixel")

	text = strings.ReplaceAll(text, "\r\n", "\n")

	plot_size, _ := strconv.Atoi(plot_size_str)
	lw, _ := strconv.Atoi(lw_str)

	if plot_size <= 0 || lw <= 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "Pixel für ein Zeichen oder Linienstärke darf nicht null oder geringer sein!",
		})
		return
	}

	max := 0
	for _, t := range strings.Split(text, "\n") {
		if len(t) > max {
			max = len(t)
		}
	}

	buchstaben, err := maureralphabet.ConvertToBuchstaben(text)

	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": err,
		})
		return
	}

	im := maureralphabet.ImageBuchstaben(buchstaben, maureralphabet.ImageOptions{
		Padding:    12,
		Stride:     10,
		Wordlength: max,
		Lines:      len(strings.Split(text, "\n")),
		PixSize:    lw,
		PixPlot:    plot_size,
	})

	var buf bytes.Buffer
	err = png.Encode(&buf, im)

	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": err,
		})
		return
	}

	encode := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": text,
		"img": encode,
	})
}
