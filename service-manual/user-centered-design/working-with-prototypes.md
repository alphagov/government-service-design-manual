---
layout: detailed-guidance
title: Working with prototypes
subtitle: Learning by doing
category: user-centered-design
type: guide
audience:
  primary: tech-archs
  secondary: service-managers, designers, developers, user-researchers
status: draft
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User centered design
    url: /service-manual/user-centered-design
---

The best way to understand a product is to try to build it. Prototyping lets you get the feel of a product, begin to estimate the work involved and provide something you can quickly test with real users.

## Building prototypes

This is a vital part of a process often known as '[product discovery](/service-manual/phases/)': 

* understanding your users and their needs
* developing a sense of how you might serve those needs
* estimating the effort involved in building and running a service to do so

We built alpha.gov.uk as a prototype of what would later become the single domain www.gov.uk. It was built quickly without much concern to scalability, resilience, or any of the other considerations of a 'real' product, because none of those matter unless the core concept is sound. 

This process allowed us to get feedback early and also understand some of the trickier concepts we would have to grapple with, like the fuzzy lines between different audiences and the operational processes that would be required.

## Ideas can be ugly

Prototyping can start on paper with sketches. Hand-drawn sketches of what a service might involve are a good way to begin thinking things through. 

We encourage everyone to get to running code as quickly as possible. It's only when you start working in the same medium your users will be using (for online services that's generally a web browser, but it may also be via an API (application programming interface)) that you can really understand the experience you need to provide.

The [smart answer format](https://www.gov.uk/maternity-benefits "example of a smart answer") for [GOV.UK](https://www.gov.uk) began as a series of paper sketches refined over a week by a small team. That process gave them a good sense of the boundaries of the problem they were trying to solve. 

We quickly prototyped the format using HTML and javascript so that it could be experienced in a web browser. That revealed more constraints, for example users might expect to be able to go back and amend an answer without realising that would change their whole journey through the format. This let us quickly adjust the user interface to be clearer for its users before we started the work of building the full system.

Running code forces you to think about your integrations with other services, like existing databases and whether you need to send email, and how they might work. A prototype will rarely include these integrations but having a clear picture of them is vital if you're going to understand the real effort involved in building and operating the service.

##Prototype code is not production code

It can be tempting to use code written for prototypes in production as it often uses similar technology and appears to 'work'. However, it's important to maintain separation between the 2 codebases, as they have different and opposing goals:

**Prototyping**

Fast changes, the ability to easily and quickly try many different approaches.

**Production code**

Security, maintainability, performance - in short good quality code. This takes time and excellent front-end skills.
