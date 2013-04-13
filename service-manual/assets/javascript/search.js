$(function() {
  /*
  $(".js-search-focus")
    .on("focus",function(e){
      $(e.target).addClass("hide-label")
    })
    .on("blur",function(e){
      if((e.target).val() != ""){
        $(e.target).removeClass("hide-label")
      }
    })
*/


  var $searchFocus = $('.js-search-focus');
  $searchFocus.on('focus', function(e){
    $(e.target).addClass('focus');
  });
  $searchFocus.on('blur', function(e){
    if($searchFocus.val() === ''){
      $(e.target).removeClass('focus');
    }
  });


});


