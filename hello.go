package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"github.com/gothinkster/golang-gin-realworld-example-app/articles"
	"github.com/gothinkster/golang-gin-realworld-example-app/common"
	"github.com/gothinkster/golang-gin-realworld-example-app/homes"
	"github.com/gothinkster/golang-gin-realworld-example-app/invitations"
	"github.com/gothinkster/golang-gin-realworld-example-app/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
	invitations.AutoMigrate()
}

func main() {
	config := common.LoadConfig()

	db := common.Init(*config)
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))
	homes.HomeResources(v1.Group("/homes"))
	invitations.Invitation(v1.Group("/invitations"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
	invitations.AdminInvitation(v1.Group("/admin/invitations"))
	invitations.InvitationSettings(v1.Group("/admin/settings"))

	articles.ArticlesRegister(v1.Group("/articles"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.Server.Port) // listen and serve on 0.0.0.0:8080
}
