package api

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	userApi := route.Group("/user")
	{
		userApi.GET("/CreatUser")
	}
}
