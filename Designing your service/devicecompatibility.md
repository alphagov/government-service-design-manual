#Device compatibility
Services must be accessible across devices by default; not as an optional extra. However, you should not try to build services for every possible device and mobile operating system but use data to inform that decision making process.

##Guidance/Tool
Mobile traffic accounts for [13% of all internet use in the UK] (http://gs.statcounter.com/#mobile_vs_desktop-GB-monthly-201211-201211-bar 'Mobile vs Desktop in United Kingdom on November 2012 | Statcounter Global Stats') and around 20% of traffic to GOV.UK. Where we might previously have developed separate mobile and desktop versions of a service or bought bespoke apps we should now use [responsive design] (http://en.wikipedia.org/wiki/Responsive_design 'Responsive Web Design - Wikipedia') and [progressive enhancement] (http://en.wikipedia.org/wiki/Progressive_enhancement 'Progressive enhancement - Wikipedia') to provide a consistently usable experience for the public.

The responsive design approach builds websites that adapt to suit the dimension of the screen being used to view it. In this way the same platform can be used to give visitors an experience tailored to their device whether it's a television, monitor, tablet, or phone.

Progressive enhancement recognises that different bits of technology have different capabilities and that data usage is an important consideration for those accessing services on the move. It ensures that core functionality is available for everybody while giving those using more sophisticated technology or access to greater bandwidth an enhanced experience.

###Audience
Every service has an audience and you should investigate yours to see whether it has particular characteristics that you need to be aware of. This data may make the case for deprioritising device based access: although most of GOV.UK is designed to work across all screen sizes the Trade Tariff team chose not to tailor their tool to the smaller screen given that it is used by office workers between 9 and 5.

However, with channel shift in mind it is important to consider your potential audience. It is anticipated that device data from GOV.UK will be published as part of the GDS performance platform and this will provide a valuable insight into the audience for government services. Before launch we noted a marked difference between the existing non-government and government audiences so you should also investigate the data provided by [NetMarketShare] (http://www.netmarketshare.com 'NetMarketShare') and [GlobalStats] (http://gs.statcounter.com 'Statcounter GlobalStats') who can provide UK and global trends.

###Platforms
There are 4 platforms in the device market: Apple's iOS, Google's Android, Research in Motion's Blackberry and Microsoft's Windows. However, this apparent simplicity masks the complex relationship between device manufacturers, operating system versions and mobile browsers. In the first month of launch GOV.UK was accessed by 900 separate devices and 2012 estimates have put the number of [Android models in circulation at 4,000] (http://opensignalmaps.com/reports/fragmentation.php 'Android device fragmentation - OpenSignalMaps').

Based on government web data (usage from DirectGov, BusinessLink and some departments), wider data from external sources and the personal input of GDS developers we were able to identify the most likely screen sizes, form factors and operating system versions. This allowed us to take a two tier approach: stating that we wanted the site to work well on the most commonly used devices, and to be functional on others. 

We need to distinguish between those devices whose popularity is either increasing or holding steady and those for which the opposite is true. Because this is fluid it is important to set clear thresholds for retiring support or adopting new and emergent devices. Device compatibility is not something that can be done at the start of the project and forgotten about. Transformed digital services need to reflect what people have in their hands on an ongoing basis.

##Code/Templates
If you're giving people code r copy to cuta and paste then here is where it will go.

##Why we do this
Tom Byers from GDS explores the practical ways in which GOV.UK has been [designed for different devices] (http://digital.cabinetoffice.gov.uk/2012/11/02/designing-for-different-devices/ 'Designing for different devices - Tom Byers, GDS') (November 2012)

The Guardian introduce their [use of responsive design] (http://www.guardian.co.uk/help/developer-blog/2012/oct/18/responsive-design-guardian-introduction 'Responsive design at the Guardian: an introduction') (October 2012)

Helpful summary of [progressive enhancement] (http://www.alistapart.com/articles/understandingprogressiveenhancement/ 'Understanding progressive enhancement') (October 2008)

##Further reading
[Browser compatibility] (http://gsdm.herokuapp.com/handbook/38/ 'Browser compatibility')