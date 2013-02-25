---
layout: default
title: Monitoring
subtitle: Tracking the status of a service
section: guidance
subsection: Maintaining services
status: draft
type: 
audience: 
    primary: tech-arch, developer
    secondary:
theme: Running a service
assets: local
---

Any online application should have some tools dedicated to alerting the
people running the service to problems. This might involve low level
issues involving the infrastructure underpinning the service to a sudden
high rate of user errors.

## Guidance/Tool

### Setup monitoring early

Monitoring is not something that should be left to the end, to be tacked
on as part of running the final production service. By talking about
monitoring, and agreeing an approach each, you are more likely to build
useful checks as you go along. Writing tests at the same time as writing
code is common, monitoring checks can be viewed as tests for the running
system.

### Include high level checks

Often this is seen through a very technical lens, so looking only at web
application performance, available disk space or memory usage. Although
these are important it's also essential to track these alongside more
business related metrics. For example being able to compare page loading
tests with failed transactions and application errors can alert you to
problems and help identify the cause at the same time. It also grounds
conversations about low level problems (disk space, slow performance) in
terms of the service performance. 

### Errors are interesting

When errors occur they should be recorded and tracked over time. Errors
always contain interesting information, potentially about a user
problem, attack in progress, failing system or just a capacity problem.
It's important to be able to see errors both at the level of the entire
system and related to a particular application or machine.

### Widely available

The monitoring system, or rather any dashboards, interactive tools or
reports, should be as widely available as possible. They should ideally
be useful outside just the group responsible for the day to day
operations and systems administration.

## Why we do this

Knowing the current state of your service and infrastructure can help
identify problems before they happen, as well as alert you to issues
that need immediate attention.

The main goals are:

* To be alerted to problems affecting the availability of the service so
  they can be resolved
* To aid with capacity planning activities by providing metrics over
  time
* To identify potential future problems
* To find areas of improvement, for instance badly performing
  systems or inefficient services

## Further reading

* [Radiating
  Information](http://digital.cabinetoffice.gov.uk/2012/02/08/radiating-information/)
Examples of some of the monitoring dashboards used while building GOV.UK
* [Open Source
  Monitoring](https://speakerdeck.com/obfuscurity/the-state-of-open-source-monitoring) 

