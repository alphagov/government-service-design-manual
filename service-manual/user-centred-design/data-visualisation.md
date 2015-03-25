---
layout: detailed-guidance
title: Data visualisation
subtitle: Creating valuable and meaningful graphics to help analyse data
category: user-centred-design
type: guide
audience:
  primary: performance-analysts
  secondary: service-managers, qa
status: draft
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

As we surface more data about government services, we need to make sure that the visualisations of it are easy to understand, visually compelling and prompt action. To do that, we need to have a consistent visual grammar, for use both within GDS and across government.

## Introduction

This guide sets out 4 principles of good data presentation, with easy to follow checklists to help you achieve this. For context, we've added examples of how the principles have been employed at GDS. The principles and examples found in this guide are likely to evolve as we find new challenges and applications for them.

### Best practice
There are many examples of best practice style guides already in place. For example, [The Economist](http://www.economist.com/) has a clearly defined house style that allows its readers to readily identify and understand their visualisations. They publish a new visualisation every day in their [Graphic Detail](http://www.economist.com/blogs/graphicdetail). This guide attempts to build on the best practice from a range of organisations.

### GDS example
The [GOV.UK Performance Platform](https://www.gov.uk/performance) helps the government make decisions based on data, often presented through innovative visualisations (built using [D3.js](http://d3js.org/)). The example below compares weekly visitors to GOV.UK with the two main websites it replaced.
![Weekly uniques](/service-manual/assets/images/data-visualisation/weekly-uniques.png)

## Telling the story

To effectively tell the story behind the numbers, you need to understand both your audience and the data.

Only use visualisations if they make the story clearer. In many cases, a good table or words may communicate better than a visualisation. If there are very few data points (eg top rate income tax down 5%, all other rates unchanged), it’s clearer to write a sentence than draw a picture. If every item must be read precisely (to several decimal places) then a table is best.

A good table will be clear and uncluttered. The data should be easy to read with the same decimal places or rounded and sorted into a logical order. Don't use too many different types of font, and make sure your data is referenced.

But visualisations often are the right answer and the data is the most important feature. It should tell its own story and it's best to not try to say too much in one go. Keep charts simple, cutting down on unnecessary items and jargon. 

Explanatory text will be needed in some cases, but it should not simply repeat the story being told in the visualisation. A well written chart title can reinforce the story of the data and reduce the amount of additional text needed.


### Checklist
1. Who are your audience?
2. How much detail do they need?
3. What story does the data tell?
4. Do you need a visualisation?

## Choosing your visualisation and templates

Choosing the right visualisation will help the data tell its own story and give powerful insight. There are many ways of displaying information visually. 

Most computer programmes come with a range of visualisations. There also visualisation tools available online: for example, this [blog showcases some free ones](http://www.computerworld.com/article/2507728/enterprise-applications-22-free-tools-for-data-visualization-and-analysis.html) and GDS has produced [a guide to infographics](/service-manual/assets/documents/Infographics_Guide.pdf)

Each chart has its own strength. Below are the core 5 with templates (a [Google spreadsheet of these is available](https://docs.google.com/spreadsheet/pub?key=0AtlK1009bUPidGVET1h5Z0xNSFJvVXA5MFF2ZHBKakE&output=html)):

###Column chart
![Column Chart - Image 1](https://docs.google.com/drawings/pub?id=1kgh4dXLVTHtQ5Lhj9SevdaNqij3ncFJRnj0hSPzDmc0&w=358&h=250)

Strengths - comparing items, or a small number of time periods.

* Negative values below the x-axis.
* If needed, the target should be a single line that is visible but not too thick.
* Limit stacked columns to 3 segments.

###Bar chart
![Bar Chart - Image 2](https://docs.google.com/drawings/pub?id=1m5nrx667px-PbiHMHxYJO9M4qq1oLrvJqVhyMf-rBmY&w=358&h=223)

Strengths - comparing items, especially if they have long names or many items.

* Arrange bars in size order, from biggest to smallest (unless there's good reason, ie data needs to be represented alphabetically).
* Negative values to the left of the y axis.

###Line chart
![Line Chart - Image 3](https://docs.google.com/drawings/pub?id=1EPgsC32JClcIjApetWMFHLPNSZQEoViSacc0ySmclp4&w=337&h=237)

Strengths - comparing over time or between variables for a single item (eg site traffic vs site performance)

* Limit number of data sets to 4.
* Keep axis labels horizontal.
* Use line points to differentiate between data types (use line dots for projections and estimates only).

###Pie chart
![Pie Chart - Image 4](https://docs.google.com/drawings/pub?id=1t5ECozT_ClsmPFNID9j6mYKkUVeM9t0bW_vAah5T_5A&w=337&h=237)

Strengths - simple share of total. Use with caution as column or bar charts are often better. Limit to 2 segments to avoid confusion.

* The largest segment should be at 12 o’clock going clockwise.
* Label the chart directly and avoid text inside segments.

###Scatter chart
![Scatter Chart - Image 5](https://docs.google.com/drawings/pub?id=1yBH8YxdFGHM4N9xXuZeD4M9CIVcxYicfRznktcSLDKw&w=337&h=237)

Strengths - relationships between variables where there are many items (eg volume vs cost for numerous transactions)

* Limit items to two to avoid confusion.
* Include trend line if required. This should be a single solid line.

There is more help on choosing the right chart [here](http://www.verstaresearch.com/types-of-charts.jpg). It is important to not confuse your audience.

###Checklist
1. What visualisations are available?
2. Have you chosen the right visualisation for the data?
3. What about [infographics](https://en.wikipedia.org/wiki/Information_graphics)?

###Worked example
Choosing the correct visualisation is important and at GDS we reviewed what was being used in the performance dashboard. As the example below shows, Pie charts with many items are not clear. We used a stacked bar chart to better represent the data.

####Before
![Bad Pie Chart - Image 6](https://docs.google.com/drawings/pub?id=1LWpYB-dagFps939s7RP9xRvMgqn0UsJcUv_d_RxY8Xc&w=337&h=237)

* Too many segments make them hard to compare.
* Too many colours made the chart confusing.
* The chart includes cost recovery as a negative which cannot be represented properly in a pie.

####After
![Good stacked chart - Image 7](https://docs.google.com/drawings/pub?id=12tme8E2COvhWn2fXRrzJNtVo_a1ClwizHeMSssn2Wbw&w=337&h=237)

* This stacked chart is much clearer.
* Comparisons can easily be made and sorting the data provides quick insight.
* The stacked chart provides additional information which could not be visualised in the pie.


## Creating your visualisation

Keep in mind these useful tips when creating your charts:

* start axes at zero unless there's good reason not to (ie data is clustered at high values)
* don't say too much, limit the number of data sets
* if needed, put legend at the top of the chart in the same order as the data in the chart
* maximise the space available to the chart and remove chartjunk
* include units of measurement in the chart title or directly on the axis, avoid doing both
* keep colours simple, do not repeat/alternate or use opposites - use the GDS Colour palette
* use the same colour when reporting a single data set

Keeping your visualisation simple will help the data tell its own story.

[Chartjunk](https://en.wikipedia.org/wiki/Chartjunk) is anything in your visualisation using ink that actively reduces clarity. Avoid:

* 3D effects
* borders
* unnecessary axes lines
* random colours or backgrounds
* unnecessary text

Know your audience so you give the right amount of supporting information. External or non-technical audiences will need more explanation but internal or expert audiences may find this tedious. Don't use the text to simply repeat what's being said by the data.

Visualisations should avoid too much data. Only include what's relevant. If the trend is obvious, don’t include a trendline. Sometimes it may be more effective to focus on high-value items only (if you're being selective, be open and clear about this).

Poor colour choice can change how the data is perceived in a visualisation. For example, red is strongly associated with negative performance so is best avoided for positive/normal figures.

Colour blindness makes it difficult for a user to differentiate between data sets. Labelling charts directly and different line styles can help. If your visualisations are likely to be printed it's important the colours work in greyscale as not all users will have high quality printers.

###Checklist
1. Have you removed the [chartjunk](https://en.wikipedia.org/wiki/Chartjunk)?
2. Have you got the right amount of supporting information?
3. Have you got the right amount of data?
4. Have you used the GDS colour palette?
5. Will the colours work for people with colour blindness or greyscale?
6. Are your colours appropriate for the data?

###Worked example: GDS performance dashboard
The example below from the GDS senior management dashboard shows how avoiding chart junk and limiting the number of datasets can enhance your visualisation.

####Before
![Poor stack bar - image 8](https://docs.google.com/drawings/pub?id=1X2JceYqVi-2LhiGJXstzPgGKkhv-ACGPNhjS09KewJU&w=337&h=237)

The legend accounts for a quarter of display space. The Y axis quotes £ and not £m. The segments are profiles and proportionate for each time period, so the bright colouring adds no extra detail. The mix of bar and line is confusing with so much information in the chart.

####After
![Good line chart - image 9](https://docs.google.com/drawings/pub?id=1IUt51enQ6kfdXz07R1n4lC68Z3Hl-bMkMGL9cmAeIdc&w=337&h=237)

The stacked column gave a level of detail which wasn't needed. This has been rationalised to best suit the audience. Axes have been standardised. The legend has been relocated giving the chart more space. Heavy grid lines and axes have been removed to give a clearer display.


## Be open and honest
>“We want transparency to become an absolutely core part of every bit of government business.”
>- Francis Maude

Being open and transparent supports the [Open Data White Paper](/government/uploads/system/uploads/attachment_data/file/78946/CM8353_acc.pdf). Similarly, our Open Public Services agenda is built on transparency. Sourcing data builds trust and credibility. Providing contact information shows ownership but also helps collaboration and information sharing.

When presenting data be aware of axes and scales. Data can be misrepresented by only showing a selection if it isn't clear why an extract has been chosen. Consider where the visualisation might be published. For example, if published alongside other visualisations, the reader is likely to assume the scales are consistent. This might change how your data is perceived.

###Checklist
1. Have you referenced data with a URL?
2. Have you provided contact details (eg a mailto link)?
3. Is it clear whether data is internal or public?
4. Have you been open and transparent with data?

###Worked example

![Bad Line chart - image10](https://docs.google.com/drawings/pub?id=1i4V_-F50fyfpJ6TwZZkJJAhI5siJcGX0sgwUKjN-dYs&w=337&h=237)

* Chart only shows a select few data points
* The small range on y-axis has exaggerated the fluctuations in data
* Y-axis starts at 60 but chart suggests data has reached lowest possible value

![Good line chart - image 11](https://docs.google.com/drawings/pub?id=1D0njUcS428O0vvA8-3Gla-ijAhv1jgo4rO97vVS-qJo&w=337&h=237)

* Y-axis has been formatted to show the full range of data
* The data fluctuations can be seen in the context of the wider data series.


## Further reading

More information about the [Performance Platform](/service-manual/measurement/performance-platform.html).

This [chart chooser](http://www.verstaresearch.com/types-of-charts.jpg) from Andrew Abela builds on the work of Gene Zelazny’s classic book Saying it with Charts. This [interactive tool](http://labs.juiceanalytics.com/chartchooser.html) from Juice Analytics helps guide your chart choice through filters.

Brain Suda’s [*A Practical Guide to Designing with Data*](http://designingwithdata.co.uk/) provides a comprehensive understanding of how to best engage the audience with your data. Here is a video of [Brian Suda presenting on a section of his book at the 2012 DIBI conference](https://vimeo.com/27483938).

Dona M. Wong’s [*The Wall Street Journal, Guide to Information Graphics*](http://donawong.com/) details the do’s and don’ts of presenting data.

Edward Tufte’s [*The Visual Display of Quantitative Information*](http://www.edwardtufte.com/tufte/books_vdqi) is a seminal work on data visualisations and introduces the concept of chartjunk. Here is a [video of Edward Tufte discussing his theories on visual thinking and analytical design](https://www.youtube.com/watch?v=Th_1azZA2OY).

This article from the [Peltier Tech blog](http://peltiertech.com/ten-chart-design-principles-guest-post/) covers the ten chart design principles.

The [Flowing Data](http://flowingdata.com/) blog is a source of data visualisation news.
