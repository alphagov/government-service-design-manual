---
layout: gsdm
title: Hosting
section: guidance
subsection: Maintaining services
status: draft
---

* Table of contents
{:toc}
    
# Hosting
The software running your service will need servers to run on. This
guide aims to help with deciding how you host your applications and the
things to think about if selecting a vendor.

## Guidance

The recommended approach is to involve a small cross-functional group of
people to quickly access different options, shortlist suppliers,
interview  and finally make a descision. This group should include people with a
knowledge of the available Procurement options and acceptable costs but
must include people with a hands-on technical understanding of the service and
underlying software.

It is important to keep good notes from any interviews or deliberation
sessions and to access different suppliers equally. A scoring matric can
help here.

It is worth stating that it's very common to use multiple suppliers.
This may be due to them offering different but compatible services or
potentially for additional redundancy. This can be technically
challenging but for larger projects can provide extra resiliance.

### Types of provider

There are a number of different approaches taken by suppliers of hosting
services, which can make comparing offerings difficult. The following is
intended only as a brief introduction.

This advice is complicated because many service providers redefine the
meanings for marketing reasons. In particular Infrastructure as a
service and Platform as a service are marketable at the moment and often
used incorrectly. Alwats look into the details of the services being
offered.

#### Ownership

For particularly large projects with very specific requirements you may
decide that purchasing hardware and even running a dedicated data centre
is suitable. The costs and timescales involved here are very high and
this is unlikley to be the best option in most cases.

#### Co-location

Many providers offer co-location services which is where you purchase
your own hardware to put into a managed data centre. This provides a
great deal of flexibility but can introduce lead times and other
physical constrains. It also requires a wide range of technical
specialist skills. 

#### Shared or managed hosting

Lots of service providers have a shared or managed hosting option. This
tends to mean renting specific virtual or physical machines for fixed
periods of time. Different suppliers other different management
services, some just manage the underlying machine while others will
support the operating system and even specific applications running on
the machines.

#### Infrastrucure as a service

In the last several years Infrastructure as a service has become a
common approach to managing hosting and infrastructure requirements.
This tends to involve a capability to rapidly add or remove capacity,
often in minutes, and to be billed only for what is used. This provides
a great deal of flexibility and the ability to hold costs down but also
requires a degree of technical skill to manage well. 

#### Platform as a service

Similar to Infrastructure as a service above, Platform as a service
offerings tend to allow for quickly adding or removing capacity and fine
grained pay on demand pricing. The difference is is that you are
abstracted away from the underlying infrastructure completely. The unit
here is the running application, not a virtual or physical machine.
Using a Platform as a service places a number of constrains on the
software architecture but can move the support burden for parts of the
stack onto the supplier.

## Why we do this

Making a decision about your hosting supplier involves weighing up a
wide range of different components, including:

* Technical requirements of your software applications
* Future capacity requirements
* User support
* Relisance and redundancy
* Mandated Government information security requirements
* Government network connectivity
* Cost of service
* Flexibility and on-demand billing
* Different operating models (physical, co-location, infrastructure as a service,
  platform as a service)

This is unfortunately a technical field with many options. Seemingly
similar services can have wildly different architectures or different
cost models can result in large differences in total cost of ownership.
It is recommended to involve technical colleagues or trusted third
parties in any discussions and descision.

## Further reading

* [Why IaaS?](http://digital.cabinetoffice.gov.uk/2012/09/25/why-iaas/)
  a blog post about why the Government Digital Service chose an
  Infrastructure as a service for GOV.UK
* [G-Cloud](http://gcloud.civilservice.gov.uk/) the procurement
  framework intended for purchasing Infrastructure as a Service,
  Platform as a Service and Software as a Service products



# Monitoring
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

