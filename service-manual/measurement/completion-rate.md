---
layout: detailed-guidance
title: Completion rate
subtitle: Services all users can finish
status: draft
category: measurement
type: guide
audience:
  primary:
  secondary: service-managers, performance-analysts
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

When users are unable to complete a digital transaction it can lead to avoidable contact through other channels. This leads to low levels of digital take-up and customer satisfaction, and a higher cost per transaction.

Measuring end-to-end completion rates helps to identify whether users have problems completing a transaction. Analysing the drop-out rate at each step of a transaction will pinpoint the specific processes that users fail to complete.

## What you will be measuring

The completion rate measures the proportion of people who start a transaction and are able to complete it.

A transaction is a self-contained process that's defined by the service manager, and is typically only completed when a user has been through the whole thing end to end.

A transaction completed only partially online means the user may complete a single process within the overall transaction, like booking an appointment or completing a part of an application.

## How to measure completion rate

1.    Work out the number of completed transactions.
2.    Divide it by the number of started transactions, expressed as a percentage.

### Before the transaction starts

Before the transaction has started, users should clearly be told:

* what the outcome of the transaction will be
* who can use it
* how long it is likely to take
* what they’ll need to complete it (eg a reference number or credit card)

Also provide users with the eligibility criteria and the costs to complete the transaction. This will help you to reduce dropouts later on in the transaction.

Typically the eligibility criteria will be given on a single page. In some cases, there may be a set of pages that checks a user’s eligibility based on information they provide.

Don’t make it possible for users to bypass your service’s start page via links or search engine results. Users who try to directly access another page in the transaction should be sent back to your start page, unless they are resuming a previously saved transaction.

### Start and end points

Transactions will begin and end on GOV.UK to allow GDS to monitor completion rates.

A transaction is considered to have been started and completed only when the user:

* begins at the start page
* reaches the end page

Data on the number of started and completed transactions will be shared with service owners.

![Completion rate](/service-manual/assets/images/measuring-completion-rates.png)

### Dropouts

Service managers will be responsible for measuring and monitoring dropout rates within transactions which aren't hosted on GOV.UK. You should analyse this data and use it to improve your service.

Build your service with unique URLs for each step or page. This will make your service much easier to measure.

### Saving or resuming progress during a transaction

Some services allow users to save a transaction mid flow and to resume it another time.

For the purpose of completion rate, make it possible to match your saved transactions with resumed transactions so they can be treated as one continuous process.

Consider applying a nominal time limit to saved transactions after which, if they haven’t been resumed, they are classed as failed. Alternatively, saved transactions could be set to expire after a given length of time.

### Multiple endpoints

Some services have complex journeys and there may be several points where a transaction can be successfully completed. Service Managers define these areas, and make sure they point to appropriate end pages on GOV.UK so they count towards the completed transaction total.

### Offline fulfilment

Some transactional services have both digital and non-digital areas. For example, when granting a lasting power of attorney, users can start and finish the transaction online, but are required to print, sign and post a form in the middle of the process.

In these situations treat the second online part of the transaction as a user continuing on from the first part. Where it’s not possible to match the two, treat the single digital parts of the service as separate tasks, with the completion rate measured for each.

Offline parts of the process can still be measured, but this is likely to be done through qualitative feedback (eg from surveys, diary studies, focus groups).

## When to measure completion rate

To successfully measure your service completion rate, GDS recommend that:

* during the discovery phase you
  * make sure all transaction pages have unique URLs
  * develop a plan to measure completion rate throughout product development
  * assess the available analytics tools
* in your alpha phase you
  * benchmark task completion rate via usability testing and establish reasons for failed transactions
  * procure digital analytics tools
* during your beta phase you
  * conduct another round of usability testing to ensure that task completion rates improve
* when a service is live you
  * analyse task completion rates so you can continually improve the service
  * carry out further usability testing to continually identify any usability problems and feed into service design

### How frequently should I measure task completion rate?

Completion rate should be measured continuously. Once the service is live this will be done on GOV.UK.

### How can I test the completion rate before launch?

Before launch, measure completion rates with usability testing. This should be iteratively tested with at least 5 people.

This should identify [over 85%](http://www.nngroup.com/articles/why-you-only-need-to-test-with-5-users/) of usability problems. Users will be given a pre-determined set of tasks that reflect what needs to be done to use the service. These tasks will include all aspects of using the service that apply, such as:

* registering
* applying
* submitting
* verifying
* amending
* unsubscribing

If users are having difficulty completing tasks, carry out more user testing after each new development.

### What about post launch?

Digital analytics will be the primary method for measuring task completion rates post launch. Please note that this relies on extra configuration in the analytics tool. It will not be available by default.

The aim of your service team’s activities will be to continually improve completion rates by monitoring where users are dropping out of the transaction process, and testing out new designs for those pages. End-to-end completion rates will be piped automatically from GOV.UK’s digital analytics into the [Performance Platform](/service-manual/measurement/performance-platform.html) and will be publicly available from the point of launch.

Further usability testing should also be carried out once a service have gone live to measure use of the service and identify any issues and improvements that can be made.
