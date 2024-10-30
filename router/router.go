package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8080") // listen and serve on
}
