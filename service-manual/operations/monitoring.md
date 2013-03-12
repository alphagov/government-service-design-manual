---
layout: detailed-guidance
title: Monitoring
subtitle: Tracking the status of a service
category: operations
type: guide
status: draft
audience: 
    primary: tech-arch, developer
    secondary:
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
---

Any online application should have some tools dedicated to alerting the people running the service to problems. This might involve low level issues involving the infrastructure underpinning the service to a sudden high rate of user errors.

Knowing the current state of your service and infrastructure can help identify problems before they happen, as well as alert you to issues that need immediate attention. The main goals are:

* to be alerted to problems affecting the availability of the service so they can be resolved
* to aid with capacity planning activities by providing metrics over time
* to identify potential future problems
* to find areas of improvement, for instance badly performing systems or inefficient services

## Setup monitoring early

Monitoring is not something that should be left to the end, to be tacked on as part of running the final production service. By talking about monitoring, and agreeing an approach, you are more likely to build useful checks as you go along. Writing tests at the same time as writing code is common, monitoring checks can be viewed as tests for the running system.

## Include high level checks

Often monitoring is seen through a very technical lens, so teams may only look at web application performance, available disk space or memory usage. Although these are important it's also essential to track these alongside more business related metrics. 

For example being able to compare page loading tests with failed transactions and application errors can alert you to problems and help identify the cause at the same time. It also grounds conversations about low level problems (disk space, slow performance) in terms of the service performance. 

## Errors are interesting

When errors occur they should be recorded and tracked over time. Errors always contain interesting information, potentially about a user problem, attacks in progress, failing systems or just a capacity problem. 

It's important to be able to see errors both at the level of the entire system and related to a particular application or machine.

## Widely available

The monitoring system, or rather any dashboards, interactive tools or reports, should be as widely available as possible. They should ideally be useful outside just the group responsible for the day to day operations and systems administration.

## Further reading

* [Radiating Information](http://digital.cabinetoffice.gov.uk/2012/02/08/radiating-information/)Examples of some of the monitoring dashboards used while building GOV.UK
* [Open Source Monitoring](https://speakerdeck.com/obfuscurity/the-state-of-open-source-monitoring) 

