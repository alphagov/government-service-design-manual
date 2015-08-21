---
layout: detailed-guidance
title: Service management
subtitle: Managing a running service
category: operations
type: guide
status: draft
audience:
  primary: web-ops, service-managers, tech-archs
  secondary: developers
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
---


The [Digital by Default standard](/service-manual/digital-by-default)
says that organisations should (emphasis on *operate* added):

> Put in place a sustainable multidisciplinary team that can design, build
and *operate* the service, led by a suitably skilled and senior service manager
with decision-making responsibility.

This implies a change to how many organisations have traditionally run services,
often with a team or organisation building a service separate from the one running it.
This change however does not mean ignoring existing good practice when it comes to service
management.

## Agile and service management

The principles of IT service management (ITSM) and those of agile do not necessarily
conflict -- issues can arise however when organisations implement rigid processes
without considering wider service delivery matters, or design and build services
without thinking about how they will be operated.

The [agile manifesto](http://agilemanifesto.org/) makes the case for:

* Individuals and interactions over processes and tools
* Working software over comprehensive documentation
* Customer collaboration over contract negotiation
* Responding to change over following a plan

It is too easy to position service management as opposed to agile as
traditional service management practices can be viewed as focusing on processes, tools,
documentation, planning and contract negotiation -- the items on the right hand
side of the points above.

However, the agile manifesto goes on to say:

> That is, while there is value in the items on the right, we
value the items on the left more.

To build and run a successful service you will need to work on suitable
processes and manage third party relationships. Using existing service management
frameworks (especially as a starting point) is one approach to this problem.

## ITIL

[ITIL](https://www.axelos.com/itil) (the Information Technology Infrastructure
Library) is one such framework. ITIL does a particularly good job of facilitating shared
language. For instance it's definition of a service is:

> A service is a means of delivering value to customers by facilitating outcomes
customers want to achieve.

The current version of ITIL currently provides 5 volumes and 26 processes
describing in detail various aspects of service management:

### Service Strategy

* IT service management
* Service portfolio management
* Financial management for IT services
* Demand management
* Business relationship management

### Service Design

* Design coordination
* Service Catalogue management
* Service level management
* Availability management
* Capacity Management
* IT service continuity management
* Information security management system
* Supplier management

### Service Transition

* Transition planning and support
* Change management
* Service asset and configuration management
* Release and deployment management
* Service validation and testing
* Change evaluation
* Knowledge management

### Service Operation

* Event management
* Incident management
* Request fulfillment
* Problem management
* Identity management

### Continual Service Improvement

### Functions

ITIL also describes four functions that should cooperate together to form
an effective service management capability.

* Service operations
* Technical management
* Application management
* Operations management

## The importance of implementation

The above processes and functions make for an excellent high level list of topics
to discuss when establishing an operating model for your service, whether
or not you adopt the formal methods. In many cases if you have well understood,
well established and well documented processes in place for all of the above
you should be in a good position to run your service.

When looking to combine much of the rest of guidance on [the service manual](/service-manual)
with ITIL or other service management frameworks it is important to challenge
existing implementations. This is less about the actual implementation and more often
about the problems that implementation was designed to originally solve.

### An example -- service transition

As an example ITIL talks a great deal about Service Transition -- getting working functionality
into the hands of the users of the service. This is a key topic for The Digital Service Standard
too which says that teams should:

> Make sure that you have the capacity and technical flexibility to update and
improve the service on a very frequent basis.

[GOV.UK](https://www.gov.uk) for instance made more than 100 production releases
[during its first two weeks](https://gds.blog.gov.uk/2012/11/02/regular-releases-reduce-risk/) after launch.

This high rate of change tends to challenge existing processes designed for a
slower rate of change. If you are releasing your service every month or every 6 months
then a manual process (like a weekly or monthly in-person change approval board or CAB) may be
the most suitable approach. If you're releasing many times a day then the approach to how
change is evaluated, tested and managed tends towards being more automated. This
moves effort from occasional but manual activities to upfront design and automation
work. More work is put in to assure the processes rather than putting all the effort
into assuring a specific _transition_.

Service management frameworks tend to acknowledge this, for instance ITIL has a
concept of a standard change (something commonly done, with known risks
and hence pre-approved), but a specific implementation in a given organisation
might not.

## Other frameworks exist

It is important to note that other service management frameworks and standards
exist, including some that are of a similar size and scope to ITIL:

* [ISO/IEC 20000](https://en.wikipedia.org/wiki/ISO/IEC_20000)
* [Microsoft Operations Framework](https://www.microsoft.com/MOF)

Many organisations also use smaller processes and integrate them together.
The needs of your service and organisation will determine what works best for you.

## Problematic concepts

Some traditional language tends to cause some confusion when discussing service
management alongside agile. It's generally best to avoid the following terms when possible,
although given their widespread usage this isn't always possible. It is however worth
being aware of the problems these concepts raise.

### Projects

Projects tend to imply a start and an end. The main goal of project work is to
complete it, to reach the end. Especially for software development the project can
too often be viewed as _done_ when the software is released. What happens after that
is another problem entirely -- and often someone else's problem.

However when building services the main goal is to meet [user needs](/service-manual/user-centred-design/user-needs).
These needs may change over time, and are only met by software that is running in
production and available to those users.

This doesn't mean not breaking work down into easily understandable parts, but [stories](/service-manual/agile/writing-user-stories),
[sprints](/service-manual/agile/features-of-agile#sprints) and epics are much more
suited to agile service delivery.

### Business as usual

The concept of business as usual also clashes with a model of continuous
service improvement. It immediately brings to mind before and after
states, often with the assumption that change is both much slower and
more constrained during _business as usual_. In reality, until you put your
service in the hands of real users as part of an [alpha](/service-manual/phases/alpha)
or [beta](/service-manual/phases/beta) you won't have all the information
needed to build the correct service. And even once you pass the [live standard](/service-manual/digital-by-default/maintaining-the-standard)
you will be expected to:

> continuously update and improve the service on the basis of user feedback,
performance data, changes to best practice and service demand

## Further reading

* [Functions in service management](http://www.slideshare.net/nuwulang/functions-in-service-operation)
* [ITIL on Wikipedia](https://en.wikipedia.org/wiki/ITIL)
* [Introduction to ITIL v3](http://www.best-management-practice.com/gempdf/itsmf_an_introductory_overview_of_itil_v3.pdf)
* [Discussion of DevOps and ITIL](http://blog.ingineering.it/post/59414765140/itil-vs-devops-slugfest-or-lovefest)
* [Discussion of ITIL and Continuous Delivery](http://changeandrelease.com/2014/04/05/devops-and-itil-continuous-delivery-doesnt-stop-at-software/)
