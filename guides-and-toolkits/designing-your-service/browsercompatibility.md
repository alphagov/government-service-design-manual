---
layout: gsdm
title: Browsers and devices
subtitle: Which ones to test with, and how best to support them
section: guidance
audience: 
  primary: designer, developer
type: resource
status: draft
---
    
#Designing for browsers and devices

Services should be universally accessible, regardless of how the user is choosing to access digital services.  
Due to the large range of browsers, devices and resolutions of access routes, it is to be expected and accepted that the experience will differ as technical capabilities and appropriate user experience on these browsers vary, as such, it is expected that you will verify that your service works across a representative range of these devices and browsers and accomodations for creating acceptable experiences across each should be made, as appropriate.

##Guidance

###Verified browsers 

These are the browsers we recommend testing on when developing your service.  This recommended testing list expects testing on a sample of browsers created within the last 3 years that cover the largest representation of the user base.  

This list is based upon usage statistics for www.gov.uk as a whole.  It allows for a 95% coverage of all browsers used (the remaining browsers are individually insignificant).  
Browsers not listed may still work well, and it should be noted that this is not a list that intends to suggest that these are the *only* browsers the service will work on - this is simply a benchmark for testing against to ensure that the service will likely work for as many users as possible along side appropriate cost-effectiveness and development overhead.  
Services should allow a clear way for users to report problems they may find, and additional testing and adjustments should be made upon receiving such a report.

*Note*: An exception is made for IE6, as this is still in large-scale use through-out government departments.

Two distinct levels of support are given and denoted next to each browser as F or C and are defined as:
* COMPLIENT
* FUNCTIONAL

Where "lastest version" is listed, it means the lastest stable version plus one version back, as these browsers regularly self-update.

####Desktop
* Windows 7
  * Internet Explorer 7 (C)
  * Internet Explorer 8 (C)
  * Google Chrome (latest version) (C)
  * Mozilla FireFox (latest version) (C)
  * Microsoft Internet Explorer 9
* Windows XP
  * Internet Explorer 6 (F)
* Windows Vista
* MacOS
  * Apple Safari
  * Google Chrome (latest version)
  * Mozilla FireFox (latest version)

####Mobile (and other smaller screen devices)
* iOS
  * Apple Safari
* Android 4.x

whilst ensuring it was at least functional on these operating systems:
* Android 2.3
* Blackberry 6.0 and above


###Developing services that are universally accessible

Digital by default services must take into consideration the limitations of the browsers people use to access them. One important idea for achieving this is [progressive enhancement](http://en.wikipedia.org/wiki/Progressive_enhancement 'Progressive enhancement - Wikipedia'). This recognises that different bits of technology have different capabilities. Whilst everybody gets access to core functionality those using more sophisticated technology get an enhanced experience.

Progressive enhancement is also important in delivering a consistent experience to people using mobile devices and who may have limited bandwidth. Because mobile traffic now accounts for [13% of all internet use in the UK](http://gs.statcounter.com/#mobile_vs_desktop-GB-monthly-201211-201211-bar 'Mobile vs Desktop in United Kingdom on November 2012 | Statcounter Global Stats') and around 20% of traffic to GOV.UK this mode of access is no longer an optional ‘extra’. Where we might previously have developed separate mobile and desktop versions of a service, or bought bespoke apps, we should now design with one website in mind. This should be done using a [responsive design](http://en.wikipedia.org/wiki/Responsive_design 'Responsive Web Design - Wikipedia') approach. This means websites adapt to suit the dimension of the screen being used to view it.

Don’t try to build services for every possible combination of operating system and browser and avoid the temptation of designing for the obvious without first researching your audience.

###Audience
Every service has an audience and you should investigate yours to see whether it has particular characteristics that you need to be aware of. Do you have existing data for the browsers and devices that your audience has been using already? If so, analyse it to see if you can identify any patterns in usage, or any combinations of:
* operating system  
* browser
* browser version
* screen size/resolution
* mobile device

This data may sometimes support the case for deprioritising certain development work: although most of GOV.UK is designed to work across all screen sizes the Trade Tariff team chose not to tailor their tool to the smaller screen given that it is used by office workers between 9 and 5. Equally, if your audience is likely to include those from within the public sector there may be higher use of older, more limited browsers.

In addition channel shift means you must also consider your potential audience. It is anticipated that operating system, browser and device data from GOV.UK will be published as part of the GDS performance platform and this will provide a valuable insight into the audience for government services. Before launch we noted a marked difference between the existing non-government and government audiences so you should also investigate the data provided by [NetMarketShare](http://www.netmarketshare.com 'NetMarketShare') and [GlobalStats](http://gs.statcounter.com 'Statcounter GlobalStats') who can provide UK and global trends.


###Continuous compatibility
We need to distinguish between those browsers and operating systems whose popularity is either increasing or holding steady and those for which the opposite is true. Although Internet Explorer versions 6 and 7 have only been used by a minority (almost 5% of the total visits to GOV.UK since launch) this still accounts for a significant number of individuals who government services must take into account. 

This will change over time and so it is important to set thresholds for abandoning support and for adopting new and emergent platforms. The iPad Mini, Kindle Fire, Windows 8 and Internet Explorer 10 highlight this dilemma - they are recently launched products so do not appear in any data but it is likely that they will enjoy widespread mainstream usage. 

Decisions about compatibility can not be something you specify at the start of the project and then forget about. Transformed digital services need to reflect the broader internet context of the country on an ongoing basis.

##Further reading
James Weiner writes about the [decisions made about browser support](http://digital.cabinetoffice.gov.uk/2012/01/25/support-for-browsers/ 'Support for browsers - James Weiner, GDS') for the Beta of GOV.UK (January 2012)

Ben Welby discusses [the operating systems, browsers and devices supported](http://digital.cabinetoffice.gov.uk/2012/10/11/what-devices-are-we-supporting-at-launch-and-why/ 'What devices are we supporting at launch, and why? - Ben Welby, GDS') for the launch of GOV.UK (October 2012).

Tom Byers explores the practical ways in which GOV.UK has been [designed for different devices](http://digital.cabinetoffice.gov.uk/2012/11/02/designing-for-different-devices/ 'Designing for different devices - Tom Byers, GDS') (November 2012).

Dafydd Vaughan with an update on [browser usage on GOV.UK](http://digital.cabinetoffice.gov.uk/2012/12/12/browser-usage-on-gov-uk/ ‘Browser usage on GOV.UK - Dafydd Vaughan, GDS’) post-launch.

The Guardian introduce their [use of responsive design](http://www.guardian.co.uk/help/developer-blog/2012/oct/18/responsive-design-guardian-introduction 'Responsive design at the Guardian: an introduction') (October 2012).

Helpful summary of [progressive enhancement](http://www.alistapart.com/articles/understandingprogressiveenhancement/ 'Understanding progressive enhancement') (October 2008).

