/**
 * Created by Jagmohan on 4/9/16.
 */
var mongo = require("./mongo");
var mongoURL = "mongodb://localhost:27017/cmpe281";
exports.login=function(req,res){

    var email = req.param('username');
    var password = req.param('password');
    
    mongo.connect(mongoURL, function(){
        console.log('Connected to mongo at: ' + mongoURL);
        var coll = mongo.collection('userDetails');
        coll.findOne({"email:": email, "password":password }, function(err, user){
            if(user)
            {
                console.log(user.email); 
                res.status(200).send({"status":"Login Successful"});
            } 
            else
            {
                 res.status(401).send({"status":"Login Failed"});
                
            }
        
        });
        
    });
   

};

exports.signup=function(req,res) {
    var email = req.param('username');
    var password = req.param('password');
    var firstName = req.param('firstName');
    var lastName = req.param('lastName');
    var mobileNumber = req.param('mobileNumber');
    console.log(email);
    console.log(password);
    console.log(firstName);
    console.log(lastName);
    console.log(mobileNumber);
    mongo.connect(mongoURL, function(){
        console.log('Connected to mongo at: ' + mongoURL);
        var coll = mongo.collection('userDetails');
        coll.findOne({email: email}, function(err, user){
            if(user)
            {
                // checking if user already exists
                json_responses = {"statusCode" : 401};
                res.send(json_responses);
            }
            else
            {    
                
                coll.insertOne({"email": email,"firstName":firstName,"lastName":lastName,"password":password,"mobileNumber":mobileNumber}, function(err, user){
                    if(err)
                    {
                        throw err;
                    }
                    else
                        res.status(200).send({"status":"Registeration Success"});
                });
            }
            });
    });
    
};


