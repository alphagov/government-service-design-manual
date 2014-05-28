// Used on the "Form example - Registration" page in user-centred design
// resources. Takes the text of the clicked 'option' and assigns it as
// a class to the element named in the 'data-for' attribute
$(function () {
  $('.class-toggle .option').click(function(){
    $('.class-toggle .option').removeClass('selected');
    $(this).addClass('selected');
    var selectedClass = $(this).text();
    var selectedElement = "#" + $(this).parents('.class-toggle').data("for");
    $(selectedElement).removeClass().addClass(selectedClass);
    return false;
  });
});
