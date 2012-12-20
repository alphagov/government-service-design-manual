#Incorporating analytics data into design
As a service manager, you will need to hire [skilled digital analysts](/handbook/11/) to continually monitor how users are interacting with your service and report back to you on a regular basis. This guidance describes what those reports should cover, when you should ask for more detailed analysis and what you should ask for.

##Completion rate
Completion rate is one of the four [digital by default service KPIs](/handbook/134/) and should form the basis of any conversations with your digital analysts on how well the service is performing. Completion rates should be reported on a daily basis and continually monitored to identify any significant changes, either positive or negative. Longer term trends should also be monitored to identify where performance is in line with the expected trajectory or where it deviates significantly from the mean.

##When to ask for more detailed analysis
Most measurements taken on a random sample follow what is called a _normal distribution_ or _bell curve_. This means that we can estimate the probability that a measurement falls within a certain range of values, known as _standard deviation_.

![Normal distribution or bell curve](http://upload.wikimedia.org/wikipedia/commons/8/8c/Standard_deviation_diagram.svg)

For example, the [mean height of American men](http://en.wikipedia.org/wiki/Human_height#Average_height_around_the_world) is 70” with a standard deviation of 3”. Because the heights follow a bell curve, we can estimate that 68% of US men are within the 67” - 73” height range, 95% are between 64” and 76”, and 99.7% are between 61” and 79”.

We take a similar approach to digital analytics, using these probabilities to help us decide when to do more detailed analysis into the causes of high or low completion rates. As a rough guide, use the following table:

<table>
<tr><th>Performance</th><th>Action</th><th>Approximate frequency for daily measure</th>
<tr><td>Within 1 standard deviation of the mean</td><td>No further analysis required. This will account for 68% of all outcomes.</td><td>5 days per week</td></tr>
<tr><td>1-2 standard deviations from the mean</td><td>Further analysis may be needed to explain why performance has peaked/dipped.</td><td>Twice a week</td></tr>
<tr><td>Over 2 standard deviations from the mean</td><td>Further analysis is needed.There is only a 5% probability of this outcome.</td><td>Every three weeks</td></tr>
</table>


##What to ask for
If completion rates have significantly increased or decreased you may want to request more detail from your digital analysts. There are several avenues that you should ask them to explore:
* Task breakdown: this will probably appear on a standard report but if not a breakdown of completion rates for each task that contributes to the overall KPI will reveal if there are any tasks which are proving particularly easy or difficult for users to complete
* Funnel analysis: each task will be configured in your analytics tool to monitor the specific pages where users are dropping out of the transaction process. Each page will have a percentage of users dropping out and you will want to probe this further.
* Referral: how are users getting to the service? Is there a difference in completion rates for users searching on particular terms or being referred from a particular source?
* Device: is the device a user is using to access the service having a positive or negative impact on completion rate?
* Time: do completion rates vary by time of day or day of week?
* Geography: do completion rates vary by geographical region?

Your digital analysts will be able to provide these insights for you if they are not part of the standard KPI report.

##Visibility and visualisation
Incorporating analytics data into the service design process is challenging and can require a cultural change within the organisation. Publishing performance data is one way to bring about that cultural change. This is something that we have done with our [Performance Platform](https://www.gov.uk/performance).

Using [best practices in data visualisation](/handbook/141/) can also help to ensure that data is taken seriously within your organisation. Simple, clear, minimal charts can increase the impact of data and prevent reports from being overlooked or ignored.

##Why we do this
We use data to drive decision making wherever possible to ensure that user needs are met and to guard against intuitive decision making that can be highly subjective.

##Further reading
Article on [data driven delivery](http://digital.cabinetoffice.gov.uk/2012/07/24/data-driven-delivery/) by Richard Sargeant on the GDS blog
