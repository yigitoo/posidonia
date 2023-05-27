package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"github.com/yigitoo/posidonia/server/lib"
)

var config lib.Config = lib.NewConfig()
var iPORT int16 = config.GetPort()
var PORT = fmt.Sprintf(":%s", strconv.Itoa(int(iPORT)))

func SetupApi() *gin.Engine {
	config.SetApiKeys()

	r := gin.Default()

	r.GET("/coordinates/:latitude/:longitude", func(ctx *gin.Context) {
		latitude := ctx.Params.ByName("latitude")
		longitude := ctx.Params.ByName("longitude")

		query_url := fmt.Sprintf(
			"https://api.geoapify.com/v1/geocode/reverse?lat=%s&lon=%s&apiKey=%s",
			latitude,
			longitude,
			os.Getenv("API_KEY_GEOCODE"),
		)
		println(query_url)
		response, err := http.Get(query_url)
		logError(err)

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		logError(err)

		ctx.JSON(response.StatusCode, gin.H{
			"message": gojsonq.New().FromString(string(body)),
		})
	})

	return r
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}
