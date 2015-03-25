---
layout: detailed-guidance
title: Using data
subtitle: How to make use of the performance information your service collects
status: draft
category: measurement
type: guide
audience:
  primary: service-managers, performance-analysts
  secondary: designers, qa, content-designers
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Measurement
    url: /service-manual/measurement
---

Simply collecting information about how your service is running isn’t enough to make judgements about how to improve it. A process of continual iteration and close measurement will help you to see what needs improvement and how to investigate ways of improving your service.

![Diagram showing seven stage iterative process](/service-manual/assets/images/kpis/monitor.png)

##Understand user needs/decide what to measure

Make sure the [core key performance indicators (KPIs)](/service-manual/measurement#what-will-you-measure) established in the [service standard](/service-manual/digital-by-default/index.html), and any [other KPIs](/service-manual/measurement/other-kpis.html) you choose to measure, accurately reflect the needs of your users and stakeholders.

This will allow you to measure the ability of your service to meet user needs.

##Install and configure platforms
While building your service, make sure that appropriate analytics tools are being used to monitor the service, collecting the data necessary to produce accurate and timely measurements.

For more information, read our guidance on [choosing analytics tools](/service-manual/making-software/analytics-tools.html).

##Establish a baseline

Establish a ‘baseline’ based on the current performance trends of each channel. Use this to judge the effectiveness of any changes to your service. This will help you pinpoint the effectiveness of your initiatives, and identify what worked.

Look at performance trends over time, rather than taking a snapshot at a particular point. Peaks and dips in performance can then be measured in relation to this base (or trend) line, which helps to:

* identify the effect of your communications or design initiatives
* reveal seasonal variations in performance

##Aggregate data

Collect and combine performance information from multiple sources and across multiple channels. Make sure you understand what this will mean in terms of system requirements.

Combining data often reveals useful insights, eg into service efficiency (cost per transaction and total cost to serve) or the amount of usage by channel (percentage digital take-up vs post, phone etc).

##Analyse and visualise data

Communicate performance information to your users through the appropriate:

* dashboards
* reports
* alerts

Highlight specific segments that you know users are interested in, and make sure that your visualisations are simple, useful and contain minimal amounts of [chartjunk](https://en.wikipedia.org/wiki/Chartjunk).

Typical segments include:

* channel used to access service -- through which channel(s) did the user find out about and attempt to use the service?
* new compared with repeat visitors -- are first time users behaving differently to those who have used the service before?
* geographical region -- how popular is the digital service by region and how does that compare with online penetration in general?
* product type -- does the user experience vary depending on the type of product or service?
* value -- is performance dependent on the monetary value of the product or service being sought?
* dashboards -- based on objectives and will help inform decisions, often with the help of real-time data
* reports -- providing regular, scheduled snapshots of data, but tend to need extra context and time to absorb
* alerts -- used to inform the users about a change or an event, often using attention-grabbing tools to do so

By making your visualisations easy to understand you significantly increase the likelihood that the information will be used to improve your service.

Best practices include:

* keeping charts plain -- don’t use shading, 3D or other distracting effects
* removing clutter -- don’t use trend lines, grid lines, unnecessary labelling
* not using pie charts -- they require more space and are harder to read than bar charts
* using text where a chart adds nothing

##Monitor, iterate, and improve

1.    Test a range of actions for improving performance and monitor to see which work well. Pilot these on a portion of your users to minimise risk.
2.    Implement the best performing solutions widely and then repeat this process continuously as what you measure will change over the course of your product or project’s lifetime.

Any service that meets user needs will include an element of user feedback. Monitor this and act on it. Doing so will help to continually improve the service for your users.

Many options are available for improving the overall performance of your service. The following examples are based on the [4 Ps of marketing](https://en.wikipedia.org/wiki/Marketing_mix):

* price -- change the price eg to attract people to the digital channel
* product -- improve the user experience (eg from user feedback, user testing, A/B testing, multivariate testing)
* placement -- place the digital service URL on printed materials and in voice recordings
* promotion -- increase your use of email and social media to encourage repeated use of your digital service

Taking an iterative approach to service development will increase the speed of improvements and minimises the risk of failure. Don’t wait until the end to do this -- do it continuously throughout the process.

##Further reading

To find out more about ways to use data, try reading:

* an article on [online customer satisfaction scores](http://www.theguardian.com/money/2010/dec/22/amazon-top-consumer-satisfaction) for retailers, based on the Customer Satisfaction Index, written for *The Guardian*
* [The role of the data scientist](http://radar.oreilly.com/2010/06/what-is-data-science.html) by Mike Loukides
* [Costing customer time](/government/publications/costing-customer-time), a method for calculating the cost of a transaction for both the service provider and the user (HMRC developed a method for calculating the cost of a users time when interacting with government -- this is important because some channels may be quicker to use than others)
* [Designing with Data](http://designingwithdata.co.uk/), an excellent book by Brian Suda which helps you to design beautiful and powerful data visualisations
* [Juice Analytics](http://www.juiceanalytics.com/), a website with lots of useful resources on how to design and develop useful data visualisations and dashboards
* Edward Tufte’s [The Visual Display of Quantitative Information](http://www.edwardtufte.com/tufte/books_vdqi), an original piece of work on data visualisation and introduces the concept of chartjunk
* the [Flowing Data](http://flowingdata.com/) blog by Nathan Yau, a useful source of data visualisation news
* the [D3 Gallery](https://github.com/mbostock/d3/wiki/Gallery), a stunning collection of data visualisations
* a [nicely presented overview](http://selection.datavisualization.ch/) of some of the tools available for data visualisation
* this [article in Wired](http://www.wired.com/2012/04/ff_abtesting/all/1), which shows how A/B testing was used to good effect in Obama’s election campaign
* this [article in eConsultancy](https://econsultancy.com/blog/2454-q-a-lovefilm-s-craig-sullivan-on-a-b-and-multi-variate-testing), which shows how multivariate testing was used to improve conversion rates at Lovefilm

*[KPIs]: key performance indicators
