MongoDB REST endpoints

1.	/mongoserver/signup
*	Purpose - To save new user data
*	POST
{
    "userid" : "jagg",
    "password" : "jaggpass",
    "email" : "jagg@sjsu.edu",
    "name" : "jaggi"
}
*	Returns:
*	success
{
		status : “success”,
		details : “”
}
*	failure
{
	status : “failure”,
	details : “<failure message>”
}

2.	/mongoserver/login/<userid>
*	Purpose - To return user data by user id
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
	userid : “”,
	password : “”,
	email : “”,
	name : “”
}

3.	/mongoserver/teas
*	Purpose - To get all tea data
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
[
  {
    "category": "tea",
    "price": 50,
    "count": 12,
    "item_id": "tea_1",
    "name": "Teavana Peach Tranquility Full-Leaf Sachets",
    "brand": "tazo",
    "type": "ice tea",
    "tea_form": "K-Cup Pods"
  },
  {
    "category": "tea",
    "price": 8.55,
    "count": 12,
    "item_id": "tea_2",
    "name": "Teavana Oprah Cinnamon Chai Full-Leaf Sachets",
    "brand": "teavana",
    "type": "black tea",
    "tea_form": "tea bags"
  }
]

4.	/mongoserver/coffees
*	Purpose - To get all coffee data
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
[
  {
    "region": "latin america",
    "category": "coffee",
    "price": 40,
    "item_id": "coffee_1",
    "name": "Sulawesi, Whole Bean",
    "flavor": "flavored",
    "quantity": 5,
    "roast": "blonde",
    "type": "decaffinated"
  },
  {
    "region": "multi",
    "category": "coffee",
    "price": 50,
    "item_id": "coffee_2",
    "name": "3 Region Blend, Whole Bean",
    "flavor": "flavored",
    "quantity": 5,
    "roast": "medium",
    "type": "regular"
  }
]

5.	/mongoserver/drinkwares
*	Purpose - To get all drinkware data
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
[
  			{
    				"category": "drinkware",
"price": 15.95,
    				"item_id": "drinkware_1",
    				"name": "Spring Garden Traveler, 12 fl oz"
  			},
{
    				"category": "drinkware",
    				"price": 18.55,
    				"item_id": "drinkware_2",
    				"name": "Stainless Steel Bouquet Tumbler, 16 fl oz"
  			}
]

6.	/mongoserver/inventory/:itemid
*	Purpose - To get current stock for a item_id
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  			"item_id": "coffee_9",
  			"stock": 106
}

7.	/mongoserver/inventory/addOne/:itemid
*	Purpose - To increment the current stock for a item_id by 1
*	PUT
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  			"item_id": "coffee_9",
  			"stock": 107
}

8.	/mongoserver/inventory/removeOne/:itemid
*	Purpose - To decrement the current stock for a item_id by 1
*	PUT
*	Returns:
*	failure
{
  "status": "failure",
  "details": "Item coffee_9 stock is 0. Cannot process."
}
*	success
{
  			"item_id": "coffee_9",
  			"stock": 105
}

9.	/mongoserver/tea/:itemid
*	Purpose - To get tea data
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
"category": "tea",
  		"price": 50,
  		"count": 12,
  		"item_id": "tea_1",
  		"name": "Teavana Peach Tranquility Full-Leaf Sachets",
  		"brand": "tazo",
  		"type": "ice tea",
  		"tea_form": ""
}

10.	/mongoserver/coffee/:itemid
*	Purpose - To get coffee data
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  		"region": "latin america",
  		"category": "coffee",
  		"price": 40,
  		"item_id": "coffee_1",
  		"name": "Sulawesi, Whole Bean",
  		"flavor": "flavored",
  		"quantity": 5,
  		"roast": "blonde",
  		"type": "decaffinated"
}

11.	/mongoserver/drinkware/:itemid
*	Purpose - To get drinkware data
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  "category": "drinkware",
  "price": 15.95,
  "item_id": "drinkware_1",
  "name": "Spring Garden Traveler, 12 fl oz"
}

12.	/mongoserver/cart/addItem/:userid/:itemid
*	Purpose - To add one item to cart against a user
*	PUT
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  	"userid": "zuser@gmail.com",
  	"item_ids": [
    	"coffee_13",
    	"coffee_14",
    	"coffee_15",
    	"coffee_16"
  	]
}

13.	/mongoserver/cart/removeItem/:userid/:itemid
*	Purpose - To remove one item from cart against a user
*	PUT
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  	"userid": "zuser@gmail.com",
  	"item_ids": [
    	"coffee_13",
    	"coffee_14",
    	"coffee_15",
  	]
}
*	success – cart deleted
{
  "status": "success",
  "details": "cart for user zuser@gmail.com deleted"
}

14.	/mongoserver/cart/:userid
*	Purpose - To get the cart against a user
*	GET
*	Returns:
*	failure
{
	status : “failure”,
	details : “<failure message>”
}
*	success
{
  	"userid": "zuser@gmail.com",
  	"item_ids": [
    	"coffee_13",
    	"coffee_14",
    	"coffee_15",
  	]
}
