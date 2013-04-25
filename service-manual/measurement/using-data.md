---
layout: detailed-guidance
title: Using data
subtitle: How to make use of the performance information your service collects
status: draft
category: measurement
type: guide
audience:
  primary: service-managers, performance-analysts
  secondary: designers, developers, qa, content-designers
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

This guide was developed by GDS to help you move from identifying and measuring KPIs, through to presenting and acting on the data.

The guide follows a seven stage iterative process:

![Diagram showing seven stage iterative process](/service-manual/assets/images/kpis/monitor.png)

##Establish a baseline

Establish a 'baseline' based on current performance trends by channel, against which changes to the service will be judged. This will help you pinpoint the effectiveness of your initiatives, and identify what worked.

###Checklist
1. Have you measured your current performance?
2. Have you compared your performance with similar types of service (e.g. by transaction category)?
3. Do you know who your users are in terms of age, disability, socio-economic group and internet usage?

It is good practice to look at performance trends over time, rather than take a snapshot at a particular point in time. Peaks and dips in performance are then measured relative to this base (or trend) line which helps to identify the effect of communications or design initiatives. It also reveals seasonal variations in performance.

##Aggregate data

Collect and aggregate performance information from multiple sources and across multiple channels. Make sure you understand what this will mean in terms of system requirements.

###Checklist

1. Have you collected data on costs, usage and performance?
2. Have you collected performance data from digital and non-digital channels?
3. Do you know how many people use the service, by channel?
4. Have you aggregated performance data to enable it to be easily combined?

Combining data often reveals useful insights, for example into service efficiency (eg cost per transaction and total cost to serve) or proportional usage by channel (eg percentage digital uptake vs post, phone etc).

##Analyse and visualise data

Communicate performance information to your users through the appropriate dashboards, reports and alerts. Highlight specific segments that you know users are interested in, and make sure that your visualisations are simple, actionable and contain minimal amounts of [chart junk](http://en.wikipedia.org/wiki/Chartjunk).

###Checklist

1. Have you done any segmentation (i.e. analysed performance data by segment)?
2. Have you designed the appropriate dashboards, reports and alerts?
3. Are your data visualisations visible to their intended audience?
4. Have you followed best practice design principles for data visualisation?

###Typical segments include:

- **Channel used to access service:** through which channel(s) did the user find out about and attempt to use the service?
- **New vs. repeat visitors:** are first time users behaving differently to those who have used the service before?
- **Geographical region:** how popular is the digital service by region and how does that compare with online penetration in general?
- **Product type:** does the user experience vary depending on the type of product or service?
- **Value:** is performance dependent on the monetary value of the product or service being sought?
- **Dashboards:** are objective-focused and will help inform decisions, often with the help of real-time data.
- **Reports:** provide regular, scheduled snapshots of data and tend to require extra - context and time to digest.
- **Alerts:** are used to inform the users about a change or an event, often using attention-grabbing delivery mechanisms.

By making your visualisations clearly visible you maximise the likelihood that the information will be acted upon - and services thereby improved.

Best practices include:

- **keeping charts plain:** don’t use shading, 3D or other distracting effects
- **removing clutter:** don’t use trend lines, grid lines, unnecessary labelling
- **not using pie charts:** they require more space and are harder to read than bar charts
- **using text where a chart adds nothing.**

###Example

**Developer dashboards**

On GOV.UK dashboards keep our developers and editors up to date with the latest status reports and highlight any critical problems. They are displayed on large screens close to developers’ desks so that they can be easily monitored. This means that bugs can be spotted quickly and fixed.
Read this [blog article](http://digital.cabinetoffice.gov.uk/2012/02/08/radiating-information/) to find out how these dashboards were developed.

##Monitor, iterate, and improve

Test a range of performance improvement initiatives and monitor to see which work well. These can be piloted on a subset of your users to minimise risk. Implement the best performing solutions widely and then repeat this process relentlessly: what you measure will change over the course of a product or project's lifetime.

###Checklist

1. Are you monitoring and acting upon user feedback?
2. Have you done any  [A/B testing](http://en.wikipedia.org/wiki/Ab_testing) or [multivariate testing?](http://en.wikipedia.org/wiki/Multivariate_testing)
3. Have you taken an iterative approach to developing your service?

Any service that meets user needs will include an element of user feedback. This should be monitored and acted upon so as to continually improve the service for users.
A range of options are available for improving the overall performance of a service. The following examples are based on the [4 Ps](http://en.wikipedia.org/wiki/Marketing_mix) of marketing:

**Price:** can the price be changed, for example to attract people to the digital channel?

**Product:** can the user experience be improved (eg from user feedback, user testing, A/B testing, multivariate testing)?

**Placement:** can the digital service URL be placed on printed materials and voice recordings?

**Promotion:** can greater use of email and social media be used to promote repeated use of the digital service?
Taking an iterative approach to service development increases the pace of improvement and minimises the risk of failure. Don’t wait until the end to do this, it should happen continuously throughout the process.

##Further reading

[This article](http://www.guardian.co.uk/money/2010/dec/22/amazon-top-consumer-satisfaction) in The Guardian shows online customer satisfaction scores for retailers. These scores are based on the Customer Satisfaction Index.

[Great article](http://radar.oreilly.com/2010/06/what-is-data-science.html) by Mike Loukides on the role of the data scientist.

Total Cost to Serve is a [method (PDF, 79k)](http://www.hmrc.gov.uk/research/cost-of-time.pdf) for calculating the cost of a transaction for both the service provider and the user. HMRC have developed a method for calculating the cost of users time when interacting with government. This is important because some channels may be quicker to use than others.

[Designing with Data](http://www.fivesimplesteps.com/products/a-practical-guide-to-designing-with-data) is an excellent book by Brian Suda which helps you to design beautiful and powerful data visualisations.

[Juice Analytics](http://www.juiceanalytics.com/) is a great website with loads of useful resources on how to design and develop useful data visualisations and dashboards.

Edward Tufte’s [The Visual Display of Quantitative Information](http://www.edwardtufte.com/tufte/books_vdqi) is a seminal work on data visualisation and introduces the concept of chartjunk.

The [Flowing Data](http://flowingdata.com/) blog by Nathan Yau is a useful source of data visualisation news.

The [D3 Gallery](https://github.com/mbostock/d3/wiki/Gallery) is a stunning collection of data visualisations.

[Nicely presented overview](http://selection.datavisualization.ch/) of some of the tools available for data visualisation.

This [article in Wired](http://www.wired.com/business/2012/04/ff_abtesting/all/1) shows how A/B testing was used to good effect in Obama’s election campaign.

This [article in eConsultancy](http://econsultancy.com/uk/blog/2454-q-a-lovefilm-s-craig-sullivan-on-a-b-and-multi-variate-testing) shows how multivariate testing was used to improve conversion rates at Lovefilm.