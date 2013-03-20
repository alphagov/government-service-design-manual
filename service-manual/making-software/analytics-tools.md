---
layout: detailed-guidance
title: Analytics tools
subtitle: Choosing the right ones for your service
category: making-software
type: guide
audience:
  primary: performance-analysts 
  secondary: service-managers, tech-archs, qa
phases:
  - beta
  - live
audience: 
    primary: analyst
    secondary: service-managers
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
---

There are various web analytics tools available to help you measure how people are using your service. You will need to assess how well a particular tool meets your needs before deciding which one to use.

This guidance describes some of the criteria you should consider and reviews some of the main analytics tools against them.

## High level requirements

When deciding which analytics tool is most appropriate for your service, you should consider the following:


* the total cost of ownership as well as cost in comparison to turnover of service
* the volume of data being sampled
* who owns the data (it should be your organisation!)
* the cost of additional profiles and/or custom variables
* the admin system users have access to
* whether it is hosted by the vendor or in-house
* whether it tracks offline channel usage
* whether it provides a comprehensive set of standard reports (including social interactions and multimedia capturing)
* it **must** provide an open API with no export restrictions
* if it can measure transactions through funnel analysis and measure goals
* support and training
* the cookies it requires

For each of these criteria, you should identify which are fulfilled as part of the standard quoted package and what is charged for any additional features.

## Privacy

The privacy and security of data is of the utmost importance. Make sure your analytics solution and processes take the following into account:

* do not collect and process any personal information (the terms and conditions of your analytics provider will probably expressly forbid you from doing this)
* turn off any data sharing (some suppliers may collect data anonymously more internal benchmarking)
* anonymise IP addresses that your analytics provider collects, by removing the last octet of the address
* obtain and review your vendor’s privacy and security policy
* you should own and be able to export analytics data (pay attention to terms and conditions for any free products)
* does the solution meet the EU privacy directive and the European Commission’s Directive on Data Protection?
* where is collected data held?
* do data centres meet EU/British data security standards?
* how long is data held for?
* what will happen to the data on termination of the contract- can you export it?
* data aggregation and sharing (does the vendor mine your data for cross-customer benchmarking/trends or to provide usage data to any advertising channels)
* what access your vendors employees have to your data (make sure there are adequate administration tools to control appropriate access for your own staff)

## Vendor comparison

There are a range of digital/web analytics vendors in the market-place, together with open source solutions.

A search for [Analytics Tools Comparison](http://www.bing.com/search?q=Analytics+Tools+Comparison&FORM=QSRE4) provides a number of useful resources where you can compare the capabilities and strengths of different services. (See some examples in [Further reading](#further-reading) below).

## Configuring analytics tools

Install and configure analytics tools that meet your needs. Where possible, use platforms that enable the data to be piped automatically into other systems.
Using [APIs](http://en.wikipedia.org/wiki/Application_programming_interface) (Application Programming Interfaces) will stop you having to input data manually and allows for aggregation across multiple platforms.
You will need to answer the following:

* have you installed web analytics software?
* have you configured your web analytics software with the appropriate [conversion funnels](http://en.wikipedia.org/wiki/Conversion_funnel)?
* do you have the capability to run user satisfaction surveys?
* do you have the capability to do [A/B testing](http://en.wikipedia.org/wiki/Ab_testing) and [multivariate testing](http://en.wikipedia.org/wiki/Multivariate_testing)?

## Further reading

A selection of sites that provide vendor comparison information:

* [About analytics](http://www.aboutanalytics.com/)
* [Web Analytics Tools Comparison: A Recommendation](http://www.kaushik.net/avinash/web-analytics-tools-comparison-a-recommendation/) Old, but thoughtful
* [Analytics tool comparison](http://www.slideshare.net/shvmdhwn/analytics-tool-comparison)
* [Enterprise Web Analytics: A Buyer’s Guide](http://searchengineland.com/buyers-guides/enterprise-web-analytics-tools-in-the-facebook-era-a-buyers-guide)

For for information on digital analytics:

* [Occam’s Razor by Avinash Kaushik](http://www.kaushik.net/avinash/) has a wealth of useful, easily digestible information on analytics
* [Web Analytics Demystified](http://www.webanalyticsdemystified.com/) and the [Digital Analytics Association](http://www.digitalanalyticsassociation.org/) have free whitepapers.

For transactions, it’s important to understand funnels:

* [Conversion Funnels](http://wiki.clicktale.com/Article/Conversion_Funnels)
* [What is a Conversion Funnel?](http://www.webics.com.au/blog/conversion-tracking/conversion-funnel/)
* Blog article by Morgan Brown with a [good discussion on user flows and conversion funnels](http://uxdesign.smashingmagazine.com/2012/01/04/stop-designing-pages-start-designing-flows/)

