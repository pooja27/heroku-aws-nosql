
/**
 * Module dependencies.
 */

var express = require('express')
  , routes = require('./routes')
  , user = require('./routes/user')
    ,userRegisteration = require('./routes/userRegisteration')
    ,products = require('./routes/products')
    ,cart = require('./routes/cart')
    ,billing = require('./routes/billing')
  , http = require('http')
  , path = require('path');

var app = express();
app.use(express.bodyParser());
app.use(express.cookieParser());
app.use(express.session({ secret: 'starbucks_001', duration: 30*60*1000}));
// all environments
app.set('port', process.env.PORT || 3000);
app.set('views', __dirname + '/views');
app.set('view engine', 'ejs');
app.use(express.favicon());
app.use(express.logger('dev'));
app.use(express.bodyParser());
app.use(express.methodOverride());
app.use(app.router);
app.use(express.static(path.join(__dirname, 'public')));



// development only
if ('development' == app.get('env')) {
  app.use(express.errorHandler());
}

//User Sign up APIs
app.get('/', routes.index);
app.post('/signup',userRegisteration.signup);
app.post('/login',userRegisteration.login);
app.get('/allCoffee',products.coffees);
app.get('/alltea',products.teas);
app.get('/allDrinkware',products.drinkwares);
app.post('/addCart',cart.addCart);
app.get('/getCart',cart.getCart);
app.get('/logout',userRegisteration.logout);
app.post('/rmvCart',cart.removeCart);
app.post('/checkOut',billing.checkOut);

//app.get('*',errorPage.errorPage);


http.createServer(app).listen(app.get('port'), function(){
  console.log('Express server listening on port ' + app.get('port'));
});
