var GOVUK = window.GOVUK || {};

GOVUK.helpPopup = {
  selectSubsection : function (e) {
    var $this = $(this);
    
    if ($this.parents('ul.subsections').length < 1) {
      e.preventDefault();
      window.location.href = '#/';
      slug = 'popular';
      $('.subsections')
        .find('.active')
        .removeClass('active');
    } else {
      var $section = $(this).parent(),
          slug = $section
                  .data('filter')
                  .match(/section-([\w-]+)/)
                  .pop();

      $section
        .addClass('active')
        .siblings()
        .removeClass('active');
    }

    $('#' + slug)
      .removeClass('hidden')
      .siblings()
      .addClass('hidden');
  },
  init : function () {
    var hash = window.location.hash,
        instObj = this;

    // Help sections click actions
    if ($('.help-sections').length > 0) {
      $('.help-navigation')
        .find('li.open')
        .on('click', 'a', function (e) {
          instObj.selectSubsection.call(this, e); 
        });  
    }

    if (hash !== '' && hash !== '#/') {
      // on page load parse hash
      hash = hash.substring(2);

      $('.subsections li').filter(function() {
        return $(this).data('filter') == 'section-' + hash;
      })
        .find('a')
        .click();
    }
    
    $('.help-navigation')
      .find('ul:eq(0) li ul li')
      .each(function (idx) {
        var section = $(this).find('a').text(),
            slug = $(this) 
              .data('filter')
              .match(/section-([\w-]+)/)
              .pop();
      });
  }
};

$(document).ready(function() {
  GOVUK.helpPopup.init();
});
