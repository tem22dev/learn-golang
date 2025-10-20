package main

import (
	"learn-golang/internal/api/v1/handler"
	"learn-golang/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := handler.NewUserHandler()
			user.GET("/", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUsersByIdV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUsersByUuidV1)
			user.POST("/", userHandlerV1.PostUsersV1)
			user.PUT("/:id", userHandlerV1.PutUsersV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersV1)
		}

		product := v1.Group("/products")
		{
			productHandlerV1 := handler.NewProductHandler()
			product.GET("/", productHandlerV1.GetProductsV1)
			product.GET("/:slug", productHandlerV1.GetProductsBySlugV1)
			//product.GET("/:id", productHandlerV1.GetProductsByIdV1)
			product.POST("/", productHandlerV1.PostProductsV1)
			product.PUT("/:id", productHandlerV1.PutProductsV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductsV1)

		}

		categories := v1.Group("/categories")
		{
			categoryHandlerV1 := handler.NewCategoryHandler()
			categories.GET("/:category", categoryHandlerV1.GetProductsByCategoryV1)
			categories.POST("/", categoryHandlerV1.PostCategoryV1)

		}

		news := v1.Group("/news")
		{
			newsHandlerV1 := handler.NewNewsHandler()
			//news.GET("/", newsHandlerV1.GetNewsV1)
			//news.GET("/:slug", newsHandlerV1.GetNewsBySlugV1)
			news.POST("/", newsHandlerV1.PostNewsV1)
			news.POST("/upload-file", newsHandlerV1.PostUploadFileNewsV1)
			news.POST("/upload-multiple-file", newsHandlerV1.PostUploadMultipleFileNewsV1)

		}
	}

	router.StaticFS("/images", gin.Dir("./uploads", false))

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
