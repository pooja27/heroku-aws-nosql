/**
 * JSON FOR COFFEE-TEA-DRINKWARES DATA
 */

var app = angular.module('myApp',[]);
app.controller("Controller", ['$scope','$http','$window',function($scope,$http,$window) {
$scope.cart=[];
    $scope.caart=[];
    $scope.address=true;
    $scope.bill=0;
   var totalSum=function(){

        var sum=0;
       for(var i=0;i<$scope.cart.length;i++)
        sum=sum+$scope.cart[i].price;
        return sum;
    };
    var flag=[];

    $http({
        method: 'GET',
        url: '/alltea',

    }).
    then(function(response) {

        console.log(response.data);
        $scope.teacataloge=response.data;

    },function(response){

    });
    $http({
        method: 'GET',
        url: '/allDrinkware',

    }).
    then(function(response) {

        console.log(response.data);
        $scope.drinkcataloge=response.data;

    },function(response){

    });
    $http({
        method: 'GET',
        url: '/allCoffee',

    }).
    then(function(response) {

        console.log(response.data);
        $scope.coffecataloge=response.data;

    },function(response){

    });
    $scope.addCart=function(item_id,category,quantity){

        item={"item_id":item_id,"category":category,"quantity":quantity};
        $http({
            method: 'POST',url: '/addCart',
            data: item
        }).
        then(function(response) {
            angular.forEach(response.data.items,function(items){
                var count=nCount(items);

                if(! flag[items.item_id])
                {flag[items.item_id] = true;
                    $scope.cart.push({"item_id":items.item_id,"name":items.name,"count":count,"price":items.price*count});
                    $scope.totalPrice+=items.price*count;
                }


            })
            $scope.totalPrice=totalSum();
            function  nCount(items)
            {var count = response.data.items.filter(function(d) { return d.item_id === items.item_id; }).length;
                return count;}
            console.log($scope.cart);
            $scope.bagItems=response.data.items.length;
            $scope.noItem=false;
            $scope.placeOrder=false;

        },function(response){

        });
    }
    $scope.calcTotal = function(item){
        return 0+item.price;
    }
    $scope.cancel=function(){
        $scope.bill=$scope.cart;
        $scope.percentage=1;

        $http({
            method: 'POST',url: '/checkOut',
            data: $.param($scope.cart),
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
            }
        }).
        then(function(response) {
            $scope.orderPlaced=true;

            $scope.cart=[];
            $scope.bagItems=0;
            $scope.noItem = true;
            $scope.caart=true;
            $scope.payment=false;
            $scope.shipping=false;
            $scope.placeOrder=true;

        },function(response){

        });
    }
    $scope.buy=function(){
        $scope.bill=$scope.cart;
        $scope.percentage=100;

        $http({
            method: 'POST',url: '/checkOut',
            data: $scope.cart
        }).
        then(function(response) {
        $scope.orderPlaced=true;
    $scope.confirmOrder=response.data.data;
            $scope.cart=[];
            $scope.bagItems=0;
            $scope.noItem = true;
            $scope.caart=true;
            $scope.payment=false;
            $scope.shipping=false;
            $scope.placeOrder=true;
            $scope.orderShow=false;

        },function(response){
            $scope.confirmOrder=response.data.data;
            $scope.orderShow=true;
            $scope.orderPlaced=true;
            $scope.cart=[];
            $scope.bagItems=0;
            $scope.noItem = true;
            $scope.caart=true;
            $scope.payment=false;
            $scope.shipping=false;
            $scope.placeOrder=true;

        });
    }
    $scope.logout=function(){



        $http({
            method: 'GET',url: '/logout'
        }).
        then(function(response) {
            $window.location.href = '/';
        },function(response){

        });
    }
    $scope.percentage=30;
   $scope.rmvCart=function(item,quantity){

flag=[];
       $scope.cart=[];
       item={"item_id":item,"quantity":quantity};
        $http({
            method: 'POST',url: '/rmvCart',
            data: item
        }).
        then(function(response) {

            if(response.data.details){
                $scope.bagItems=0;
                $scope.noItem = true;
                $scope.placeOrder=true;
                $scope.totalPrice=0;
            }
            angular.forEach(response.data.items,function(items){
                var count=nCount(items);

                if(! flag[items.item_id]) {
                    flag[items.item_id] = true;
                    $scope.cart.push({"item_id": items.item_id,"name":items.name,"count": count, "price": items.price * count});
                    $scope.totalPrice += items.price * count;
                }


            })
            $scope.bagItems=response.data.items.length;

            $scope.totalPrice=totalSum();
            function  nCount(items)
            {var count = response.data.items.filter(function(d) { return d.item_id === items.item_id; }).length;
                return count;}
            console.log("***************"+$scope.cart);



        },function(response){


        });
    }

    $scope.caart=true;
    $scope.payment=false;
    $scope.shipping=false;
        $scope.noItem = true;
        $scope.placeOrder=true;

    $scope.cShop=function(){
        $scope.orderPlaced=false;
        $scope.cart=[];

    }
    $scope.rAdd=function(){
        $scope.shipping=true;
        $scope.caart=false;
        $scope.payment=false;
        $scope.percentage=50;
    }
    $scope.rCart=function(){
        $scope.caart=true;
        $scope.payment=false;
        $scope.shipping=false;
        $scope.percentage=30;
    }
    $scope.rPay=function(){
        $scope.caart=false;
        $scope.payment=true;
        $scope.shipping=false;
        $scope.percentage=80;
    }

$scope.getPercentage=function(){return $scope.percentage;}

        $http({
            method: 'GET',url: '/getCart'

        }).
        then(function(response) {
            //angular.forEach(response.data.items,function(items){
            //
            //});

            angular.forEach(response.data.items,function(items){
                var count=nCount(items);

                    if(! flag[items.item_id])
                    {flag[items.item_id] = true;
                    $scope.cart.push({"item_id":items.item_id,"name":items.name,"count":count,"price":items.price*count});

                    }



            })
            $scope.totalPrice=totalSum();
           function  nCount(items)
            {var count = response.data.items.filter(function(d) { return d.item_id === items.item_id; }).length;
                return count;}
            console.log($scope.cart);
           $scope.bagItems=response.data.items.length;
            $scope.placeOrder=false;
            $scope.noItem=false;

        },function(response){


        });


}]);
/*	Preloader
/* ========================================================================= */

jQuery(window).load(function(){

	$("#preloader").fadeOut("slow");

});

/* ========================================================================= */
/*  Welcome Section Slider
/* ========================================================================= */

$(function() {

    var Page = (function() {

        var $navArrows = $( '#nav-arrows' ),
            $nav = $( '#nav-dots > span' ),
            slitslider = $( '#slider' ).slitslider( {
                onBeforeChange : function( slide, pos ) {

                    $nav.removeClass( 'nav-dot-current' );
                    $nav.eq( pos ).addClass( 'nav-dot-current' );

                }
            } ),

            init = function() {

                initEvents();
                
            },
            initEvents = function() {

                // add navigation events
                $navArrows.children( ':last' ).on( 'click', function() {

                    slitslider.next();
                    return false;

                } );

                $navArrows.children( ':first' ).on( 'click', function() {
                    
                    slitslider.previous();
                    return false;

                } );

                $nav.each( function( i ) {
                
                    $( this ).on( 'click', function( event ) {
                        
                        var $dot = $( this );
                        
                        if( !slitslider.isActive() ) {

                            $nav.removeClass( 'nav-dot-current' );
                            $dot.addClass( 'nav-dot-current' );
                        
                        }
                        
                        slitslider.jump( i + 1 );
                        return false;
                    
                    } );
                    
                } );

            };

            return { init : init };

    })();

    Page.init();

});



$(document).ready(function(){

	/* ========================================================================= */
	/*	Menu item highlighting
	/* ========================================================================= */

	jQuery('#nav').singlePageNav({
		offset: jQuery('#nav').outerHeight(),
		filter: ':not(.external)',
		speed: 2000,
		currentClass: 'current',
		easing: 'easeInOutExpo',
		updateHash: true,
		beforeStart: function() {
			console.log('begin scrolling');
		},
		onComplete: function() {
			console.log('done scrolling');
		}
	});
	
    $(window).scroll(function () {
        if ($(window).scrollTop() > 400) {
            $(".navbar-brand a").css("color","#fff");
            $("#navigation").removeClass("animated-header");
        } else {
            $(".navbar-brand a").css("color","inherit");
            $("#navigation").addClass("animated-header");
        }
    });
	
	/* ========================================================================= */
	/*	Fix Slider Height
	/* ========================================================================= */	

    // Slider Height
    var slideHeight = $(window).height();
    
    $('#home-slider, #slider, .sl-slider, .sl-content-wrapper').css('height',slideHeight);

    $(window).resize(function(){'use strict',
        $('#home-slider, #slider, .sl-slider, .sl-content-wrapper').css('height',slideHeight);
    });
	
	
	
	$("#works, #testimonial").owlCarousel({	 
		navigation : true,
		pagination : false,
		slideSpeed : 700,
		paginationSpeed : 400,
		singleItem:true,
		navigationText: ["<i class='fa fa-angle-left fa-lg'></i>","<i class='fa fa-angle-right fa-lg'></i>"]
	});
	
	
	/* ========================================================================= */
	/*	Featured Project Lightbox
	/* ========================================================================= */

	$(".fancybox").fancybox({
		padding: 0,

		openEffect : 'elastic',
		openSpeed  : 650,

		closeEffect : 'elastic',
		closeSpeed  : 550,

		closeClick : true,
			
		beforeShow: function () {
			this.title = $(this.element).attr('title');
			this.title = '<h3>' + this.title + '</h3>' + '<p>' + $(this.element).parents('.portfolio-item').find('img').attr('alt') + '</p>';
		},
		
		helpers : {
			title : { 
				type: 'inside' 
			},
			overlay : {
				css : {
					'background' : 'rgba(0,0,0,0.8)'
				}
			}
		}
	});
	
});



// ========== END GOOGLE MAP ========== //

var wow = new WOW ({
	offset:       75,          // distance to the element when triggering the animation (default is 0)
	mobile:       false,       // trigger animations on mobile devices (default is true)
});
wow.init();

