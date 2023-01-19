package controllers

import (
	"bytes"
	"context"
	"ecommerce/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword (password string) string{
	bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword (userPassword string, givenPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "Login or Password is incorrect"
		valid = false
	}
	return valid, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err !=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocument(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H("error": err))
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H("error": "user already exists"))
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})

		defer cancel()
		if err := nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H("error": err))
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H("error":"This phone no. is already in use"))
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshToken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token =&refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr := nil {
			c.JSON(http.StatusInternalServerError, gin.H("error":"the user did not get created"))
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, "Successfully signed in")
	}
}

func Login() gin.HandlerFunc {
	return fun(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.Userif err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H("error": err))
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": "login or password incorrect"})
		return
	}

	PasswordIsValid, msg := verifyPassword(*user.Password, *founduser.Password)

	defer cancel()

	if !PasswordIsValid [
		c.JSON(http.StatusInternalServerError, gin.H("error": msg))
		fmt.Println(msg)
	]

	token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, user.User_ID)
	defer cancel()

	generate.UpdateAllTokens(token, refreshToken, foun)
}

func ProductViewerAdmin() gin.HandlerFunc {
	
}

func searchProduct() gin.HandlerFunc {
	return func(c *gin.Context){

		var productList []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		ProductCollection.Find(ctx, bson.D{})
		if er != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong, please try after some time")
			return
		}

		err = cursor.All(ctx, &productList)

		if err != nill {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return
		}
		defer cursor.Close()

		if err := cursor.err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer cancel()
		c.IndentedJSON(200, productList)
	}
}

func searchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context){
		var searchProducts []models.Product
		queryParam := c.Query("name")

	// check if it's empty
		if queryParam == ""{
			log.Println("Query is emppty")
			c.Header("Content-Type", "application/json")
			c.JSON(HTTP.StatusNotFound, gin.H{"Error":"Invalid search index"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex":queryParam}})

		if err != nil {
			c.IndentedJSON(404, "something went wrong while fetching the data")
			return
		}

		err = searchquerydb.All(ctx, &searchProducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
		}

		defer searchquerydb.Close(ctx)

		if err := searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}

		defer cancel()
		c.IndentedJSON(200, searchProducts)
	}
}