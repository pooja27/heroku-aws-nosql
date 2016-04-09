## Week #1
* Team project to be a Starbucks shopping cart application
* 2 team members will work on Heroku, Node.js, other 2 on Backend DB
* Second DB not yet decided, so both DB team members to start working on MongoDB
	* To be decided within next few weeks - Riak, Cassandra, Redis
* Team deciding on features of Starbucks Application:
	* Must have features - login page, product catalogue, user sign up
		* To store user data, address or only email is sufficient ?
	* Additional - UI improvement
* Exploring MongoDB
	* MongoDB architecture
		* Strong consistency - shopping cart ?
		* Flexible data model - Documents, BSON, json
	* MongoDB query model
		* Native support for node.js - MongoDb drivers
		* Mongo Shell ?
		* MongoDB compass for testing
		* Indexes ?
			* Unique index - key
	* Data management
		* Auto-sharding - drivers
* Create sample database on MongoLab
	* Monglo lab can be connected to Node.js
* UI wirefreams
	* Product catalogue
		* Must have - Products, price, cart icon on top right, single scroll page ?, descriptions
		* Should have - sort by price, sort by latest, ratings and reviews, images
		* Can have - make custom cofee
	* Cart checkout
		* Must have - number of items, price, quantity, total, buy button, go back, remove from cart
		* Extend - pay on separate page - Credit card, address, confirm order
* For week 2 - 
	* explore Riak
	* MongoDB connectivity with Node.js