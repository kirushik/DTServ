package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

type handlers struct {
	greeter  gin.HandlerFunc
	goodbyer gin.HandlerFunc
}

func initHandlers(greeting string) handlers {
	return handlers{
		greeter: func(c *gin.Context) {
			name := c.DefaultQuery("name", "незнакомец")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%s, %s", greeting, name),
			})
		}, goodbyer: func(c *gin.Context) {
			c.String(http.StatusOK, "Bye-bye")
		}}
}

func GetMainEngine(greeting string) *gin.Engine {
	router := gin.Default()
	handlers := initHandlers(greeting)
	router.GET("/hello", handlers.greeter)
	return router
}

func main() {
	router := GetMainEngine("Алоха")
	router.Run(":8080")
}
