/**
 * Created by Jagmohan on 4/9/16.
 */
exports.login=function(req,res){

    var email = req.param('username');
    var password = req.param('password');
    console.log(email);
    console.log(password);
    res.status(200).send({"status":"Success"});

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
    res.status(200).send({"status":"Success"});
};


