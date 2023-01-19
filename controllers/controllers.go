package controllers

import(
	""
)

func HashPassword (password string) string{

}

func VerifyPassword (userPassword string, givenPassword string) (bool, string){

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

}

func searchProductByQuery() gin.HandlerFunc {

}