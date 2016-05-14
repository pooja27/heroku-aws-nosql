
/*
 * GET home page.
 */

exports.index = function(req, res){
  if(req.session.data)
  res.render('home1');
  else
  res.render('login');
};