package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID		primitive.ObjectID	`json:"_id" bson:"_id"`
	First_Name	*string
	Last_Name	*string
	Password	*string
	Email		*string
	Phone		*string
	Token		*string
	Refresh_Token	*string
	Created_At	time.Time
	Updated_At	time.Time
	User_ID		string
	UserCart	[]ProductUser
	Address_Details	[]Address
	OrderStatus	[]Order
}

type Product struct{
	Product_ID	primitive.ObjectID
	Product_Name	*string
	Price		*uint64
	Rating		*uint64		
	Image		*string
}
type ProductUser struct{
	Product_ID	primitive.ObjectID
	Product_Name	*string
	Price		*uint64
	Rating		*uint64
	Image		*string
}
type Address struct{
	Adress_ID	primitive.ObjectID
	House		*string
	Street		*string
	City		*string
	Zipcode		*string
}
type Order struct{
	Order_ID	primitive.ObjectID
	Order_Cart	[]ProductUser
	Ordered_At	time.Time
	Price		int
	Discount	*int
	Payment_method	Payment
}
type Payment struct{
	Digital bool
	COD	bool //cash on delivery
}
