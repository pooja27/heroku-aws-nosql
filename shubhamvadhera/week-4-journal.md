## Week #4
* Team meeting
	* Discussion regarding additional MongoDB REST endpoints needed
	* There will be 3 kinds of products - tea, coffee, drinkware
	* Each type of product will be stored in a separate collection - tea, coffee, drinkware
		* Details of product as per official starbucks data
	* Inventory will be stored in separate collection - inventory
	* Shopping cart will be persisited in DB so that it is available across all devices - cart collection

* REST end points for products
	* GET - teas
	* GET - coffees
	* GET - drinkwares
		* Returns all products of given type
* REST end points for specific product
	* GET - tea
	* GET - coffee
	* GET - drinkware
		* Returns a product of given item_id
* REST end points for Inventory
		* GET - Returns the inventory against an item_id
* REST end points for Inventory updates
		* PUT - To increase inventory of an item by one
		* PUT - To decrease inventory of an item by one
* REST end points for Cart
		* GET - To get the cart against a user
* REST end points for Cart updates
		* PUT - To update the cart against a user, adding one item given the item_id
		* PUT - To update the cart against a user, removing one item given the item_id
		* If cart becomes empty, cart is deleted

* Worked on scalability of MongoDB cluster on Amazon AWS

* For last week
	* Run the project end to end and complete project