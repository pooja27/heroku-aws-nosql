package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* ------- Debugging variables ------- */
var out io.Writer
var debugModeActivated bool

/* ------- Structs ------- */

//struct to store in products collection
type Product struct {
	ProductID		string	`json:"productid" bson:"productid"`
	Name				string 	`json:"name" bson:"name"`
	Type				string 	`json:"type" bson:"type"`
	Description string 	`json:"description" bson:"description"`
	Price 			float64	`json:"price" bson:"price"`
}

//struct to return success / failure status
type ResponseStatus struct {
	Status 	string `json:"status" bson:"status"`
	Details string `json:"details" bson:"details"`
}

//to store in users collection
type UserDetails struct {
	Userid		string `json:"userid" bson:"userid"`
	Password	string `json:"password" bson:"password"`
	Email 		string `json:"email" bson:"email"`
	Name 			string `json:"name" bson:"name"`
}

//ResponseController struct to provide to httprouter
type ResponseController struct {
	session *mgo.Session
}

/* ------- REST Functions ------- */

// Login serves the Login GET request
func (rc ResponseController) GetAllProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET Request product: allProducts")

	resp, err := getAllProductsDB(rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// Login serves the Login GET request
func (rc ResponseController) GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productid := p.ByName("productid")
	fmt.Println("GET Request product: productid:", productid)

	resp, err := getProductDB(productid, rc)
	if err != nil {
		sendErrorResponse(w,err,404)
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

// SaveProduct serves the products POST request
func (rc ResponseController) SaveProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req Product

	defer r.Body.Close()

	jsonIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w,err,400)
		return
	}

	json.Unmarshal([]byte(jsonIn), &req)
	fmt.Println("POST Request product:", req)

	if err := rc.session.DB("db_test").C("products").Insert(req); err != nil {
		sendErrorResponse(w,err,500)
		return
	}
	sendSuccessResponse(w,201)
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

// DeleteLocation deletes existing user
func (rc ResponseController) DeleteDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println("DELETE Request: ID:", id)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		fmt.Println("Response: 404 Not Found")
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := rc.session.DB("db_test").C("col_test").RemoveId(oid); err != nil {
		fmt.Println("Response: 404 Not Found")
		return
	}

	fmt.Println("Response: 200 OK")
	w.WriteHeader(200)
}

/* ------- Helper Functions ------- */

//Get data corresponding to the user id
func getAllProductsDB(rc ResponseController) ([]Product, error) {
	var resp []Product

	if err := rc.session.DB("db_test").C("products").Find(nil).All(&resp); err != nil {
		return resp, err
	}
	return resp, nil
}

//Get data corresponding to the user id
func getProductDB(productid string, rc ResponseController) (Product, error) {
	var resp Product

	if err := rc.session.DB("db_test").C("products").Find(bson.M{"productid" : productid}).One(&resp); err != nil {
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
	r.GET("/mongoserver/product/:productid", rc.GetProduct)
	r.GET("/mongoserver/allProducts", rc.GetAllProducts)
	// r.POST("/mongoserver", rc.CreateDocument)
	r.POST("/mongoserver/product", rc.SaveProduct)
	r.POST("/mongoserver/signup", rc.Signup)
	// r.DELETE("/mongoserver/:id", rc.DeleteDocument)
	// r.PUT("/mongoserver/:id", rc.UpdateDocument)
	fmt.Println("Server is Ready !")
	http.ListenAndServe(":7777", r)
}
