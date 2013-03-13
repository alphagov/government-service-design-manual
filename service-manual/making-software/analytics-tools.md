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
* who owns the data
* the cost of additional profiles and/or custom variables
* the admin system users have access to
* whether it is hosted by the vendor or in-house
* whether it tracks offline channel usage
* whether it provides a comprehensive set of standard reports (including social interactions and multimedia capturing)
* whether it provides an open API with no export restrictions
* if it can measure transactions through funnel analysis and measure goals
* support and training
* the cookies it requires

For each of these criteria, you should identify which are fulfilled as part of the standard quoted package and what is charged for any additional features.

## Privacy

The privacy and security of data is of the utmost importance. Make sure your analytics solution and processes take the following into account:

* do not collect and process any personal information (the terms and conditions of your analytics provider will probably expressly forbid you from doing this)
* turn off any data sharing (in Google Analytics this is done by Admin -> <Account name> -> Account Settings)
* anonymise IP addresses that your analytics provider collects, by removing the last octet of the address
* obtain your vendor’s privacy and security policy
* you should own and be able to export analytics data (Note that Google owns data in the free Google Analytics product)
* does the solution meet the EU privacy directive and the European Commission’s Directive on Data Protection?
* where is collected data held?
* do data centres meet EU/British data security standards?
* how long is data held for?
* what will happen to the data on termination of the contract- can you export it?
* data aggregation and sharing (does the vendor mine your data for cross-customer benchmarking/trends or to provide usage data to any advertising channels)
* what access your vendors employees have to your data (make sure there are adequate administration tools to control appropriate access for your own staff)

## Vendor comparison

The following table assesses four major analytics tools against the criteria set out above. This information was compiled from Econsultancy’s [Web Analytics Buyer’s Guide 2012](http://econsultancy.com/uk/reports/web-analytics-buyers-guide) and input from GDS and other government departments. There are numerous other vendors in the market place, including open source solutions such as [Piwik](http://piwik.org/), which may also be worth considering.

<table>
<tr><th></th><th>Google Analytics Premium</th><th>Adobe Omniture</th><th>WebTrends</th><th>Digital Analytics (Comscore)</th></tr>
<tr><td>Cost</td><td>£90k per year</td><td>Based on onsite actions and page views</td><td>Charged per server call at low level at higher levels at volumetrics. Will charge mixed fee for the contract. If exceed will review charging model.

SDK mobile is  extra for Apps; but mobile site use is no extra.</td><td>Based on number of monthly server calls, type of measurements, integration of external marketing tools, size and type of user base including concurrent users. Prices quoted £200k year 1 and £400k year 2 just for licensing</td></tr>
<tr><td>Data ownership</td><td>Yes</td><td>Unsure</td><td>Unsure</td><td>Unsure</td></tr>
<tr><td>1st party cookies only</td><td>Yes</td><td>Yes</td><td>Yes</td><td>Yes</td></tr>
<tr><td>Profiles and custom variables</td><td>Unlimited profiles and 50 custom variables</td><td>Yes but possible extra cost</td><td>Unsure</td><td>Unsure</td></tr>
<tr><td>Admin system</td><td>Yes</td><td>Yes</td><td>Yes</td><td>Yes</td></tr>
<tr><td>Vendor hosted solution</td><td>Yes</td><td>Yes</td><td>Yes</td><td>Yes</td></tr>
<tr><td>Multi-channel</td><td>Multi channel funnels and attribution analysis available</td><td>Yes</td><td>Yes but not necessarily part of standard package</td><td>Yes but may incur extra costs</td></tr>
<tr><td>Standard reports</td><td>Yes</td><td>Yes</td><td>Yes</td><td>Yes</td></tr>
<tr><td>Open API</td><td>Yes</td><td>Reporting API rather than a data API</td><td>Yes</td><td>Yes</td></tr>
<tr><td>Funnel analysis</td><td>Yes</td><td>Yes</td><td>Yes</td><td>Yes</td></tr>
<tr><td>Support and training</td><td>Yes and implementation at no extra charge</td><td>Yes but at extra cost</td><td>Yes but at extra cost</td><td>Yes but at extra cost</td></tr>
<tr><td>Expert review</td><td>If you are running a lot of PPC, media campaigns, marketing then Google Analytics Premium is a good option as this is a fixed cost regardless of the volume of visits you have to the site.</td><td>Adobe Omniture is the biggest one out there. However, there are downsides, this is expensive and any additional functionality you require comes at a cost from different modules.</td><td>Webtrends is great to keep an accurate (or as close as it can be) record of traffic to the site. and this records it on a session by session basis. and whilst not out of the box you can create any number of scenario and funnel reports for transactional activities.</td><td>Comscore is a great all-rounder and is an easy install across multiple pages. This measures every activity and for a very busy site can get very expensive as they charge per call.</td></tr>
</table>

## Configuring analytics tools

Install and configure analytics tools that meet your needs. Where possible, use platforms that enable the data to be piped automatically into other systems. 

Using [APIs](http://en.wikipedia.org/wiki/Application_programming_interface) (Application Programming Interfaces) will stop you having to input data manually and allows for aggregation across multiple platforms.

You will need to answer the following:

* have you installed web analytics software?
* have you configured your web analytics software with the appropriate [conversion funnels](http://en.wikipedia.org/wiki/Conversion_funnel)?
* do you have the capability to run user satisfaction surveys?
* do you have the capability to do [A/B testing](http://en.wikipedia.org/wiki/Ab_testing) and [multivariate testing](http://en.wikipedia.org/wiki/Multivariate_testing)?

## Examples
![Page conversion funnel](https://assets.digital.cabinet-office.gov.uk/designprinciples/page_hits-41db2fb9ab658ba4ee0d577a3d847e78.png)

We are using Google Analytics to measure how users interact with GOV.UK pages. We want to know how far down the page they are reading so that we can tweak the content, if necessary.

This involved triggering events that can be picked up by Google Analytics at various points down the page: 25%, 50%, 75% and 100%.

The funnel visualisation shows the proportion of users who move on to the next page and the number who exit the site (i.e. the drop out rate).

!['Register to vote' conversion funnel](https://assets.digital.cabinet-office.gov.uk/designprinciples/funnel-conversion-5f179f569db3adde3c0cb02e58385cb5.png) 

We are developing a new service to allow people to join the Electoral Register online. The service is in [alpha](http://en.wikipedia.org/wiki/Software_release_life_cycle#Alpha) at the time of writing but we have already installed Google Analytics and configured it to measure how users flow through the transaction.

The funnel visualisations show at each stage the proportion of users proceeding through and the number who exit the process. This enables us to quickly spot where users are experiencing problems and where we might need to test alternative page designs.

## Further reading
The [Google Analytics Help Centre](http://support.google.com/analytics/?hl=en) is a useful resource if you use that particular platform, see for example [this guide](http://support.google.com/analytics/bin/answer.py?hl=en&answer=1012040) to setting up goals and funnels.

Blog article by Morgan Brown with a [good discussion on user flows and conversion funnels](http://uxdesign.smashingmagazine.com/2012/01/04/stop-designing-pages-start-designing-flows/).
