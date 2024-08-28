package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-myself/kakao"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	r := gin.Default()

	// CORS 설정
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"status": "success", "message": "잘해써~"})
	})

	r.POST("/submit", func(c *gin.Context) {
		var jsonData struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required,email"`
		}

		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invaild JSON data",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Data received successfully",
			"data": gin.H{
				"name":  jsonData.Name,
				"email": jsonData.Email,
			},
		})
	})

	r.GET("/kakaoLogin", func(c *gin.Context) {

		kakao.DoLogin()

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "오케이 로그인",
		})
	})
	//
	//db.InitMongoDB("mongodb+srv://test:1111@cluster0.rhxlati.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	//db.CreateCollection("stock")
	//db.InsertOne("stock")
	r.Run(":5174")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
