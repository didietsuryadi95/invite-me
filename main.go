package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"github.com/didietsuryadi95/invite-me/articles"
	"github.com/didietsuryadi95/invite-me/common"
	"github.com/didietsuryadi95/invite-me/homes"
	"github.com/didietsuryadi95/invite-me/sites"
	"github.com/didietsuryadi95/invite-me/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
	sites.AutoMigrate()
}

func main() {
	config := common.LoadConfig()

	// db := common.Init(*config)
	// Migrate(db)
	// defer db.Close()

	// common.Setup(*config)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))
	homes.HomeResources(v1.Group("/homes"))
	sites.Site(v1.Group("/sites"))
	sites.AdminSite(v1.Group("/admin/sites"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
	sites.SiteSettings(v1.Group("/admin/settings"))

	articles.ArticlesRegister(v1.Group("/articles"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = config.Server.Port
	}

	r.Run(fmt.Sprintf("%s%s", ":", port)) // listen and serve on 0.0.0.0:8080
}
