package lib

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var config Config = NewConfig()
var iPORT int16 = config.GetPort()
var PORT = fmt.Sprintf(":%s", strconv.Itoa(int(iPORT)))

func SetupApi() *gin.Engine {
	InitializeLogger()
	config.SetApiKeys()

	InfoLogger.Println("===PROGRAM_STARTED===")

	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})
	r.Static("/public/", "./public")
	r.LoadHTMLGlob("./templates/*.html")

	InfoLogger.Println("Routes are setting into application.")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	r.POST("/imageUpload", func(ctx *gin.Context) {
		var image_upload ImageUpload
		if err := ctx.ShouldBind(&image_upload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		InfoLogger.Println("Image is being uploaded.")
		InfoLogger.Println(image_upload)

		can_pass, assoc := FileNameAnalyzer(image_upload.Image.Filename)
		if can_pass {
			InfoLogger.Println("File name is valid.")
			ctx.SaveUploadedFile(image_upload.Image, "public/uploads/"+RandomIDGenerator()+assoc)
			ctx.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	r.GET("/coordinates/:latitude/:longitude", func(ctx *gin.Context) {
		latitude := ctx.Params.ByName("latitude")
		longitude := ctx.Params.ByName("longitude")

		response, status_code, err := GeoCodeQuery(latitude, longitude)

		LogError(err)
		formatted_address := gjson.Get(response, "features.0.properties.formatted").String()

		ctx.JSON(status_code, gin.H{
			"status": status_code,
			"result": formatted_address,
		})
	})

	r.GET("/dumpPolygon", func(ctx *gin.Context) {
		success := true
		all_polygons, err := FetchAllPolygons()

		if err != nil {
			log.Fatal(err)
			success = false
		}

		if success {
			ctx.JSON(http.StatusOK, gin.H{
				"all_locations": all_polygons,
				"successful":    true,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succesful": false,
		})
	})

	r.POST("/addPolygon", func(ctx *gin.Context) {
		success := true
		var request AddPolygon_RequestPayload
		if err := ctx.BindJSON(&request); err != nil {
			LogError(err)
			success = false
		}

		polygon_list := request.Polygon
		added_by := request.AddedBy
		added_time := request.AddedTime
		isInDanger := request.IsInDanger
		err := AddPolygon(polygon_list, string(added_by), string(added_time), isInDanger)

		if err != nil {
			fmt.Println(err.Error())
			success = false
		}

		if success {
			ctx.JSON(http.StatusOK, gin.H{
				"successful": true,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succesful": false,
		})

	})

	r.GET("/bbox/:latitude/:longitude", func(ctx *gin.Context) {
		latitude := ctx.Params.ByName("latitude")
		longitude := ctx.Params.ByName("longitude")

		response, status_code, err := GeoCodeQuery(latitude, longitude)
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

	// this route is deprecated
	r.GET("/_sendMail", func(ctx *gin.Context) {
		err := EmailSender("OAUTH")
		LogError(err)
	})

	return r
}
