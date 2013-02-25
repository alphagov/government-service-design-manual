---
layout: default
title: Completion rate
section: guidance
subsection: KPIs
status: draft
---

The completion rate measures the proportion of people who start a transaction and are able to complete it.

A transaction is a self-contained process that the service manager has defined in relation to the service. Typically this will be completing an entire transaction from end to end but where a transaction can be completed only partly online, it may be completing a discrete process within the transaction: booking an appointment or completing a part of an application, for example.

## Why measure completion rate? 

When users are unable to complete a digital transaction it can lead to avoidable contact through other channels. This in turn leads to low levels of digital take-up and customer satisfaction, and a higher cost per transaction.

Measuring end-to-end completion rates helps identify whether users have problems completing a transaction; subsequent analysis of drop-out rates at each step of a transaction can then pinpoint the specific processes that users are unable to complete.

## How to measure completion rate

The end-to-end completion rate can be calculated as the number of completed transactions divided by the number of started transactions; it is expressed as a percentage. 

### Before the transaction starts

Users should be told clearly what the outcome of the transaction will be, who can use it, how long it is likely to take, and what they will need to complete it (eg a reference number or credit card) before the transaction start point.  They will also need to be provided with the eligibility criteria and the costs to complete the transaction. This will help to reduce the dropout rate later in the transaction.


Typically this information will be given on a single introduction page. In some cases, there may be a set of pages that checks a user’s eligibility based on information they provide.

It will not be possible to bypass introduction pages via links or search engine results. Users who try to access the start page directly should be referred back to the introduction page.

### Start and end points

Transactions will begin and end on GOV.UK to allow GDS to monitor completion rates. 

Transactions started will be defined as the number of referrals from the introduction page to the start page. Transactions completed is the number of referrals to the end page.

This data will be freely shared with service managers.

[IMAGE TO BE ADDED]

### Drop-outs

Service managers will be responsible for measuring and monitoring drop-out rates within transactions which are not hosted on GOV.UK. You should analyse this data and use it to improve your service.

You should build your service with unique URLs for each step / page. This will make your service much easier to measure.

### Saving or resuming progress during a transaction

Some services allow users to save a transaction mid flow and to resume it another time.

Ideally it should be possible to match saved transactions with resumed transactions so that, for the purposes of completion rate, they are treated as one continuous process.

Service managers could consider applying a nominal time limit to saved transactions after which, if they haven’t been resumed, they are classed as failed. Alternatively, saved transactions could be set to expire after a given length of time.

It should not be possible for a user to start a new transaction without having passed through the introduction page.                             

### Multiple endpoints

Some services have complex flows and there may be multiple endpoints for successful transactions. These will need to be defined by service managers and should point to appropriate end pages on GOV.UK so they count towards the completed transaction total.

### Offline fulfilment

Some transactional services have parts which are digital and others which are non-digital. For example, when granting a lasting power of attorney users can start and finish the process online, but are required to print, sign and post a copy of the form in the middle of the process. 

In these situations the second online part of the transaction should be treated as a resumption of the first part. Where this is not possible, the discrete digital parts of the service should be treated as separate tasks, with the measure completion rate measured for each. 

Offline parts of the process can still be measured but this is likely to be done through qualitative feedback (e.g from surveys, diary studies, focus groups).

## When to measure completion rate

In order to successfully measure your service completion rate we recommend follow the guidance below:

### Guidance/Tool

| | Inception | Alpha | Beta | Live |
| Must | Ensure all transaction pages have unique URLs

Develop a plan to measure completion rate throughout product development. 
Assess the available analytics tools | Benchmark task completion rate via usability testing and establish reasons for failed transactions

Procure digital analytics tool | Conduct another round of usability testing to ensure that task completion rates improve | Analyse task completion rate to continually improve the service

Carry out further usability testing to continually identify any usability problems and feed into service design |

### How frequently should I measure task completion rate?

Completion rate should be measured continuously. Once the service is live this will be done on GOV.UK.

### How can I test the completion rate before launch?

Before launch completion rates can be measured with usability testing. This should be iteratively tested with at least five people. 

This should identify [over 85%](http://www.nngroup.com/articles/why-you-only-need-to-test-with-5-users/) of usability problems. Users will be given a pre-determined set of tasks that reflect what needs to be done to use the service. These tasks will include all aspects of using the service that apply, such as registering, applying, submitting, verifying, amending and unsubscribing. If users are having difficulty completing tasks, further development should be followed by further rounds of testing.

### What about post launch?

Digital analytics will be the primary method for measuring task completion rates post launch. Please note that this relies on extra configuration in the analytics tool. It will not be available by default.

The focus of the service team’s activities will be to continually improve this by monitoring where users are dropping out of the transaction process and testing out new designs for those pages. End to end completion rates will be piped automatically from GOV.UK’s digital analytics into the [Performance Platform](https://www.gov.uk/performance) and will be publicly available from the point of launch.

Further usability should also be carried out once a service have gone live to measure use of the service and identify any issues and improvements that can be made. 