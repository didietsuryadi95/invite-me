package homes

import (
	"github.com/gin-gonic/gin"
)

func HomeResources(router *gin.RouterGroup) {
	router.GET("/", homePage)
}

func homePage(c *gin.Context) {

}
