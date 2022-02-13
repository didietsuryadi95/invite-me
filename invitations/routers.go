package invitations

import (
	"github.com/gin-gonic/gin"
)

func Invitation(router *gin.RouterGroup) {
	router.GET("/:slug", getEndUserInvitation)
	router.GET("/:slug/trackers", trackInvitation)
}

func AdminInvitation(router *gin.RouterGroup) {
	router.GET("/", getInvitationList)
	router.GET("/:id", getInvitationById)
	router.PUT("/:id", updateInvitationById)
	router.PUT("/:id/content", updateContentInvitationById)
	router.GET("/:id/tracker", getInvitationTrackerById)
	router.GET("/:id/modules", getModulesProperty)
	router.POST("/", createInvitationById)
	router.POST("/:id/order", createOrderInvitation)
}

func InvitationSettings(router *gin.RouterGroup) {
	router.GET("/templates/", getTemplateProperty)
	router.GET("/templates/:code/modules", getTemplateProperty)
	router.GET("/themes", getTemplateProperty)
	router.GET("/features", getTemplateProperty)
}

func getEndUserInvitation(c *gin.Context) {

}

func trackInvitation(c *gin.Context) {

}

func getInvitationTrackerById(c *gin.Context) {

}

func getInvitationList(c *gin.Context) {

}

func getInvitationById(c *gin.Context) {

}

func updateInvitationById(c *gin.Context) {

}

func updateContentInvitationById(c *gin.Context) {

}

func createInvitationById(c *gin.Context) {

}

func createOrderInvitation(c *gin.Context) {

}

func getModulesProperty(c *gin.Context) {

}

func getTemplateProperty(c *gin.Context) {

}