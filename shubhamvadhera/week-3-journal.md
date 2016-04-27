## Week #3
* This week I worked on an issue with AWS MongoDB setup
	* The MongoDB stops working at odd times
	* Had to reconfigure and install MongoDB from scratch
* I started working on MongoDB REST endpoints for Heroku frontend
	* The language of coding would be GO. I am using mgo drivers for the same.
	* The team discussed and finalized initial structure of the schema
	* There will be a users collection that will store the details of the users - userid, password, email, name
	* This collection would expose to REST end points:
		* A GET endpoint - /mongoserver/login/<userid>
			* This will be used to return the user details to the Heroku front-end
			* This will be used during the sign in process - when a user logs in to the website on Heroku, the Heroku service would check the credentials against the database using this REST end point.
		* A POST endpoint - /mongoserver/signup
			* This will be used to accept the user details from the Heroku front-end to be saved in the MongoDB
			* This will be used during the new user sign up process - when a new user tries to register fot the website on Heroku, the Heroku service would store the credentials in the database using this REST end point.
* For week 4
	* Need to implement more REST endpoints after discussion with fronted end team
	* Need to finish backend MongoDB work