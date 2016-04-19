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

//Debugging variables
var out io.Writer
var debugModeActivated bool

//Response struct
type Response struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Coordinate Point         `json:"coordinate" bson:"coordinate"`
}

//Point struct to hold coordinates
type Point struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

//PostRequest struct to handle POST data
type PostRequest struct {
	Coordinate  Point   `json:"coordinate"`
}

//ResponseController struct to provide to httprouter
type ResponseController struct {
	session *mgo.Session
}

//NewResponseController function returns reference to ResponseController and a mongoDB session
func NewResponseController(s *mgo.Session) *ResponseController {
	return &ResponseController{s}
}

func getSession() *mgo.Session {
    //Enter mongoLab connection string here
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Panic@getSession.Dial")
		panic(err)
	}
	return s
}

// CreateLocation serves the POST request
func (rc ResponseController) CreateDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp Response
	var req PostRequest

	defer r.Body.Close()
	jsonIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Panic@CreateLocation.ioutil.ReadAll")
		panic(err)
	}

	json.Unmarshal([]byte(jsonIn), &req)
	fmt.Println("POST Request:", req)

	resp.ID = bson.NewObjectId()

	resp.Coordinate = req.Coordinate
	if err := rc.session.DB("db_test").C("col_test").Insert(resp); err != nil {
	    httpResponse(w, nil, 500)
		fmt.Println("Panic@CreateLocation.session.DB.C.Insert")
		panic(err)
	}
	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 201)
	fmt.Println("Response:", string(jsonOut), " 201 OK")
}

// GetLocation serves the GET request
func (rc ResponseController) GetDocument(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println("GET Request: ID:", id)

	resp, err := getDBData(id, rc)
	if err != nil {
	    w.WriteHeader(404)
		fmt.Println("Response: 404 Not Found")
		return
	}

	jsonOut, _ := json.Marshal(resp)
	httpResponse(w, jsonOut, 200)
	fmt.Println("Response:", string(jsonOut), " 200 OK")
}

//Get data corresponding to the object id
func getDBData(id string, rc ResponseController) (Response, error) {
	var resp Response
	if !bson.IsObjectIdHex(id) {
		return resp, errors.New("404")
	}
	oid := bson.ObjectIdHex(id)
	if err := rc.session.DB("db_test").C("col_test").FindId(oid).One(&resp); err != nil {
		return resp, errors.New("404")
	}
	return resp, nil
}

//write http response
func httpResponse(w http.ResponseWriter, jsonOut []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", jsonOut)
}

func main() {
	//debugging variables----------------------
	debugModeActivated = true
	out = ioutil.Discard
	if debugModeActivated {
		out = os.Stdout
	}
	//---------------------debugging variables

    fmt.Fprintln(out, "Starting server...")
	r := httprouter.New()
	rc := NewResponseController(getSession())
	r.GET("/mongoserver/:id", rc.GetDocument)
	r.POST("/mongoserver", rc.CreateDocument)
	fmt.Fprintln(out, "Server is Ready !")
	http.ListenAndServe(":7777", r)
}
