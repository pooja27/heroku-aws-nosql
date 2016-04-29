
/*
 * GET home page.
 */

exports.index = function(req, res){
  if(req.session.data)
  res.render('home');
  else
  res.render('login');
};