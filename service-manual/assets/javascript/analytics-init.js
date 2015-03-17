(function() {
  "use strict";

  GOVUK.Tracker.load();
  var cookieDomain = (document.domain === 'www.gov.uk') ? '.www.gov.uk' : document.domain;

  GOVUK.analytics = new GOVUK.Tracker({
    universalId: 'UA-26179049-7',
    classicId: 'UA-26179049-1',
    cookieDomain: cookieDomain
  });

  if (window.devicePixelRatio) {
    GOVUK.analytics.setDimension(11, window.devicePixelRatio, 'Pixel Ratio', 2);
  }

  GOVUK.analytics.trackPageview();
})();
