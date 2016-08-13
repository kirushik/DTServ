package main

import "github.com/gin-gonic/gin"
import "net/http"

type handlers struct {
	greeter  gin.HandlerFunc
	goodbyer gin.HandlerFunc
}

func initHandlers(greeting string) handlers {
	return handlers{
		greeter: func(c *gin.Context) {
			name := c.DefaultQuery("name", "незнакомец")
			c.String(http.StatusOK, "%s, %s", greeting, name)
		}, goodbyer: func(c *gin.Context) {
			c.String(http.StatusOK, "Bye-bye")
		}}
}

func main() {
	router := gin.Default()
	handlers := initHandlers("Алоха")
	router.GET("/hello", handlers.greeter)
	router.GET("/goodbye", handlers.goodbyer)
	router.Run(":8080")
}
