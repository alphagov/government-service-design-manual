---
layout: default
title: Implementing completion rate
section: guidance
subsection: KPIs
status: draft
---

* Toc
{:toc}
    
#Implementing completion rate
This guidance covers implementing the completion rate KPI, including what to do when and how frequently to measure. Please refer to [Defining KPIs](/guides/kpis/definingkpis.html) for a definition of completion rate.

##Guidance/Tool
<table>
<tr><th></th><th>Inception</th><th>Alpha</th><th>Beta</th><th>Live</th></tr>
<tr><td>Must</td><td>Ensure the transaction start and end pages have unique URLs</td><td>Benchmark task completion rate via a survey or usability testing and establish reasons for failure</td><td></td><td>Measure task completion using digital analytics software. Continually improve and optimize the service using task</td></tr>
<tr><td>Should</td><td>Develop a plan to measure completion rate throughout product development. Assess the available analytics tools</td><td>Procure digital analytics tool</td><td>Conduct more than one round of usability testing to ensure that task completion rates improve</td><td>Carry out further usability testing to continually identify any usability problems and feed into service design</td></tr>
<tr><td>May</td><td>Set task completion time as a KPI if you have identified that users want to complete tasks as quickly as possible. This may not be the case for some services e.g. Lasting Power of Attorney</td><td>Benchmark task completion time</td><td>Measure task completion time via usability testing</td><td>Use A/B or multivariate testing to test alternative designs, implementing those which improve completion rates</td></tr>
</table>

##How frequently should I measure task completion rate?
The user survey (alpha) and usability testing (beta) should be carried out at least once but it is recommended that these methods are used to support the ongoing development of the service even after launch. Digital analytics measurement is a continual process and should carry on indefinitely after launch.

##What level of completion should I aim for?
You should aim for a 100% rate of completion.

The minimum you should plan to achieve by launch is 80%. This will be measured via usability testing of at least ten people in either a lab-based or remote setting. This should identify [over 90%](http://www.useit.com/alertbox/20000319.html) of usability problems. Users will be given a pre-determined set of tasks that reflect what needs to be done to use the service. These tasks will include all aspects of using the service that apply, such as registering, applying, submitting, verifying, amending and unsubscribing. An average of 80% must be achieved across all tasks for the service to go live. Further rounds of testing may be necessary if users are having difficulty completing tasks.

##What about post launch?
Digital analytics will be the primary method for measuring task completion rates post launch. Please note that this relies on extra configuration in the analytics tool. It will not be available by default.

The focus of the service team’s activities will be to continually improve this by monitoring where users are dropping out of the transaction process and testing out new designs for those pages. Completion rates will need to be piped automatically from the product’s digital analytics into the [Performance Platform](http://www.gov.uk/performance) and will be publicly available from the point of launch.

##Further reading
[Available analytics tools](/handbook/96/)
[Analytics team skills](/handbook/11/)
[Optimizing conversion rate](/handbook/73/)

#Optimizing conversion rate
Conversion rate is considered a good [web KPI](http://www.kaushik.net/avinash/rules-choosing-web-analytics-key-performance-indicators/) because it focuses on outcomes and is traceable to a website’s goal. However, looking at conversion rates in isolation fails to generate user insight or suggest actions for improvement. This guidance suggests approaches that product teams may use to optimize a service post launch.

##Methods for improving conversion rate
The eConsultancy Conversion Report 2010 lists the following methods for improving conversion rates:
* A/B testing
* User journey analysis
* Multivariate testing
* Cart abandonment analysis
* User testing
* Online surveys/user feedback
* Event-triggered/behavioural email
* Segmentation
* Copy optimisation
* Expert usability reviews/consultancy
* Pinch-point analysis

An interesting report by McKinsey entitled [E-government 2.0](https://www.mckinseyquarterly.com/E-government_20_2408) finds that governments have failed to deliver on investments in online services and argues that more should be invested in web analytics expertise. It cites the example that Google has 50 to 200 online experiments running at any one time and argues that governments should adopt a similar approach. It stresses the importance of *data-driven decision making* and specifically mentions the following methods:
* set clear goals and measure success
* track quantitative data through dashboards
* collect user feedback
* test and learn through usability testing and piloting
* test and learn through A/B and multivariate testing

##Segmentation
The eConsultancy Conversion Report 2010 found that businesses whose online conversion rates improved looked at twice as many segments compared to businesses whose online conversion rates did not. It recommends:
                                                                              
> Analyse conversion for different segments, specifically, new versus repeat visitors, channel used to access site, product and RFM (Recency, Frequency, Monetary value).

It gives the following examples of segments:

* Purchased any product: Compare the activity of those that have purchased (happy with the brand) to those that have not. This is most powerful when trying to understand what non-customers are doing on the site. Put any inactive users into a reactivation programme.
* Repeat Visitors: Understand the difference between the user behaviour of returning prospects and first time visitors.
* Channel used to access site: How the user entered your site is a big clue to how they need to be treated. If they entered with the search term “cheap holiday” send them emails on cheap deals. If they entered via a PPC ad promoting a “luxury break” send them an email about luxury breaks.
* First time buyers: Paying customers act differently to first time browsers. By splitting the data in this way you can focus on your buying customers. It is these users that are important to you. Use the information extracted via online behaviour to understand how these users like to shop.
* Product segments: It is important to have segments per product so you can analyse the activity of users based on what they like to buy. Trends in behaviour may differ depending on what product is being purchased.
* Advocates: Separate regular visitors from infrequent visitors. It is the regular visitors you want to pay attention to. Understanding these users is critical because how they behave is how you want all users to behave.
* Engagement: Recent visitors are more engaged and can be contacted more frequently.
 
Avinash Kaushik also sees value in pushing beyond a simple conversion rate, although it’s still one of his [Six Web Metrics/KPIs to die for](http://www.kaushik.net/avinash/rules-choosing-web-analytics-key-performance-indicators/). He also looks at the number of sessions taken to complete a transaction (segmented into ranges), number of days since last visit (segmented into ranges) and number of visits per visitor (segmented into ranges).

##Retention rate
The [McKinsey e-performance report](https://www.mckinseyquarterly.com/E-performance_The_path_to_rational_exuberance_975) also sees value in considering new versus repeat users and stresses the importance of customer retention. It looks at conversion within the broader process of attraction, conversion and retention. That is, getting users into a service, getting them through it and getting them to re-use it.

*Perhaps our most important finding is the power of retention. A 10 percent increase in the rate at which visitors are converted into repeat customers drives a 10 percent improvement in the net present value of a company’s expected cash flows*

An analysis of costs and revenues per visitor, per customer and per repeat customer shows that repeat customers are cheaper to acquire and worth more. McKinsey says more effort is needed to ensure customer loyalty. In practical terms this may mean keeping information so a user never has to provide the same information twice, or emailing users to remind them to use a service again.

