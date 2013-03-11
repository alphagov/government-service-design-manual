// Emphasised checkbox and radio button label styles
    
var $emphasised = $('.emphasised input');
$emphasised.filter(':checked').parent().addClass('checked');
$emphasised.change(function() {
    $emphasised.parent().removeClass('checked');
    $emphasised.filter(':checked').parent().addClass('checked');
});