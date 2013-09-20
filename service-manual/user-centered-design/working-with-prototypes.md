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

The best way to understand a product is to try to build it. Prototyping is an essential process to get a feel for the shape and edges of a product, to begin to estimate the work involved and to provide something you can quickly test with real users.

## Building prototypes

This is a vital part of a process often known as '[product discovery](/service-manual/phases/)': understanding your users and their needs, developing a sense of how you might serve those needs, and estimating the effort involved in building and running a service to do so.

We built alpha.gov.uk as a prototype of what would later become the single domain www.gov.uk. It was built quickly without much concern to scalability, resilience, or any of the other considerations of a 'real' product, because none of those matter unless the core concept is sound. That allowed us to get feedback early and also understand some of the trickier concepts we would have to grapple with such as the fuzzy lines between different audiences, the operational processes that would be required, and so on.

## Ideas can be ugly

Prototyping can start on paper with sketches. Hand-drawn sketches of what a service might involve are a good way to begin thinking things through. We encourage everyone to get to running code as quickly as possible. It's only when you start working in the same medium your users will be using (for online services that's generally a web browser, but it may also be via an API) that you can really understand the experience you need to provide.

The [smart answer format](https://www.gov.uk/maternity-benefits "example of a smart answer") for [GOV.UK](https://www.gov.uk) began as a series of paper sketches refined over a week by a small team. That process gave them a good sense of the boundaries of the problem they were trying to solve. As quickly as possible we prototyped the format using HTML and javascript so that it could be experienced in a web browser. That revealed more constraints, such as the fact that users might expect to be able to go back and amend an answer without realising that would change their whole journey through the format. This allowed us to quickly adjust the user interface to be clearer for its users before we started the work of building out the full system.

Running code also forces you to think about your integrations with other services and how they might work – do you need to send email? integrate with an existing database? and so on. A prototype will rarely actually include these integrations but having a clear picture of them is vital if you're going to understand the real effort involved in building and operating the service.
