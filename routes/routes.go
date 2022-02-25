package routes

import (
	"go-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandlerRequest() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/articles", controllers.GetAllArticles )
	router.GET("/articles/:limit/:offset", controllers.PaginationPostedArticle )
	router.GET("/article/:id", controllers.GetArticleById )
	router.POST("/article", controllers.CreateNewArticle )	
	router.PUT("/article/:id", controllers.UpdateArticle )
	router.DELETE("/article/:id", controllers.DeleteArticleById )

	router.Run(":8080")
}