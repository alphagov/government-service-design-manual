---
layout: detailed-guidance
title: Analytics tools
subtitle: Choosing the right tools for your service
category: making-software
type: guide
audience:
  primary: performance-analysts 
  secondary: service-managers, qa
phases:
  - beta
  - live
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
---

There are various web analytics tools available to help you measure how people are using your service. Be sure to assess how well a particular tool meets your needs before deciding on which tools to use.

## High-level requirements

It’s essential that your chosen analytics tools provide an open [Application Programming Interface](/service-manual/making-software/apis.html) (API) with no restrictions on exporting data.

Also, when deciding which analytics tool is most appropriate for your service, consider:


* the total cost of ownership, as well as the cost in comparison to turnover of service
* the volume of data being sampled
* who owns the data (it should be your organisation!)
* the cost of additional profiles and/or custom variables
* the admin system users have access to
* whether it is hosted by the vendor or in-house
* whether it tracks offline channel usage
* whether it provides a comprehensive set of standard reports (including social interactions and multimedia capturing)
* if it can measure transactions through funnel analysis and measure goals
* support and training
* the cookies it requires

For each of these criteria, identify which ones are fulfilled as part of the standard quoted package and what is charged for any additional features.

## Privacy

The privacy and security of data is of absolute importance, so review your vendor’s privacy and security policy.

Make sure your analytics solution and processes:

* won't collect and process any personal information (the terms and conditions of your analytics provider will probably expressly forbid you from doing this)
* turn off any data sharing (some suppliers may collect data anonymously for internal benchmarking)
* anonymise IP addresses that your analytics provider collects by removing the last octet of the address
* meets the EU privacy directive and the European Commission’s Directive on Data Protection
* has data centres that meet EU/British data security standards
* allows you to restrict your vendor’s employees access to your data as appropriate -- also make sure there are adequate administration tools to control appropriate access for your own staff
* allows you to own export analytics data (pay attention to terms and conditions for any free products)

As part of keeping your users' data secure, find out:

* where collected data is held
* how long data is held for?
* what happens to the data on termination of the contract - can you export it?
* if your vendor mines data (analyses data from different perspectives, turning it into useful summaries of information) for cross-customer benchmarking/trends or to provide usage data to any advertising channels

## Vendor comparison

There are many digital/web analytics vendors in the marketplace, along with open source solutions.

A [search for 'analytics tools comparison'](https://www.bing.com/search?q=analytics+tools+comparison) provides a number of useful resources where you can compare the capabilities and strengths of different services. [See some more examples](#further-reading).

## Configuring analytics tools

Install and configure analytics tools that meet your needs. Where possible, use platforms that enable the data to be piped automatically into other systems.

Use [APIs](https://en.wikipedia.org/wiki/Application_programming_interface) (Application Programming Interfaces) as they will stop you from having to input data manually and allow the grouping of data across multiple platforms.

You will need to answer the following:

* have you installed web analytics software?
* have you configured your web analytics software with the appropriate [conversion funnels](https://en.wikipedia.org/wiki/Conversion_funnel)?
* do you have the capability to run user satisfaction surveys?
* do you have the capability to do [A/B testing](https://en.wikipedia.org/wiki/Ab_testing) and [multivariate testing](https://en.wikipedia.org/wiki/Multivariate_testing)?

## Further reading

Sites that provide vendor comparison information:

* [AboutAnalytics](http://www.aboutanalytics.com/)
* [Web Analytics Tools Comparison: A Recommendation](http://www.kaushik.net/avinash/web-analytics-tools-comparison-a-recommendation/) - old, but thoughtful
* [Analytics tool comparison](http://www.slideshare.net/shvmdhwn/analytics-tool-comparison)
* [Enterprise Web Analytics: A Buyer’s Guide](http://marketingland.com/buyers-guides/enterprise-web-analytics-tools-in-the-facebook-era-a-buyers-guide)

Sites that provide information on digital analytics:

* [Occam’s Razor by Avinash Kaushik](http://www.kaushik.net/avinash/) has a wealth of useful information on analytics that's easy to understand
* [Web Analytics Demystified](http://www.webanalyticsdemystified.com/) and the [Digital Analytics Association](http://www.digitalanalyticsassociation.org/) have free whitepapers

Sites that provide information on funnels (important for transactions):

* [Conversion Funnels](http://wiki.clicktale.com/Article/Conversion_Funnels)
* [What is a Conversion Funnel?](http://www.webics.com.au/blog/conversion-tracking/conversion-funnel/)
* [Blog post by Morgan Brown with a good discussion on user flows and conversion funnels](http://www.smashingmagazine.com/2012/01/04/stop-designing-pages-start-designing-flows/)

