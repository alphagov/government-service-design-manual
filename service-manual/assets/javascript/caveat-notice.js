function setCookie(name, value, days){
  var cookieString = name + "=" + value + "; path=/";
  if (days) {
    var date = new Date();
    date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
    cookieString = cookieString + "; expires=" + date.toGMTString();
  }
  if (document.location.protocol == 'https:'){
    cookieString = cookieString + "; Secure";
  }
  document.cookie = cookieString;
}

function getCookie(name){
  var nameEQ = name + "=";
  var ca = document.cookie.split(';');
  for(var i = 0, len = ca.length; i < len; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1, c.length);
    }
    if (c.indexOf(nameEQ) === 0) {
      return c.substring(nameEQ.length, c.length);
    }
  }
  return null;
}

$(function() {
  var $message = $('#caveat-notice');

  if ($message.length && getCookie('govuk_service_manual_beta') === null) {
    $message.show();
    setCookie('govuk_service_manual_beta', 'yes', 28);
  }
});
