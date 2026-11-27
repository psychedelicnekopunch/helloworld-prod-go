
package main


import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	// 環境変数
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file\n")
	}
	key := "TEST"

	// ルーティング
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
			"env": fmt.Sprintf("ENV: [%s] %s", key, os.Getenv(key)),
		})
	})

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8080")
}

