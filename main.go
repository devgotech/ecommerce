package main

import (
	"log"
	"os"

	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/middleware"
	"ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	// router.GET("/", showIndexPage)
	// userRoutes := router.Group("/user")
	// {
	// 	userRoutes.GET("/signup", showRegistrationPage)
	// 	userRoutes.POST("/signup", controllers.SignUp())
	// }

	//other routes apart from user
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Println("serving:"+port)
	log.Fatal(router.Run(":" + port))
}
