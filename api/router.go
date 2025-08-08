package api

import (
	"oms-test/internal/product"
	"oms-test/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userController := user.UserContorller{}
	productController := product.ProductController{}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/", userController.SearchUser)
		userRouter.POST("/", userController.CreateUser)
		userRouter.GET("/:id", userController.FetchUser)
		userRouter.PUT("/:id", userController.UpdateUser)
		userRouter.DELETE("/:id", userController.DeleteUser)
	}

	productRouter := router.Group("/product")
	{
		productRouter.GET("/", productController.SearchProduct)
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/:id", productController.ProductFromId)
		productRouter.PUT("/:id", productController.UpdateProduct)
		productRouter.DELETE("/:id", productController.DeleteProduct)
		productRouter.POST("/inflow/:id", productController.InflowProduct)
		productRouter.POST("/outflow/:id", productController.OutflowProduct)
	}

	return router
}
