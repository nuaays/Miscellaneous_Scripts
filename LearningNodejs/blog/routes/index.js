var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

/*hello world*/
router.get('/ns', function(req, res){
  res.send("Hello World!");
});

/*registry*/
router.get('/reg', function(req, res){
  res.render('index', { title: '注册'});
});
module.exports = router;
