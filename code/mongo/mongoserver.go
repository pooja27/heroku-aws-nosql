package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
    Database = "starbucks"
)

/* ------- Debugging variables ------- */
var out io.Writer
var debugModeActivated bool

/* ------- Global Variables ------- */
var rc *ResponseController
var mongoTimeout time.Duration

/* ------- Structs ------- */

//struct to store in Inventory collection
type Inventory struct {
	Item_ID string	`json:"item_id"	bson:"item_id"`
	Stock		int64		`json:"stock"		bson:"stock"`
}

//struct to store in Cart collection
type Cart struct {
	Userid		string 		`json:"userid" 	bson:"userid"`
	Items 		[]Item	`json:"items"	bson:"items"`
}

type Item struct {
	Item_ID		string 		`json:"item_id" 	bson:"item_id"`
	Price 		float64		`json:"price"			bson:"price"`
	Name 			string		`json:"name" 			bson:"name"`
}

//struct to store in coffee collection
type ProductCoffee struct {
	Region		string	`json:"region"		bson:"region"`
	Category	string 	`json:"category"	bson:"category"`
	Price			float64 `json:"price" 		bson:"price"`
	Item_ID 	string 	`json:"item_id" 	bson:"item_id"`
	Name 			string	`json:"name" 			bson:"name"`
	Flavor		string	`json:"flavor" 		bson:"flavor"`
	Quantity	int64		`json:"quantity" 	bson:"quantity"`
	Roast			string	`json:"roast" 		bson:"roast"`
	Type			string	`json:"type" 			bson:"type"`
}

//struct to store in tea collection
type ProductTea struct {
	Category	string 	`json:"category"	bson:"category"`
	Price			float64 `json:"price" 		bson:"price"`
	Count			int64		`json:"count" 		bson:"count"`
	Item_ID 	string 	`json:"item_id" 	bson:"item_id"`
	Name 			string	`json:"name" 			bson:"name"`
	Brand			string	`json:"brand" 		bson:"brand"`
	Type			string	`json:"type" 			bson:"type"`
	TeaForm		string	`json:"tea_form" 	bson:"tea_form"`
}

//struct to store in drinkware collection
type ProductDrinkware struct {
	Category	string 	`json:"category" 	bson:"category"`
	Price			float64 `json:"price" 		bson:"price"`
	Item_ID 	string 	`json:"item_id" 	bson:"item_id"`
	Name 			string	`json:"name" 			bson:"name"`
}

//struct to return success / failure status
type ResponseStatus struct {
	Status 	string `json:"status" 	bson:"status"`
	Details string `json:"details" 	bson:"details"`
}

//to store in users collection
type UserDetails struct {
	Userid		string `json:"userid" 	bson:"userid"`
	Password	string `json:"password" bson:"password"`
	Email 		string `json:"email" 		bson:"email"`
	Name 			string `json:"name" 		bson:"name"`
}

//ResponseController struct to provide to httprouter
type ResponseController struct {
	session *mgo.Session
}

/* ------- REST Functions ------- */

// HealthCheck serves the healthcheck GET request
func HealthCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("HealthCheck")
	if err := checkMongoSession(w); err !=nil {
		return
	}
	sendSuccessResponse(w,200)
}

// DeleteUserCart deletes the user cart
func DeleteUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userid := p.ByName("userid")
	fmt.Println("DELETE Request Delete user Cart: userid:", userid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	if err := rc.session.DB(Database).C("cart").Remove(bson.M{"userid" : userid}); err != nil {
		sendErrorResponse(w,err,500)
		return
	}

	sendSuccessResponseCartDeleted(w, userid)

}

// RemoveFromUserCart serves the user cart remove item PUT request
func RemoveFromUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 itemid := p.ByName("itemid")
 fmt.Println("PUT Request Remove From Cart: userid:", userid, " item_id:", itemid)

 if err := checkMongoSession(w); err !=nil {
	 return
 }

 var newList []Item
 currCart, err := getUserCartDB(userid)
 if err == nil {
	 size := len(currCart.Items)
	 if size < 2 {
		 if currCart.Items[0].Item_ID == itemid {
			 if err := rc.session.DB(Database).C("cart").Remove(bson.M{"userid" : userid}); err != nil {
				 sendErrorResponse(w,err,500)
			 } else {
				 sendSuccessResponseCartDeleted(w, userid)
			 }
			 return
		 } else {
			 sendErrorResponse(w,errors.New("item_id invalid"),404)
			 return
		 }
	 }
	 newList = make([]Item, size-1)
	 i := 0
	 j := 0
	 found := false
	 for ; i<size && j<size-1; i++ { //keep adding items until not found, if found add all remaining items
		 if found || currCart.Items[i].Item_ID != itemid {
			 newList[j] = currCart.Items[i]
			 j++
		 } else {
			 found = true
		 }
	}

	if found == false && currCart.Items[i].Item_ID != itemid { //if item is not found and last item is also not equal to given item
		sendErrorResponse(w,errors.New("item_id invalid"),404)
		return
	}

	currCart.Items = newList
	change := bson.M{"$set": bson.M{"items" : currCart.Items}}

  if err := rc.session.DB(Database).C("cart").Update(bson.M{"userid": userid}, change); err != nil {
 	 sendErrorResponse(w,err,500)
 	 return
  }
	jsonOut, _ := json.Marshal(currCart)
  httpResponse(w, jsonOut, 200)
  fmt.Println("Response:", string(jsonOut), " 200 OK")
	return
 } else {
	 sendErrorResponse(w,err,404)
 }

}

// AddToUserCart serves the user cart add item PUT request
func AddToUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 itemid := p.ByName("itemid")
 typee := p.ByName("type")
 fmt.Println("PUT Request Add To Cart: userid:", userid, " item_id:", itemid, " type:", typee)

 if err := checkMongoSession(w); err !=nil {
	 return
 }

 var price float64
 var name string
 errMsg := "item id not found for type " + typee

 if typee == "tea" {
	 product, err := getProductTeaDB(itemid)
	 if err != nil {
		 sendErrorResponse(w,errors.New(errMsg),404)
		 return
	 } else {
		 price = product.Price
		 name = product.Name
	 }
 } else if typee == "coffee" {
	 product, err := getProductCoffeeDB(itemid)
	 if err != nil {
		 sendErrorResponse(w,errors.New(errMsg),404)
		 return
	 } else {
		 price = product.Price
		 name = product.Name
	 }
 } else if typee == "drinkware" {
	 product, err := getProductDrinkwareDB(itemid)
	 if err != nil {
		 sendErrorResponse(w,errors.New(errMsg),404)
		 return
	 } else {
		 price = product.Price
		 name = product.Name
	 }
 } else {
	 sendErrorResponse(w,errors.New("item type is invalid"),404)
	 return
 }

 var newList []Item
 currCart, err := getUserCartDB(userid)
 size := 1
 if err == nil {
	 size = len(currCart.Items) + 1
	 newList = make([]Item, size)
	 for i, x := range currCart.Items {
		newList[i] = x
	}
 } else {
	 currCart.Userid = userid
	 newList = make([]Item, 1)
 }
 newList[size-1].Item_ID = itemid
 newList[size-1].Price = price
 newList[size-1].Name = name

 currCart.Items = newList

 change := bson.M{"$set": bson.M{"items" : currCart.Items}}

if _,err := rc.session.DB(Database).C("cart").Upsert(bson.M{"userid": userid}, change); err != nil {
	 sendErrorResponse(w,err,500)
	 return
 }

 jsonOut, _ := json.Marshal(currCart)
 httpResponse(w, jsonOut, 200)
 fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetUserCart serves the user cart GET request
func GetUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 fmt.Println("GET Request Cart: userid:", userid)

 if err := checkMongoSession(w); err !=nil {
	 return
 }

 resp, err := getUserCartDB(userid)
 if err != nil {
	 sendErrorResponse(w,err,404)
	 return
 }

 jsonOut, _ := json.Marshal(resp)
 httpResponse(w, jsonOut, 200)
 fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// UpdateInventoryRemoveOne serves the Inventory PUT request to decrement by one
func UpdateInventoryRemoveOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	itemid := p.ByName("itemid")
	fmt.Println("PUT Request: inventoryRemoveOne : itemid:", itemid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	var resp Inventory

	resp, err := getInventoryDB(itemid)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	if resp.Stock < 1 {
		sendOutOfStockResponse(w, itemid)
		return
	}

	resp.Stock--

	change := bson.M{"$set": bson.M{"stock" : resp.Stock}}

	if err := rc.session.DB(Database).C("inventory").Update(bson.M{"item_id": itemid}, change); err != nil {
		sendErrorResponse(w,err,500)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// UpdateInventoryAddOne serves the Inventory PUT request to increment by one
func UpdateInventoryAddOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	itemid := p.ByName("itemid")
	fmt.Println("PUT Request: inventoryAddOne : itemid:", itemid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	var resp Inventory

	resp, err := getInventoryDB(itemid)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	resp.Stock++

	change := bson.M{"$set": bson.M{"stock" : resp.Stock}}

	if err := rc.session.DB(Database).C("inventory").Update(bson.M{"item_id": itemid}, change); err != nil {
		sendErrorResponse(w,err,500)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetInventory serves the Inventory GET request
func GetInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	itemid := p.ByName("itemid")
	fmt.Println("GET Request inventory: itemid:", itemid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	resp, err := getInventoryDB(itemid)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetAllCoffee serves the coffee GET request
func GetAllCoffee(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request all coffee")

	if err := checkMongoSession(w); err !=nil {
		return
	}

	resp, err := getAllCoffeeDB()
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetAllTea serves the tea GET request
func GetAllTea(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request all tea")

	if err := checkMongoSession(w); err !=nil {
		return
	}

	resp, err := getAllTeaDB()
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetAllTea serves the tea GET request
func GetAllDrinkware(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request all drinkware")

	if err := checkMongoSession(w); err !=nil {
		return
	}

	resp, err := getAllDrinkwareDB()
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

 // GetDrinkware serves the drinkware GET request
 func GetDrinkware(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 	productid := p.ByName("itemid")
 	fmt.Println("GET Request Drinkware: itemid:", productid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

 	resp, err := getProductDrinkwareDB(productid)
 	if err != nil {
 		sendErrorResponse(w,err,404)
 		return
 	}

 	jsonOut, _ := json.Marshal(resp)
 	httpResponse(w, jsonOut, 200)
 	fmt.Println("Response:", string(jsonOut), " 200 OK")
 }

 // GetCoffee serves the coffee GET request
 func GetCoffee(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 	productid := p.ByName("itemid")
 	fmt.Println("GET Request Coffee: itemid:", productid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

 	resp, err := getProductCoffeeDB(productid)
 	if err != nil {
 		sendErrorResponse(w,err,404)
 		return
 	}

 	jsonOut, _ := json.Marshal(resp)
 	httpResponse(w, jsonOut, 200)
 	fmt.Println("Response:", string(jsonOut), " 200 OK")
 }

// GetTea serves the tea GET request
func GetTea(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productid := p.ByName("itemid")
	fmt.Println("GET Request Tea: itemid:", productid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	resp, err := getProductTeaDB(productid)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// Signup serves the signup POST request
func Signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var req UserDetails

	defer r.Body.Close()

	jsonIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w,err,400)
		return
	}

	json.Unmarshal([]byte(jsonIn), &req)
	fmt.Println("POST Request signup:", req)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	if err := rc.session.DB("db_test").C("users").Insert(req); err != nil {
		sendErrorResponse(w,err,500)
		return
	}
	sendSuccessResponse(w,201)
}

// Login serves the Login GET request
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userid := p.ByName("userid")
	fmt.Println("GET Request login: userid:", userid)

	if err := checkMongoSession(w); err !=nil {
		return
	}

	resp, err := getUserDetailsDB(userid)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

/* ------- Helper Functions ------- */

//Get cart corresponding to the userid
func getUserCartDB(userid string) (Cart, error) {
	var resp Cart

	if err := rc.session.DB(Database).C("cart").Find(bson.M{"userid" : userid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get inventory corresponding to the item id
func getInventoryDB(itemid string) (Inventory, error) {
	var resp Inventory

	if err := rc.session.DB(Database).C("inventory").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get all tea products from DB
func getAllDrinkwareDB() ([]ProductDrinkware, error) {
	var resp []ProductDrinkware

	if err := rc.session.DB(Database).C("drinkware").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get all tea products from DB
func getAllTeaDB() ([]ProductTea, error) {
	var resp []ProductTea

	if err := rc.session.DB(Database).C("tea").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get all coffee products from DB
func getAllCoffeeDB() ([]ProductCoffee, error) {
	var resp []ProductCoffee

	if err := rc.session.DB(Database).C("coffee").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get Drinkware data corresponding to the productid
func getProductDrinkwareDB(itemid string) (ProductDrinkware, error) {
	var resp ProductDrinkware

	if err := rc.session.DB(Database).C("drinkware").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get Coffee data corresponding to the productid
func getProductCoffeeDB(itemid string) (ProductCoffee, error) {
	var resp ProductCoffee

	if err := rc.session.DB(Database).C("coffee").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get Tea data corresponding to the productid
func getProductTeaDB(itemid string) (ProductTea, error) {
	var resp ProductTea

	if err := rc.session.DB(Database).C("tea").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get data corresponding to the user id
func getUserDetailsDB(userid string) (UserDetails, error) {
	var resp UserDetails

	if err := rc.session.DB("db_test").C("users").Find(bson.M{"userid" : userid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func sendOutOfStockResponse(w http.ResponseWriter, itemid string) {
	var resp ResponseStatus
	resp.Status = "failure"
	resp.Details = "Item " + itemid + " stock is 0. Cannot process."
	fmt.Println(resp)
	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 412)
	fmt.Println("Response:", string(jsonOut), 412)
}

func sendErrorResponse(w http.ResponseWriter, err error, httpCode int) {
	var resp ResponseStatus
	resp.Status = "failure"
	resp.Details = err.Error()
	fmt.Println(err)
	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, httpCode)
	fmt.Println("Response:", string(jsonOut), httpCode)
}

func sendSuccessResponse(w http.ResponseWriter, httpCode int) {
	var resp ResponseStatus
	resp.Status = "success"
	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, httpCode)
	fmt.Println("Response:", string(jsonOut), httpCode)
}

func sendSuccessResponseCartDeleted(w http.ResponseWriter, userid string) {
	var resp ResponseStatus
	resp.Status = "success"
	resp.Details = "cart for user " + userid + " deleted"
	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), 200)
}

//write http response
func httpResponse(w http.ResponseWriter, jsonOut []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", jsonOut)
}

//gets a new mongo session, return error if not able to get mongo session
// this function has been added to handle mongodb sessions stopping in background
func checkMongoSession(w http.ResponseWriter) error {
	sess, err := getSession()
	rc = NewResponseController(sess)
	if (err != nil) {
		sendErrorResponse(w,err,500)
		return err
	}
	return nil
}

//NewResponseController function returns reference to ResponseController and a mongoDB session
func NewResponseController(s *mgo.Session) *ResponseController {
	return &ResponseController{s}
}

//returns a mongoDB session
func getSession() (session *mgo.Session, err error) {
  //Enter mongoLab connection string here
	s, err := mgo.DialWithTimeout("mongodb://localhost", mongoTimeout)
	if err != nil {
		fmt.Println("Unable to get Mongo session")
		return nil, errors.New("Unable to get Mongo session")
	}
	return s,nil
}

/* ------- Main Function ------- */

func main() {
	//debugging variables----------------------
	debugModeActivated = true
	out = ioutil.Discard
	if debugModeActivated {
		out = os.Stdout
	}
	//---------------------debugging variables

	mongoTimeout = time.Duration(2 * time.Second)

  fmt.Println("Starting server...")
	r := httprouter.New()

	sess, _ := getSession()
	rc = NewResponseController(sess)

	r.GET("/mongoserver/healthcheck", HealthCheck)

	r.GET("/mongoserver/login/:userid", Login)

	r.GET("/mongoserver/teas", GetAllTea)
	r.GET("/mongoserver/coffees", GetAllCoffee)
	r.GET("/mongoserver/drinkwares", GetAllDrinkware)

	r.GET("/mongoserver/tea/:itemid", GetTea)
	r.GET("/mongoserver/coffee/:itemid", GetCoffee)
	r.GET("/mongoserver/drinkware/:itemid", GetDrinkware)

	r.GET("/mongoserver/inventory/:itemid", GetInventory)
	r.PUT("/mongoserver/inventory/addOne/:itemid", UpdateInventoryAddOne)
	r.PUT("/mongoserver/inventory/removeOne/:itemid", UpdateInventoryRemoveOne)

	r.PUT("/mongoserver/cart/addItem/:userid/:type/:itemid", AddToUserCart)
	r.PUT("/mongoserver/cart/removeItem/:userid/:itemid", RemoveFromUserCart)
	r.GET("/mongoserver/cart/:userid", GetUserCart)
	r.DELETE("/mongoserver/cart/:userid", DeleteUserCart)

	r.POST("/mongoserver/signup", Signup)

	fmt.Println("Server is Ready !")
	fmt.Println(http.ListenAndServe(":7777",r))

}
