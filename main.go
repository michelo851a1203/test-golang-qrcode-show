package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func MainHandler(ctx *gin.Context) {
	var png []byte
	inputInfo := "色慾之罪 way"
	png, err := qrcode.Encode(inputInfo, qrcode.Medium, 256)
	if err != nil {
		log.Fatalf("this is qrCode error %v\n", err)
	}

	base64String := base64.StdEncoding.EncodeToString(png)
	fmt.Println(base64String)

	imageHtml := fmt.Sprintf("<img src=\"data:image/png;base64,%s\" />", base64String)
	ctx.Data(http.StatusOK, "text/html; charset:utf-8", []byte(imageHtml))
}

func main() {
	router := gin.Default()
	router.GET("/", MainHandler)
	router.Run()
}
