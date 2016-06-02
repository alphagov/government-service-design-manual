---
layout: detailed-guidance
title: Browsers and devices
subtitle: Which ones to test with, and how best to support them
category: user-centred-design
type: guide
audience:
  primary: designers, developers, qa, service-managers
  secondary: delivery-managers, performance-analysts, content-designers
status: draft
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
exclude_from_search: true
---

Services should be universally accessible, regardless of how the user is choosing to access them.

Due to the large range of browsers, devices and screen sizes available, users' experience of your service will vary according to the technical capabilities of their browsers and devices. You must verify that your service works across a representative range of these browsers and devices. Creating your service with web standards in mind will give it the best possible chance of working across all devices.

##Verified browsers

These are the browsers we recommend testing on when developing your service.  This list strongly recommends testing on a range of browsers created within the last 3 years that cover the largest representation of the user base.

The list is based upon usage statistics for GOV.UK and represents approximately 95% of the most popular browsers (the remaining browsers are individually insignificant).

Browsers not listed may still work well, and it should be noted that this is not a list that intends to suggest that these are the only browsers the service will work on -- this is simply a benchmark for testing against to ensure that the service will likely work for as many users as possible alongside appropriate cost-effectiveness and development overhead.

Services should ensure there's an obvious way for users to report problems they may find, and additional testing and adjustments should be made upon receiving such a report.

Two distinct levels of support are given and denoted next to each browser: 'compliant' or 'functional'. Compliant means that the service should look as good in this browser as in other modern browsers. Functional means that while it might not look perfect the service is still usable. In both cases the user should be able to access the information they need or complete their service.

Where 'latest version' is listed, it means the latest stable version plus one version back, as these browsers regularly update without intervention from the user.

### Desktop

<table>
  <thead>
    <tr>
      <th scope="col">Operating system</th><th scope="col">Browser</th><th scope="col">Support</th>
    </tr>
  </thead>
  <tr>
    <th scope="row" rowspan="4">Windows</th><td>Internet Explorer 8+</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Edge (latest version)</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Google Chrome (latest version)</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Mozilla Firefox (latest version)</td><td>Compliant</td>
  </tr>
  <tr>
    <th scope="row" rowspan="3">Mac OS X</th><td>Safari 8+</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Google Chrome (latest version)</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Mozilla Firefox (latest version)</td><td>Compliant</td>
  </tr>
</table>

### Small screen devices

<table>
  <thead>
    <tr>
      <th scope="col">Operating system</th><th scope="col">Version</th><th scope="col">Browser</th><th scope="col">Support</th>
    </tr>
  </thead>
  <tr>
    <th scope="row" rowspan="2">iOS</th><td rowspan="2">7+</td><td>Mobile Safari</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Google Chrome</td><td>Compliant</td>
  </tr>
  <tr>
    <th scope="row" rowspan="2">Android</th><td rowspan="2">4.x</td><td>Google Chrome</td><td>Compliant</td>
  </tr>
  <tr>
    <td>Android Browser</td><td>Compliant</td>
  </tr>
  <tr>
    <th scope="row">Windows Phone</th><td>8.1</td><td>Internet Explorer</td><td>Compliant</td>
  </tr>
</table>

##Developing universally accessible services

Digital by default services must take into consideration the limitations of the browsers people use to access them. One important idea for achieving this is [progressive enhancement](/service-manual/making-software/progressive-enhancement.html). This recognises that different bits of technology have different capabilities. While everybody gets access to core functionality, those using more sophisticated technology get an enhanced experience.

Progressive enhancement is also important in providing a consistent experience to people using mobile devices or those who may have limited bandwidth. Because mobile traffic now accounts for [30% of all internet use in the UK over the last year](http://gs.statcounter.com/#desktop+mobile-comparison-GB-monthly-201410-201509-bar 'Mobile vs Desktop in United Kingdom from October 2014 to September 2015 Statcounter Global Stats') and [40% of traffic to GOV.UK][pp-govuk-device-type], this mode of access is not an optional extra to consider.

Where we might previously have developed separate mobile and desktop versions of a service, or bought bespoke apps, design should now be done with one website in mind. This should be done using a [responsive design](https://en.wikipedia.org/wiki/Responsive_design 'Responsive Web Design -- Wikipedia') approach. This means websites adapt to suit the dimension of the screen being used to view it.

Don’t try to build services for every possible combination of operating system and browser. Avoid the temptation of designing for the obvious without first researching your users.

###Audience
Every service has an audience and you should investigate yours to see whether it has particular characteristics that you need to be aware of. Do you have existing data for the browsers and devices that your audience has been using already? If so, analyse it to see if you can identify any patterns in usage, or any combinations of:

* operating system
* browser
* browser version
* screen size/resolution
* mobile device

This data may sometimes support the case for deprioritising certain development work. Although most of GOV.UK is designed to work across all screen sizes, the [Trade Tariff](/trade-tariff) team chose not to tailor their tool to the smaller screen as it is largely used by office workers working during office hours. Equally, if your audience is likely to include those working in the public sector, there may be higher use of older, more limited browsers.

Channel shift means you must also consider your potential future audience.
[Device data for GOV.UK][pp-govuk-device-type] is displayed on the [Performance Platform][pp]
and this will be expanded to include operating system and browser data.

Before the launch of GOV.UK we noted a marked difference between the existing non-government and government audiences so you should also investigate the data provided by [NetMarketShare](http://www.netmarketshare.com 'NetMarketShare') and [GlobalStats](http://gs.statcounter.com 'Statcounter GlobalStats') who can provide UK and global trends.


##Continuous compatibility
It's important to distinguish between those browsers and operating systems whose popularity is either increasing or holding steady and those for which the opposite is true. Although Internet Explorer version 8 is only used by a minority (3.6% of visits to GOV.UK in the last year) this still accounts for a significant number of individuals that government services must take into account.

However, over time this will change. So it's important to set thresholds for abandoning support and for adopting new and emergent platforms. Windows 10 and Microsoft Edge highlight this dilemma -- recently launched products might not appear in any data but it is likely that they'll eventually enjoy widespread use.

Decisions about compatibility can't be something you specify at the start of the project and then forget about. Transformed digital services need to reflect and adapt to the broader internet context of their users on an ongoing basis.

##Further reading
James Weiner writes about the [decisions made about browser support](https://gds.blog.gov.uk/2012/01/25/support-for-browsers/ 'Support for browsers -- James Weiner, GDS') for the Beta of GOV.UK (January 2012).

Ben Welby discusses [the operating systems, browsers and devices supported](https://gds.blog.gov.uk/2012/10/11/what-devices-are-we-supporting-at-launch-and-why/ 'What devices are we supporting at launch, and why? -- Ben Welby, GDS') for the launch of GOV.UK (October 2012).

Tom Byers explores the practical ways in which GOV.UK has been [designed for different devices](https://gds.blog.gov.uk/2012/11/02/designing-for-different-devices/ 'Designing for different devices -- Tom Byers, GDS') (November 2012).

Dafydd Vaughan with an update on [browser usage on GOV.UK](https://gds.blog.gov.uk/2012/12/12/browser-usage-on-gov-uk/ 'Browser usage on GOV.UK -- Dafydd Vaughan, GDS') post-launch.

Edd Sowden shows some of the [interesting browser stats](https://gds.blog.gov.uk/2013/03/11/interesting-browser-stats/ 'Interesting browser stats') on www.gov.uk (March 2013).

The Guardian introduce their [use of responsive design](http://www.theguardian.com/help/developer-blog/2012/oct/18/responsive-design-guardian-introduction 'Responsive design at the Guardian: an introduction') (October 2012).

[pp-govuk-device-type]: /performance/site-activity/device-type
[pp]: /service-manual/measurement/performance-platform
