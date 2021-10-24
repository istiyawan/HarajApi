package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HarajApi struct {
	Body         string `json:"body"`
	City         string `json:"city"`
	CommentCount int    `json:"commentCount"`
	Date         int    `json:"date"`
	ThumbURL     string `json:"thumbUrl"`
	Title        string `json:"title"`
	Username     string `json:"username"`
}

var harajAPI = []HarajApi{
	{Title: "ددسن 85 غمارتين", Username: "ffaris6", ThumbURL: "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/900x506_40794ac6-4d70-4b56-95cd-40be11922c44.jpg", CommentCount: 10, City: "الطايف", Date: 1621848879, Body: "السيارة: نيسان - ددسن\nالموديل: 1985\nبدون لوحات وبدون استمارة\nحالة السيارة: مستعملة\nالقير: قير عادي\nنوع الوقود: بنزين\nالسيارة لا تشكو من شي \nبدون لوحات واستماره \nتنفع للمزارع والحلال"},
	{Title: "85", Username: "ffaris6", ThumbURL: "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/900x506_40794ac6-4d70-4b56-95cd-40be11922c44.jpg", CommentCount: 10, City: "الطايف", Date: 1621848879, Body: "السيارة: نيسان - ددسن\nالموديل: 1985\nبدون لوحات وبدون استمارة\nحالة السيارة: مستعملة\nالقير: قير عادي\nنوع الوقود: بنزين\nالسيارة لا تشكو من شي \nبدون لوحات واستماره \nتنفع للمزارع والحلال"},
}

func getHarajCar(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, harajAPI)
}

func getHarajCarByTitle(c *gin.Context) {
	title := c.Param("title")
	fmt.Println("title", title)

	for _, car := range harajAPI {
		if car.Title == title {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Car Not Found"})

}

func main() {
	router := gin.Default()
	router.GET("/", getHarajCar)
	router.GET("/:title", getHarajCarByTitle)
	// router.Run("localhost:8080")
	router.Run(":" + os.Getenv("PORT"))
}
