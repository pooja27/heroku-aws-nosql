MongoDB Collections
1.	users – stores user details including userid and password.
{
	userid : “”,
	password : “”,
	email : “”,
	name : “”
}
2.	tea – stores details of tea
{
    "category": "tea",
    "price": 50,
    "count": 12,
    "item_id": "tea_1",
    "name": "Teavana Peach Tranquility Full-Leaf Sachets",
    "brand": "tazo",
    "type": "ice tea",
    "tea_form": "K-Cup Pods"
  }
3.	coffee – stores details of coffee
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
4.	drinkware – stores details of drinkwares
{
    "category": "drinkware",
    "price": 15.95,
    "item_id": "drinkware_1",
    "name": "Spring Garden Traveler, 12 fl oz"
 }

5.	inventory – stores the item stock
{
	"item_id" : "coffee_9",
	"stock" : 7
}

6.	cart – stores the user cart
{
        "userid" : "kuser@gmail.com",
        "item_ids" : [
                "coffee_15",
                "tea_13"
        ]
}
