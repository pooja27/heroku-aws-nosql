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

//struct to store in products collection
// type Product struct {
// 	ProductID		string	`json:"productid" bson:"productid"`
// 	Name				string 	`json:"name" bson:"name"`
// 	Type				string 	`json:"type" bson:"type"`
// 	Description string 	`json:"description" bson:"description"`
// 	Price 			float64	`json:"price" bson:"price"`
// }

//struct to store in Inventory collection
type Inventory struct {
	Item_ID string	`json:"item_id"	bson:"item_id"`
	Stock		int64		`json:"stock"		bson:"stock"`
}

//struct to store in Inventory collection
type Cart struct {
	Userid		string 		`json:"userid" 	bson:"userid"`
	Item_IDs 	[]string	`json:"item_ids"	bson:"item_ids"`
}

//struct to receive Inventory update request
// type Stock struct {
// 	Stock		int64		`json:"stock"		bson:"stock"`
// }

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

// RemoveFromUserCart serves the user cart remove item PUT request
func (rc ResponseController) RemoveFromUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 itemid := p.ByName("itemid")
 fmt.Println("PUT Request Remove From Cart: userid:", userid, " item_id:", itemid)

 var newList []string
 currCart, err := getUserCartDB(userid, rc)
 if err == nil {
	 size := len(currCart.Item_IDs)
	 if size < 2 {
		 if err := rc.session.DB(Database).C("cart").Remove(bson.M{"userid" : userid}); err != nil {
			 sendErrorResponse(w,err,500)
		 } else {
			 sendSuccessResponseCartDeleted(w, userid)
		 }
		 return
	 }
	 newList = make([]string, size-1)
	 i := 0
	 j := 0
	 found := false
	 for i := 0; i<size && j<size-1; i++ { //keep adding items until not found, if found add all remaining items
		 if found || currCart.Item_IDs[i] != itemid {
			 newList[j] = currCart.Item_IDs[i]
			 j++
		 } else {
			 found = true
		 }
	}

	if found == false && currCart.Item_IDs [i] != itemid { //if item is not found and last item is also not equal to given item
		sendErrorResponse(w,errors.New("item_id invalid"),404)
		return
	}

	currCart.Item_IDs = newList
	change := bson.M{"$set": bson.M{"item_ids" : currCart.Item_IDs}}

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

// // GetUserCart serves the user cart GET request
// func (rc ResponseController) GetUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//  userid := p.ByName("userid")
//  fmt.Println("GET Request Cart: userid:", userid)
//
//  resp, err := getUserCartDB(userid, rc)
//  if err != nil {
// 	 sendErrorResponse(w,err,404)
// 	 return
//  }
//
//  jsonOut, _ := json.Marshal(resp)
//  httpResponse(w, jsonOut, 200)
//  fmt.Println("Response:", string(jsonOut), " 200 OK")
// }

// AddToUserCart serves the user cart add item PUT request
func (rc ResponseController) AddToUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 userid := p.ByName("userid")
 itemid := p.ByName("itemid")
 fmt.Println("PUT Request Add To Cart: userid:", userid, " item_id:", itemid)

 var newList []string
 currCart, err := getUserCartDB(userid, rc)
 size := 1
 if err == nil {
	 size = len(currCart.Item_IDs) + 1
	 newList = make([]string, size)
	 for i, x := range currCart.Item_IDs {
		newList[i] = x
	}
 } else {
	 currCart.Userid = userid
	 newList = make([]string, 1)
 }
 newList[size-1] = itemid
 currCart.Item_IDs = newList

 change := bson.M{"$set": bson.M{"item_ids" : currCart.Item_IDs}}

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

// // Login serves the Login GET request
// func (rc ResponseController) GetAllProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Println("GET Request product: allProducts")

 //	resp, err := getAllProductsDB(rc)
 //	if err != nil { 		sendErrorResponse(w,err,404)
 //		return
 //	}
 //	jsonOut, _ := json.Marshal(resp)
 //	httpResponse(w, jsonOut, 200)
 //	fmt.Println("Response:", string(jsonOut), " 200 OK")
 //}

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

// // SaveProduct serves the products POST request
// func (rc ResponseController) SaveProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	var req Product
//
// 	defer r.Body.Close()
//
// 	jsonIn, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		sendErrorResponse(w,err,400)
// 		return
// 	}
//
// 	json.Unmarshal([]byte(jsonIn), &req)
// 	fmt.Println("POST Request product:", req)
//
// 	if err := rc.session.DB("db_test").C("products").Insert(req); err != nil {
// 		sendErrorResponse(w,err,500)
// 		return
// 	}
// 	sendSuccessResponse(w,201)
// }

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

// CreateDocument serves the POST request
// func (rc ResponseController) CreateDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	var resp Response
// 	// var req PostRequest
// 	// var i interface{}
//
// 	defer r.Body.Close()
// 	jsonIn, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("Panic@CreateLocation.ioutil.ReadAll")
// 		panic(err)
// 	}
//
// 	// json.Unmarshal([]byte(jsonIn), &req)
// 	fmt.Println("POST Request:", string(jsonIn))
// 	// fmt.Println("POST Request:", string(jsonIn))
//
// 	resp.ID = bson.NewObjectId()
// 	//
// 	resp.Document = string(jsonIn)
// 	if err := rc.session.DB("db_test").C("col_test").Insert(resp); err != nil {
// 	    httpResponse(w, nil, 500)
// 		fmt.Println("Panic@CreateLocation.session.DB.C.Insert")
// 		panic(err)
// 	}
// 	jsonOut, _ := json.Marshal(resp)
// 	httpResponse(w, jsonOut, 201)
// 	fmt.Println("Response:", string(jsonOut), " 201 OK")
// }

// // GetLocation serves the GET request
// func (rc ResponseController) GetDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	fmt.Println("GET Request: ID:", id)
//
// 	resp, err := getDBData(id, rc)
// 	if err != nil {
// 	    w.WriteHeader(404)
// 		fmt.Println("Response: 404 Not Found")
// 		return
// 	}
//
// 	jsonOut, _ := json.Marshal(resp)
// 	httpResponse(w, jsonOut, 200)
// 	fmt.Println("Response:", string(jsonOut), " 200 OK")
// }

// UpdateLocation serves the PUT request
// func (rc ResponseController) UpdateDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	fmt.Println("PUT Request: ID:", id)
//
// 	var req PostRequest
// 	var resp Response
//
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(404)
// 		fmt.Println("Response: 404 Not Found")
// 		return
// 	}
//
// 	defer r.Body.Close()
// 	jsonIn, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("Panic@UpdateLocation.ioutil.ReadAll")
// 		panic(err)
// 	}
//
// 	json.Unmarshal([]byte(jsonIn), &req)
// 	fmt.Println("PUT Request:", req)
//
// 	//resp.Coordinate = req.Coordinate
// 	oid := bson.ObjectIdHex(id)
// 	resp.ID = oid;
//
// 	if err := rc.session.DB("db_test").C("col_test").UpdateId(oid, resp); err != nil {
// 		w.WriteHeader(404)
// 		fmt.Println("Response: 404 Not Found")
// 		return
// 	}
//
// 	jsonOut, _ := json.Marshal(resp)
// 	httpResponse(w, jsonOut, 201)
// 	fmt.Println("Response:", string(jsonOut), " 201 OK")
// }

// // DeleteLocation deletes existing user
// func (rc ResponseController) DeleteDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	fmt.Println("DELETE Request: ID:", id)
//
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(404)
// 		fmt.Println("Response: 404 Not Found")
// 		return
// 	}
//
// 	oid := bson.ObjectIdHex(id)
//
// 	if err := rc.session.DB("db_test").C("col_test").RemoveId(oid); err != nil {
// 		fmt.Println("Response: 404 Not Found")
// 		return
// 	}
//
// 	fmt.Println("Response: 200 OK")
// 	w.WriteHeader(200)
// }

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

// //Get data corresponding to the user id
// func getAllProductsDB(rc ResponseController) ([]Product, error) {
// 	var resp []Product
//
// 	if err := rc.session.DB("db_test").C("products").Find(nil).All(&resp); err != nil {
// 		return resp, err
// 	}
// 	return resp, nil
// }
//

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

//Get data corresponding to the object id
// func getDBData(id string, rc ResponseController) (Response, error) {
// 	var resp Response
// 	if !bson.IsObjectIdHex(id) {
// 		return resp, errors.New("404")
// 	}
// 	oid := bson.ObjectIdHex(id)
// 	if err := rc.session.DB("db_test").C("col_test").FindId(oid).One(&resp); err != nil {
// 		return resp, errors.New("404")
// 	}
// 	return resp, nil
// }

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
	//r.GET("/mongoserver/:id", rc.GetDocument)
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

	r.PUT("/mongoserver/cart/addItem/:userid/:itemid", rc.AddToUserCart)
	r.PUT("/mongoserver/cart/removeItem/:userid/:itemid", rc.RemoveFromUserCart)
	r.GET("/mongoserver/cart/:userid", rc.GetUserCart)
	// r.GET("/mongoserver/product/:productid", rc.GetProduct)
	// r.GET("/mongoserver/allProducts", rc.GetAllProducts)
	// r.POST("/mongoserver", rc.CreateDocument)
	// r.POST("/mongoserver/product", rc.SaveProduct)
	r.POST("/mongoserver/signup", rc.Signup)

	//r.POST("/mongoserver/cart", rc.Signup)
	// r.DELETE("/mongoserver/:id", rc.DeleteDocument)
	// r.PUT("/mongoserver/:id", rc.UpdateDocument)
	fmt.Println("Server is Ready !")
	fmt.Println(http.ListenAndServe(":7777",r))
}
