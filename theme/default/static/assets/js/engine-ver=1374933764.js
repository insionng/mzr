jQuery(function ($) {
	
	"use strict";
	
/*-----------------------------------------------------------------------------------*/
/*	Videos & Audios
/*-----------------------------------------------------------------------------------*/
	
	
	$("body").fitVids();

	$('video, audio').mediaelementplayer();

/*-----------------------------------------------------------------------------------*/
/*	Twitter Feeds
/*-----------------------------------------------------------------------------------*/

	$('.engine_tweet').each( function() {
	
		$(this).tweet({
			username: $(this).attr('data-username'),
			join_text: "auto",
			avatar_size: 32,
			count: $(this).attr('data-number'),
			auto_join_text_default: "", 
			auto_join_text_ed: "",
			auto_join_text_ing: "",
			auto_join_text_reply: "",
			auto_join_text_url: "",
			loading_text: $(this).attr('data-loading')
		});
    
	}); 

/*-----------------------------------------------------------------------------------*/
/*	Hover Effect
/*-----------------------------------------------------------------------------------*/    
	
	$('.featured-image a img').hover( function() {
		
		//$(this).animate({ opacity: 0.8 }, 200);
		
	}, function() {
		
		//$(this).animate({ opacity: 1 }, 200);
		
	});

	
	function colourPicker(headerColour, primaryColour, secondaryColour) {
		
		var html = ".top-bar,.signup button{background-color:"+headerColour+"}.bg{background-image: url(http://demo2.designerthemes.com/buzz/files/2013/02/map.png);background-repeat: no-repeat;background-position: top center;background-attachment: scroll;background-color:"+headerColour+"}.social.nav a:hover{color:"+headerColour+"}a,a:hover,.nav.nav-pills .dropdown-menu a:hover,.widget-title,#reply-title,.first .entry-title a,.comments label,.comments cite,.trail-end{color:"+primaryColour+"}.pager-container li:hover,.pager-container .pagination a:hover,.comments #submit,.read-more,.flex-control-paging li a.flex-active,.Engine_Postcarousel.primary,.Engine_Postcarousel.primary .widget-title,.Engine_Postcarousel.primary .widget-title .title-bg,.Engine_Postcarousel.primary .flex-control-paging,.Engine_Tabbed .topics .cat-item:hover a,.widget_categories .cat-item:hover a{background:"+primaryColour+"}.header a:hover,.comments a:hover,.comments #submit:hover,.read-more:hover,.primary-menu .nav li.current-menu-item>a,.signup button:hover,.footer a:hover,.secondary-menu a:hover{color:"+secondaryColour+"}.header-search .btn-warning,.Engine_Tabbed .topics .cat-item:hover .post-count,.widget_categories .cat-item:hover .post-count,.footer .Engine_Postcarousel .flex-control-nav li a.flex-active,.footer .Engine_Postcarousel .flex-control-nav li a.hover{background:"+secondaryColour+"}.signup,.flex-direction-nav .flex-prev:hover,.flex-direction-nav .flex-next:hover,.Engine_Postcarousel.secondary,.Engine_Postcarousel.secondary .widget-title,.Engine_Postcarousel.secondary .widget-title .title-bg,.Engine_Postcarousel.secondary .flex-control-paging,.mobile-menu .btn-navbar,.navbar .btn-navbar:hover,.navbar .btn-navbar:active,.mobile-menu li a,.post-wrap.first .archive-title{background-color:"+secondaryColour+"}.bg,.footer{border-color:"+secondaryColour+"}";
		
		$('#color-scheme').html(html);
		
	}
	
	$('.red-skin').click( function() {
		colourPicker('#d31e1e', '#d31e1e', '#ffb400');
		return false;
	});
	
	$('.black-skin').click( function() {
		colourPicker('#2d2d2d', '#2d2d2d', '#ffb400');
		return false;
	});
	
	$('.green-skin').click( function() {
		colourPicker('#789048', '#789048', '#ffb400');
		return false;
	});
	
	$('.blue-skin').click( function() {
		colourPicker('#00a0b0', '#00a0b0', '#EDC951');
		return false;
	});
	
	$('.pink-skin').click( function() {
		colourPicker('#fe4365', '#fe4365', '#7FC7AF');
		return false;
	});
	
	$('.grey-skin').click( function() {
		colourPicker('#444444', '#444444', '#aaaaaa');
		return false;
	});

/*-----------------------------------------------------------------------------------*/
/*	Flexslider
/*-----------------------------------------------------------------------------------*/
	
	
	$('.gallery-flex').imagesLoaded( function() {
		
		$(this).flexslider({
			animation: "slide",
			controlNav: false,
			smoothHeight: true,
			directionNav: true,
			slideshow: false,
			prevText: '<i class="icon-chevron-left"></i>', 
			nextText: '<i class="icon-chevron-right"></i>',
			start: function(slider) {
				
				$('.slides li img').click(function(event){
					event.preventDefault();
					slider.flexAnimate(slider.getTarget("next"));
				});
			}
		});
    
	});
	
	$('.home-slider .flexslider').imagesLoaded( function() {
		
		var autoplay = parseInt($(this).attr('data-autoplay'), 10) * 1000;
		var sh = false;
		var effect = $(this).attr('data-effect');

		if( autoplay !== 0 ) {
			sh = true;	
		}
		
		if( effect === '' ) {
			effect = 'fade';	
		}
		
		jQuery(this).flexslider({
			animation: effect,
			slideshow: sh,
			video: true,
			slideshowSpeed: autoplay,
			smoothHeight: false,
			controlNav: false,
			directionNav: true,
			animationLoop: true,
			prevText: '<i class="icon-angle-left"></i>',
			nextText: '<i class="icon-angle-right"></i>'
		});
    
	});
	
	$('.home-widgets .post-carousel, .related-posts').imagesLoaded( function() {
		
		jQuery(this).flexslider({
			animation: "slide",
			slideshow: false,
			animationLoop: true,
			itemWidth: 153,
			itemMargin: 20,
			minItems: 1,
			maxItems: 3,
			move: 3,
			smoothHeight: false,
			directionNav: false
		});
		
	});
	
	$('.sidebar .post-carousel, .footer .post-carousel').imagesLoaded( function() {
		
		jQuery(this).flexslider({
			animation: "slide",
			slideshow: false,
			animationLoop: true,
			smoothHeight: true,
			directionNav: false
		});
		
	});  
		
/*-----------------------------------------------------------------------------------*/
/*	Tabs
/*-----------------------------------------------------------------------------------*/
	
	
	$('.engine-tabs li:first-child a').tab('show', 200);

	$('.engine-tabs a').click(function (e) {
		e.preventDefault();
		$(this).tab('show');
		
	});
	

/*-----------------------------------------------------------------------------------*/
/*	Tooltips
/*-----------------------------------------------------------------------------------*/
	
	
	$('.engine-tooltip').tooltip();
	

/*-----------------------------------------------------------------------------------*/
/*	Dropdowns
/*-----------------------------------------------------------------------------------*/
	
	
	$('.primary-menu .nav li, .secondary-menu .nav li').hover( function () {
		$(this).find('.dropdown-menu').stop().toggle(100);
	});

	
});



