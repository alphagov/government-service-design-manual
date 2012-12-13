#Browser compatibility
The design of a website can vary significantly from operating system to operating system, browser to browser and even version to version. People can update their systems infrequently and as a result digital services need to think about supporting the older, more problematic browsers as well as taking advantage of emerging technology and web services.


##Guidance/Tool
Digital by default services must take into consideration the known problems associated with certain browsers by adopting the idea of [progressive enhancement] (http://en.wikipedia.org/wiki/Progressive_enhancement 'Progressive enhancement - Wikipedia'). This approach recognises that different bits of technology have different capabilities. Everybody gets access to core functionality but those using more sophisticated technology get an enhanced experience.

You should not try to build services for every possible combination of operating system and browser. Equally, you should avoid the temptation to design for the obvious without first researching your audience.

###Audience
Every service has an audience and you should investigate yours to see whether it has particular characteristics that you need to be aware of. Do you have existing data for the browsers and devices that your audience has been using already? If so, analyse it to see if you can identify any patterns in usage, or any combinations of:
* operating system  
* browser
* browser version
* screen size/resolution

However, with channel shift in mind it is important to consider your potential audience. It is anticipated that device data from GOV.UK will be published as part of the GDS performance platform and this will provide a valuable insight into the audience for government services. Before launch we noted a marked difference between the existing non-government and government audiences so you should also investigate the data provided by [NetMarketShare] (http://www.netmarketshare.com 'NetMarketShare') and [GlobalStats] (http://gs.statcounter.com 'Statcounter GlobalStats') who can provide UK and global trends.

###Operating systems and browsers
There are two big vendors of operating system: Apple and Microsoft and they each have several versions of their products which each support a unique selection of browsers. Apple's Safari, Google's Chrome, Microsoft's Internet Explorer and Mozilla's Firefox are the most popular brands but whilst some 'silently update' and are therefore always current, Internet Explorer has four versions, each of which has statistically significant usage and each of which interprets web design differently. This presents a challenge to those designing services about where to focus development resource and at what level of usage should something be made a priority.

In the first month of launch GOV.UK was accessed by 2,963 different browser versions. Recognising that it is impossible to build services that have been tested across every possible combination of browser and operating system. GDS took a two tier approach: stating that we wanted the site to work well on the most commonly used browsers, and to be functional on others.

###GOV.UK
Based on our known audience (using data from DirectGov, BusinessLink and some departments) and wider data from external sources we decided to make sure GOV.UK worked well on the following operating systems:
* Windows 7
* Windows XP
* Windows Vista
* MacOS
* iOS
* Android 4.x

and these browsers:
* Apple Safari (latest version)
* Google Chrome (latest version)
* Microsoft Internet Explorer 8
* Microsoft Internet Explorer 9
* Mozilla Firefox (versions above 5)

whilst ensuring it was at least functional on these operating systems:
* Android 2.3
* Blackberry 6.0

and these browsers:
* Microsoft Internet Explorer 7
* Microsoft Internet Explorer 6

###Continuous compatibility
We need to distinguish between those browsers and operating systems whose popularity is either increasing or holding steady and those for which the opposite is true. Although Internet Explorer versions 6 and 7 have only been used by a minority (almost 5% of the total visits to GOV.UK since launch) this still accounts for a significant number of individuals who government services must take into account. 

This will change over time and so it is important to set thresholds for abandoning support and for adopting new and emergent platforms. Windows 8 and Internet Explorer 10 highlight this dilemma - this is a recently launched product so does not appear in any data but it is likely to enjoy widespread mainstream usage. Therefore decisions about compatibility can not be something you specify at the start of the project and then forget about. Transformed digital services need to reflect the broader internet context of the country on an ongoing basis.

##Code/Templates
If you're giving people code r copy to cuta and paste then here is where it will go.

##Why we do this
James Weiner writes about the [decisions made about browser support] (http://digital.cabinetoffice.gov.uk/2012/01/25/support-for-browsers/ 'Support for browsers') for the Beta of GOV.UK (January 2012)

Ben Welby follows that up with a discussion of [the operating systems, browsers and devices supported] (http://digital.cabinetoffice.gov.uk/2012/10/11/what-devices-are-we-supporting-at-launch-and-why/ 'What devices are we supporting at launch, and why?') for the launch of GOV.UK (October 2012).

##Further reading
[Device compatibility] (http://gsdm.herokuapp.com/handbook/44/ 'Device compatibility')