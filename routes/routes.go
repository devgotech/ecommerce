package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes routes for the user
func UserRoutes(incomingRoutes *gin.Engine) {
	ir := incomingRoutes
	ir.POST("/user/signup", controllers.SignUp())
	ir.POST("/user/login", controllers.Login())
	ir.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	ir.GET("/user/productview", controllers.SearchProduct())
	ir.GET("/user/search", controllers.SearchProductByQuery())

	//FRONTEND ROUTES
	ir.LoadHTMLGlob("templates/*")
	ir.GET("/", )
}
