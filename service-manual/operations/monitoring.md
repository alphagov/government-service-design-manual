---
layout: detailed-guidance
title: Monitoring
subtitle: Tracking the status of a service
category: operations
type: guide
status: draft
audience: 
  primary: web-ops
  secondary: service-managers, developers, tech-archs
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
exclude_from_search: true
---

Any online application should have tools dedicated to alerting any problems to those running the service. These can be low-level issues involving the infrastructure supporting the service, to a sudden high rate of user errors.

Know the current state of your service and infrastructure. It will help you to identify problems before they happen, as well as alert you to issues that need your immediate attention. The main goals of monitoring are to:

* be alerted to problems affecting the availability of the service so they can be resolved
* help with capacity planning activities by providing metrics over time
* identify potential future problems
* find areas of improvement, eg badly performing systems or inefficient services
* be able to identify the root cause of an outage after the fact, based on data collected during the problematic event

## Set up monitoring early

Don’t leave monitoring to the end, tacked on as part of running the final production service. 

Talk about monitoring and agree on an approach. This way you’re more likely to build useful checks as you go along. Writing tests at the same time as writing code is common, monitoring checks can be viewed as tests for the running system.

## Include high-level checks

Often monitoring is seen through a very technical lens, so teams may only look at web application performance, available disk space or memory usage. Although these are important, it’s also essential to track these alongside more business-related metrics.

For example, being able to compare page-loading tests with failed transactions and application errors can: 

* alert you to problems
* help identify the cause at the same time
* ground conversations about low-level problems (disk space, slow performance) in relation to service performance

## Errors are interesting

When you have an error, record it and track it over time. Errors always contain interesting information, which can be about:

* a user problem
* attacks in progress
* failing systems
* problems with capacity

It’s important to be able to see errors that are part of the overall system and errors specifically related to a particular application or machine.

## Make data widely available

Make the following as widely available to everyone as possible, and also useful to teams outside of operations and systems administration:

* monitoring system
* dashboards
* interactive tools
* reports

## Further reading

* the [Radiating Information](https://gds.blog.gov.uk/2012/02/08/radiating-information/) blog post gives examples of some of the monitoring dashboards used while building GOV.UK
* [open source monitoring](https://speakerdeck.com/obfuscurity/the-state-of-open-source-monitoring) explains some of the terminology and available open-source tools
