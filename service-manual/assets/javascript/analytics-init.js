(function() {
  "use strict";

  GOVUK.Analytics.load();
  var cookieDomain = (document.domain === 'www.gov.uk') ? '.www.gov.uk' : document.domain;

  GOVUK.analytics = new GOVUK.Analytics({
    universalId: 'UA-26179049-1',
    cookieDomain: cookieDomain
  });

  if (window.devicePixelRatio) {
    GOVUK.analytics.setDimension(11, window.devicePixelRatio);
  }

  GOVUK.analytics.trackPageview();
})();
