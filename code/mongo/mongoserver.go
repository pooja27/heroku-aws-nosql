package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

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

// DeleteUserCart deletes the user cart
func (rc ResponseController) DeleteUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userid := p.ByName("userid")
	fmt.Println("DELETE Request Delete user Cart: userid:", userid)

	if err := rc.session.DB(Database).C("cart").Remove(bson.M{"userid" : userid}); err != nil {
		sendErrorResponse(w,err,500)
		return
	}

	sendSuccessResponseCartDeleted(w, userid)

}

// RemoveFromUserCart serves the user cart remove item PUT request
func (rc ResponseController) RemoveFromUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 itemid := p.ByName("itemid")
 fmt.Println("PUT Request Remove From Cart: userid:", userid, " item_id:", itemid)

 var newList []Item
 currCart, err := getUserCartDB(userid, rc)
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
func (rc ResponseController) AddToUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 itemid := p.ByName("itemid")
 typee := p.ByName("type")
 fmt.Println("PUT Request Add To Cart: userid:", userid, " item_id:", itemid, " type:", typee)

 var price float64
 var name string
 errMsg := "item id not found for type " + typee

 if typee == "tea" {
	 product, err := getProductTeaDB(itemid, rc)
	 if err != nil {
		 sendErrorResponse(w,errors.New(errMsg),404)
		 return
	 } else {
		 price = product.Price
		 name = product.Name
	 }
 } else if typee == "coffee" {
	 product, err := getProductCoffeeDB(itemid, rc)
	 if err != nil {
		 sendErrorResponse(w,errors.New(errMsg),404)
		 return
	 } else {
		 price = product.Price
		 name = product.Name
	 }
 } else if typee == "drinkware" {
	 product, err := getProductDrinkwareDB(itemid, rc)
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
 currCart, err := getUserCartDB(userid, rc)
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
func (rc ResponseController) GetUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 fmt.Println("GET Request Cart: userid:", userid)

 resp, err := getUserCartDB(userid, rc)
 if err != nil {
	 sendErrorResponse(w,err,404)
	 return
 }

 jsonOut, _ := json.Marshal(resp)
 httpResponse(w, jsonOut, 200)
 fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// UpdateInventoryRemoveOne serves the Inventory PUT request to decrement by one
func (rc ResponseController) UpdateInventoryRemoveOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	itemid := p.ByName("itemid")
	fmt.Println("PUT Request: inventoryRemoveOne : itemid:", itemid)

	var resp Inventory

	resp, err := getInventoryDB(itemid, rc)
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
func (rc ResponseController) UpdateInventoryAddOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	itemid := p.ByName("itemid")
	fmt.Println("PUT Request: inventoryAddOne : itemid:", itemid)

	var resp Inventory

	resp, err := getInventoryDB(itemid, rc)
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
func (rc ResponseController) GetInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	itemid := p.ByName("itemid")
	fmt.Println("GET Request inventory: itemid:", itemid)

	resp, err := getInventoryDB(itemid, rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetAllCoffee serves the coffee GET request
func (rc ResponseController) GetAllCoffee(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request all coffee")

	resp, err := getAllCoffeeDB(rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetAllTea serves the tea GET request
func (rc ResponseController) GetAllTea(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request all tea")

	resp, err := getAllTeaDB(rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// GetAllTea serves the tea GET request
func (rc ResponseController) GetAllDrinkware(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request all drinkware")

	resp, err := getAllDrinkwareDB(rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

 // GetDrinkware serves the drinkware GET request
 func (rc ResponseController) GetDrinkware(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 	productid := p.ByName("itemid")
 	fmt.Println("GET Request Drinkware: itemid:", productid)

 	resp, err := getProductDrinkwareDB(productid, rc)
 	if err != nil {
 		sendErrorResponse(w,err,404)
 		return
 	}

 	jsonOut, _ := json.Marshal(resp)
 	httpResponse(w, jsonOut, 200)
 	fmt.Println("Response:", string(jsonOut), " 200 OK")
 }

 // GetCoffee serves the coffee GET request
 func (rc ResponseController) GetCoffee(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 	productid := p.ByName("itemid")
 	fmt.Println("GET Request Coffee: itemid:", productid)

 	resp, err := getProductCoffeeDB(productid, rc)
 	if err != nil {
 		sendErrorResponse(w,err,404)
 		return
 	}

 	jsonOut, _ := json.Marshal(resp)
 	httpResponse(w, jsonOut, 200)
 	fmt.Println("Response:", string(jsonOut), " 200 OK")
 }

// GetTea serves the tea GET request
func (rc ResponseController) GetTea(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productid := p.ByName("itemid")
	fmt.Println("GET Request Tea: itemid:", productid)

	resp, err := getProductTeaDB(productid, rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// Signup serves the signup POST request
func (rc ResponseController) Signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req UserDetails

	defer r.Body.Close()

	jsonIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w,err,400)
		return
	}

	json.Unmarshal([]byte(jsonIn), &req)
	fmt.Println("POST Request signup:", req)

	if err := rc.session.DB("db_test").C("users").Insert(req); err != nil {
		sendErrorResponse(w,err,500)
		return
	}
	sendSuccessResponse(w,201)
}

// Login serves the Login GET request
func (rc ResponseController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userid := p.ByName("userid")
	fmt.Println("GET Request login: userid:", userid)

	resp, err := getUserDetailsDB(userid, rc)
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
func getUserCartDB(userid string, rc ResponseController) (Cart, error) {
	var resp Cart

	if err := rc.session.DB(Database).C("cart").Find(bson.M{"userid" : userid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get inventory corresponding to the item id
func getInventoryDB(itemid string, rc ResponseController) (Inventory, error) {
	var resp Inventory

	if err := rc.session.DB(Database).C("inventory").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get all tea products from DB
func getAllDrinkwareDB(rc ResponseController) ([]ProductDrinkware, error) {
	var resp []ProductDrinkware

	if err := rc.session.DB(Database).C("drinkware").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get all tea products from DB
func getAllTeaDB(rc ResponseController) ([]ProductTea, error) {
	var resp []ProductTea

	if err := rc.session.DB(Database).C("tea").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get all coffee products from DB
func getAllCoffeeDB(rc ResponseController) ([]ProductCoffee, error) {
	var resp []ProductCoffee

	if err := rc.session.DB(Database).C("coffee").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get Drinkware data corresponding to the productid
func getProductDrinkwareDB(itemid string, rc ResponseController) (ProductDrinkware, error) {
	var resp ProductDrinkware

	if err := rc.session.DB(Database).C("drinkware").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get Coffee data corresponding to the productid
func getProductCoffeeDB(itemid string, rc ResponseController) (ProductCoffee, error) {
	var resp ProductCoffee

	if err := rc.session.DB(Database).C("coffee").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get Tea data corresponding to the productid
func getProductTeaDB(itemid string, rc ResponseController) (ProductTea, error) {
	var resp ProductTea

	if err := rc.session.DB(Database).C("tea").Find(bson.M{"item_id" : itemid}).One(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get data corresponding to the user id
func getUserDetailsDB(userid string, rc ResponseController) (UserDetails, error) {
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

//NewResponseController function returns reference to ResponseController and a mongoDB session
func NewResponseController(s *mgo.Session) *ResponseController {
	return &ResponseController{s}
}

//returns a mongoDB session
func getSession() *mgo.Session {
    //Enter mongoLab connection string here
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Panic@getSession.Dial")
		panic(err)
	}
	return s
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

  fmt.Println("Starting server...")
	r := httprouter.New()
	rc := NewResponseController(getSession())

	r.GET("/mongoserver/login/:userid", rc.Login)

	r.GET("/mongoserver/teas", rc.GetAllTea)
	r.GET("/mongoserver/coffees", rc.GetAllCoffee)
	r.GET("/mongoserver/drinkwares", rc.GetAllDrinkware)

	r.GET("/mongoserver/tea/:itemid", rc.GetTea)
	r.GET("/mongoserver/coffee/:itemid", rc.GetCoffee)
	r.GET("/mongoserver/drinkware/:itemid", rc.GetDrinkware)

	r.GET("/mongoserver/inventory/:itemid", rc.GetInventory)
	r.PUT("/mongoserver/inventory/addOne/:itemid", rc.UpdateInventoryAddOne)
	r.PUT("/mongoserver/inventory/removeOne/:itemid", rc.UpdateInventoryRemoveOne)

	r.PUT("/mongoserver/cart/addItem/:userid/:type/:itemid", rc.AddToUserCart)
	r.PUT("/mongoserver/cart/removeItem/:userid/:itemid", rc.RemoveFromUserCart)
	r.GET("/mongoserver/cart/:userid", rc.GetUserCart)
	r.DELETE("/mongoserver/cart/:userid", rc.DeleteUserCart)

	r.POST("/mongoserver/signup", rc.Signup)

	fmt.Println("Server is Ready !")
	fmt.Println(http.ListenAndServe(":7777",r))

}
