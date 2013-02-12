---
layout: gsdm
title: Configuring analytics tools
section: guidance
subsection: KPIs
status: draft
---
    
#Configuring analytics tools

Install and configure analytics tools that meet your needs. Where possible, use platforms that enable the data to be piped automatically into other systems. Using [APIs](http://en.wikipedia.org/wiki/Application_programming_interface) (Application Programming Interfaces) will stop you having to input data manually and allows for aggregation across multiple platforms.

##Guidance

* Have you installed web analytics software?
* Have you configured your web analytics software with the appropriate [conversion funnels](http://en.wikipedia.org/wiki/Conversion_funnel)?
* Do you have the capability to run user satisfaction surveys?
* Do you have the capability to do [A/B testing](http://en.wikipedia.org/wiki/Ab_testing) and [multivariate testing](http://en.wikipedia.org/wiki/Multivariate_testing)?

There are a number of open source products available as well as paid alternatives. If you are using a third party supplier, ensure that you have access to the raw data behind the measures specified in the contract and make sure data is not thrown away after a short period of time.

##Why we do this
![Page conversion funnel](https://assets.digital.cabinet-office.gov.uk/designprinciples/page_hits-41db2fb9ab658ba4ee0d577a3d847e78.png)
We are using Google Analytics to measure how users interact with GOV.UK pages. We want to know how far down the page they are reading so that we can tweak the content, if necessary.
This involved triggering events that can be picked up by Google Analytics at various points down the page: 25%, 50%, 75% and 100%.
The funnel visualisation shows the proportion of users who move on to the next page and the number who exit the site (ie the drop out rate).

!['Register to vote' conversion funnel](https://assets.digital.cabinet-office.gov.uk/designprinciples/funnel-conversion-5f179f569db3adde3c0cb02e58385cb5.png) We are developing a new service to allow people to join the Electoral Register online. The service is in [alpha](http://en.wikipedia.org/wiki/Software_release_life_cycle#Alpha) at the time of writing but we have already installed Google Analytics and configured it to measure how users flow through the transaction.
The funnel visualisations show at each stage the proportion of users proceeding through and the number who exit the process. This enables us to quickly spot where users are experiencing problems and where we might need to test alternative page designs.

##Further reading
The [Google Analytics Help Centre](http://support.google.com/analytics/?hl=en) is a useful resource if you use that particular platform, see for example [this guide](http://support.google.com/analytics/bin/answer.py?hl=en&answer=1012040) to setting up goals and funnels.

Blog article by Morgan Brown with a [good discussion on user flows and conversion funnels](http://uxdesign.smashingmagazine.com/2012/01/04/stop-designing-pages-start-designing-flows/).
