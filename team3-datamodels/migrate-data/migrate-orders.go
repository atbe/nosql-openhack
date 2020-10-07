package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"openhack/cart"
	"openhack/cartitem"
	"openhack/item"
	openhackUser "openhack/user"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var server = "openhacksqlsuf1svlflnun7rura.database.windows.net"
var port = 1433
var username = "openhackadmin"
var password = "Password123"
var database = "Movies"

var db *gorm.DB

var mongoDBConnectionString = "AccountEndpoint=https://cosmosdb2-team3.documents.azure.com:443/;AccountKey=uDw5eCXYABEpiz16QTKvd07GhCt68D9g9pNye6shFbalntrmHS1gI9NXCY0m8TOHjGYwtqjY7X2j19mq3Wt7hQ==;"

func main() {
	// Uncomment for debug output
	log.SetFlags(0)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, username, password, port, database)
	if conDb, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db = conDb
	}

	items := make([]cartitem.CartItem, 0)
	if err := db.Model(&cartitem.CartItem{}).Find(&items).Error; err != nil {
		panic(err)
	}
	log.Print(fmt.Sprintf("There are %v items to migrate", len(items)))

	users := make([]openhackUser.User, 0)
	if err := db.Model(&openhackUser.User{}).Find(&users).Error; err != nil {
		panic(err)
	}
	log.Print(fmt.Sprintf("There are %v users to migrate", len(users)))

	carts := make([]cart.CosmosCart, 0)
	for _, user := range users {
		cart := migrateUser(user)
		carts = append(carts, cart)
		// break
	}

	cartsStr, _ := json.MarshalIndent(carts, "", "    ")
	fmt.Printf("%v\n", string(cartsStr))

	// write the whole body at once
	err := ioutil.WriteFile("carts.json", cartsStr, 0644)
	if err != nil {
		panic(err)
	}
}

func migrateUser(user openhackUser.User) cart.CosmosCart {
	log.Print(fmt.Sprintf("Migrating user with email %v and id %v\n\n\n", user.Email, user.ID))

	cartItems := getCosmosOrderItems(user)
	log.Print(len(cartItems))

	cosmosCart := cart.CosmosCart{
		UserId: user.ID,
		Items:  cartItems,
	}
	return cosmosCart
}

func getCosmosOrderItems(user openhackUser.User) []cartitem.CosmosCartItem {
	usersCart := cart.Cart{}
	if err := db.Model(&cart.Cart{}).Where("UserId = ?", user.ID).First(&usersCart).Error; err != nil {
		panic(err)
	}
	log.Print(fmt.Sprintf("Users cart has id %v", string(usersCart.ID)))

	userCartItems := make([]cartitem.CartItem, 0)
	if err := db.Model(&cartitem.CartItem{}).Where("CartId = ?", usersCart.ID).Find(&userCartItems).Error; err != nil {
		panic(err)
	}
	log.Print(fmt.Sprintf("Found %v number of user items in the cart", len(userCartItems)))

	cosmosCartItems := make([]cartitem.CosmosCartItem, 0)
	for _, cartItem := range userCartItems {
		theItem := item.Item{}
		if err := db.Model(&item.Item{}).Preload("Category").Where("ItemId = ?", cartItem.ItemId).First(&theItem).Error; err != nil {
			panic(err)
		}

		log.Print(fmt.Sprintf("The cart item with id %v maps to the item with name %v and price %v", cartItem.CartItemId, theItem.Name, theItem.UnitPrice))

		log.Print(theItem.Category)

		cosmosCartItem := cartitem.CosmosCartItem{
			Quantity:  cartItem.Quantity,
			UnitPrice: theItem.UnitPrice,
			Item: item.CosmosItem{
				Name:                theItem.Name,
				CategoryDescription: theItem.Category.Description,
				CategoryName:        theItem.Category.Name,
				CategoryID:          theItem.Category.ID,
			},
		}

		cosmosCartItems = append(cosmosCartItems, cosmosCartItem)
	}

	return cosmosCartItems
}
