---
layout: gsdm
title: Designing for browsers and devices
section: guidance
subsection: Designing your service
status: draft
---
    
#Designing for browsers and devices

Services must be accessible wherever people are and whatever technology they’re using by default; not as an optional extra. The user experience can vary significantly from operating system to operating system, browser version to browser version and desktop to mobile. However, don’t try to build services for every possible product but use data to inform and set your priorities for development.

##Guidance

###Verified browsers 

These are the browsers we recommend testing on when developing your service.  

This list is based upon usage statistics for www.gov.uk as a whole, and considers any browser with a usage by 2%+ of users to be significant. 

An exception is made for IE6, as this is still in large-scale use through-out government departments.

Two distinct levels of support are given and denoted next to each browser as F or C and are defined as:
* COMPLIENT
* FUNCTIONAL

Where "lastest version" is listed, it means the lastest stable version plus one version back, as these browsers regularly self-update.

####Desktop
* Windows 7
  * Microsoft Internet Explorer 7 (C)
  * Microsoft Internet Explorer 8 (C)
* Windows XP
  *Internet Explorer 6 (F)
* Windows Vista
* MacOS
* iOS
* Android 4.x

these browsers:
* Apple Safari (latest version)
* Google Chrome (latest version)
* Microsoft Internet Explorer 8
* Microsoft Internet Explorer 9
* Mozilla Firefox (versions above 5)

####Mobile


whilst ensuring it was at least functional on these operating systems:
* Android 2.3
* Blackberry 6.0 and above

and these browsers:
* Microsoft Internet Explorer 7
* Microsoft Internet Explorer 6
* 
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

###Platforms
There are two big vendors of operating system: Apple and Microsoft and they each have several versions of their products which each support a unique selection of browsers. Apple's Safari, Google's Chrome, Microsoft's Internet Explorer and Mozilla's Firefox are the most popular brands. The latest versions of these all 'silently update' and are therefore always current. There are currently 4 older versions of Internet Explorer still being used enough to need attention. Each of these have statistically significant usage and varying levels of support for web standards (code written to [W3C specifications](http://www.w3.org/standards/ 'Standards - W3C')).

The device market has similar complexity. Masked by the big 4 platforms, Apple's iOS, Google's Android, Research in Motion's Blackberry and Microsoft's Windows, is a complex relationship between device manufacturers, operating system versions and mobile browsers with 2012 estimates putting the number of [Android models in circulation at 4,000](http://opensignalmaps.com/reports/fragmentation.php 'Android device fragmentation - OpenSignalMaps').

In the first month of launch GOV.UK was accessed by 2,963 different browser versions and 900 separate devices. Recognising that it is impossible to build services that have been tested across every possible combination of browser and operating system we used data from DirectGov, BusinessLink and some departments, wider data from external sources and the opinions of GDS developers to identify our priorities. This produced a two tier approach: we want the site to work well in the most common situations, and to be functional in other high profile but less popular circumstances (functional means all content being viewable (no overlapping areas) and any interactive elements not preventing access to content).



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

