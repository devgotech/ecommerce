package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gotech/ecommerce/controllers"
)

// routes for the user
func UserRoutes(incomingRoutes *gin.Engine) {
	ir := incomingRoutes
	ir.POST("/user/signup", controllers.SignUp())
	ir.POST("/user/login", controllers.Login())
	ir.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	ir.GET("/user/productview", controllers.SearchProduct())
	ir.GET("/user/search", controllers.SearchProductByQuery())
}
