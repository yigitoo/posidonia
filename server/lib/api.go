package lib

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var config Config = NewConfig()
var iPORT int16 = config.GetPort()
var PORT = fmt.Sprintf(":%s", strconv.Itoa(int(iPORT)))

func SetupApi() *gin.Engine {
	config.SetApiKeys()

	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})

	r.GET("/coordinates/:latitude/:longitude", func(ctx *gin.Context) {
		latitude := ctx.Params.ByName("latitude")
		longitude := ctx.Params.ByName("longitude")

		response, status_code, err := GeoCodeAPI(latitude, longitude)

		LogError(err)
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
		LogError(err)

		formatted_bbox_structs := gjson.Get(response, "features.0.bbox").Array()
		formatted_bbox := make([]float64, 4)
		for index, result := range formatted_bbox_structs {
			item, err := strconv.ParseFloat(result.Raw, 64)
			LogError(err)
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

	r.GET("/id/:userid", func(ctx *gin.Context) {
		user_id := ctx.Params.ByName("userid")
		user, err := GetUserByID(user_id)
		LogError(err)

		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"username": user.Username,
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var request User

		if err := ctx.BindJSON(&request); err != nil {
			panic(err)
		}

		user, err := ValidateLogin(request.Username, request.Password)
		LogError(err)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"user_id":  user.UserID,
				"username": user.Username,
			})
			// i use return because of
			return
		}
		// this context
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "We aren't able to login to you because of credentials.",
		})

	})

	return r
}

func GeoCodeAPI(latitude, longitude string) (string, int, error) {

	query_url := fmt.Sprintf(
		"https://api.geoapify.com/v1/geocode/reverse?lat=%s&lon=%s&apiKey=%s",
		latitude,
		longitude,
		config.GetApiKeys("geocode"),
	)

	response, err := http.Get(query_url)
	LogError(err)

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	LogError(err)

	return string(body), response.StatusCode, err
}
