/*
Anti-spam plugin
No spam in comments. No captcha.
wordpress.org/extend/plugins/anti-spam/
*/

jQuery(function($){

    $('.comment-form-anti-spam, .comment-form-anti-spam-2').hide(); // hide inputs from users
    var answer = $('.comment-form-anti-spam input#anti-spam-a').val(); // get answer
    $('.comment-form-anti-spam input#anti-spam-q').val( answer ); // set answer into other input

    if ( $('#comments form input#anti-spam-q').length == 0 ) { // anti-spam input does not exist (could be because of cache or because theme does not use 'comment_form' action)
        var current_date = new Date();
        var current_year = current_date.getFullYear();
        $('#comments form').append('<input type="hidden" name="anti-spam-q" id="anti-spam-q" value="'+current_year+'" />'); // add whole input with answer via javascript to comment form
    }

    if ( $('#respond form input#anti-spam-q').length == 0 ) { // similar, just in case (used because user could bot have #comments)
        var current_date = new Date();
        var current_year = current_date.getFullYear();
        $('#respond form').append('<input type="hidden" name="anti-spam-q" id="anti-spam-q" value="'+current_year+'" />'); // add whole input with answer via javascript to comment form
    }

    if ( $('form#commentform input#anti-spam-q').length == 0 ) { // similar, just in case (used because user could bot have #respond)
        var current_date = new Date();
        var current_year = current_date.getFullYear();
        $('form#commentform').append('<input type="hidden" name="anti-spam-q" id="anti-spam-q" value="'+current_year+'" />'); // add whole input with answer via javascript to comment form
    }

});