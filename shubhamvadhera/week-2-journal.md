## Week #2
* I worked on setting up AWS for team 7 which included following tasks:
	* Setting up AWS instances on which MongoDB would be installed
	* Setting up the required keypair, VPC, subnets, etc.
* Also worked on setting up MongoDB on AWS
	* MongoDB was successfully deployed to AWS, which included:
		* One NAT node and one Primary replica
* Riak DB was studied for its architecture and feasibility in this project
	* Fault-tolerant availability
	* Querying
	* Key/Value
	* Buckets
	* Replication
	* Partitions
	* Consistent Hashing
	* CAP sacrifices
	* RIAK drivers for Node JS - community supported, not official
* Result - Riak is a very specialized DB and is more suited for high performance and large scale systems, so probably not required for our project.
* Preparation for Team demo presentation on MongoDB, which included covering the following areas
	* Design philosophy of MongoDB
	* Architecture
	* Data Model
* After discussion with professor, we got more clarity and insights towards databases to implement:
	* There's only one DB needed for the backend, so MongoDB would be able suffecient to implement the features of the project.
	* DB team members to stop exploring other DBs and both team members would now work solely on MongoDB
	* Also, the backend DB needs to have a REST fronted for Heroku deployed services
* For week 3
	* Need to implement a front-end REST API for MongoDB so that the system deployed on Heroku can interact with MongoDB
	* Need to finalize the language, MongoDB drivers to implement this