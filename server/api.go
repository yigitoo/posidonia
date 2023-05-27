package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
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

		response, status_code, err := GeoCodeAPI(latitude, longitude)

		logError(err)
		formatted_address := gjson.Get(response, "features.0.properties.formatted").String()

		ctx.JSON(status_code, gin.H{
			"status": status_code,
			"result": formatted_address,
		})
	})

	r.GET("/bbox/:latitude/:longitude", func(ctx *gin.Context) {
		latitude := ctx.Params.ByName("latitude")
		longitude := ctx.Params.ByName("longitude")

		response, status_code, err := GeoCodeAPI(latitude, longitude)
		logError(err)

		formatted_bbox_structs := gjson.Get(response, "features.0.bbox").Array()
		formatted_bbox := make([]float64, 4)
		for index, result := range formatted_bbox_structs {
			item, err := strconv.ParseFloat(result.Raw, 64)
			logError(err)
			formatted_bbox[index] = item
		}

		ctx.JSON(status_code, gin.H{
			"status":    status_code,
			"bbox_list": formatted_bbox,
			"x_min":     formatted_bbox[0],
			"y_min":     formatted_bbox[1],
			"x_max":     formatted_bbox[2],
			"y_max":     formatted_bbox[3],
		})
	})

	return r
}

func GeoCodeAPI(latitude, longitude string) (string, int, error) {

	query_url := fmt.Sprintf(
		"https://api.geoapify.com/v1/geocode/reverse?lat=%s&lon=%s&apiKey=%s",
		latitude,
		longitude,
		os.Getenv("API_KEY_GEOCODE"),
	)

	response, err := http.Get(query_url)
	logError(err)

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	logError(err)

	return string(body), response.StatusCode, err
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
