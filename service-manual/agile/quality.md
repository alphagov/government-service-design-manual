---
layout: detailed-guidance
title: Quality
subtitle: How to define it, measure it and maintain it
category: agile
type: guide
audience:
  primary: service-managers, tech-archs, designers, developers, performance-analysts
  secondary:
status: draft
phases:
  - discovery
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Agile
    url: /service-manual/agile
exclude_from_search: true
---

Put quality at the forefront of your project if you want people to use and enjoy your service.

Quality is the responsibility of every member of your team -- the people who create the service create the quality of it. Every person on your project should be part of increasing quality and fixing issues.

## Defining quality

Quality will mean different things to different members of the team. Fundamentally it is about the experience of the service's users, from the start to the end of the transaction, which may include:

* accessibility for (the widest possible range of) users, across an appropriate range of devices
* how simple it is to interact with the service
* the ability of the service provider to quickly provide appropriate support
* whether the storage and handling of any data is suitable and secure
* the strength and security of essential software and the infrastructure
* making sure the [code is well-tested](/service-manual/making-software/code-testing.html)
* seeing if it responds quickly when dealing with lots of users at once ([load testing](/service-manual/operations/load-and-performance-testing.html))
* if it can be scaled rapidly to handle unusual amounts of traffic
* the ability of the team to quickly add or change features to respond to changing requirements or settings

### A note on technical debt

‘Technical debt’ is a common term in software development, and definitions vary. Generally, technical debt is used to describe compromises made to complete a product’s development, instead of creating a clean, expandable technical design.
These compromises mean there’s still work to be done, even though the overall product is finished.

It’s impossible to develop your software without building up some technical debt, so make sure your team have some way of sharing an understanding of the technical debt in their system.

A large amount of technical debt will slow down future development.  Know how much of it you have and you can prioritise work to reduce it, and can continue to iterate rapidly in the future.

## Everyone is responsible for quality

Service quality isn’t just a testing issue or the responsibility of a few individuals. The quality of a system is defined by the people who create it. 

Your team should be able to see any problems in the quality of your system. Every person on the project should be taking action to increase quality and fix issues.

## Testing

You won’t know how great your product is or if it meets your criteria without testing it -- in both normal and unusual conditions.

Dylan Robert's book ["Learning From First Responders"](http://www.oreilly.com/webops-perf/free/learning-from-first-responders.csp) is worth reviewing as an example of testing software and teams under unusual stresses (in this case, the final stages of the 2012 presidential campaign).

This manual gives further information on testing:

* [testing in an agile environment](/service-manual/making-software/testing-in-agile.html)
* [accessibility testing](/service-manual/user-centred-design/user-research/accessibility-testing.html)
* [testing code](/service-manual/making-software/code-testing.html)
* [load and performance testing](/service-manual/operations/load-and-performance-testing.html)
* [vulnerability and penetration testing](/service-manual/operations/penetration-testing.html)
* [exploratory testing](/service-manual/making-software/exploratory-testing.html)

## Team roles and quality assurance specialists

The quality of any digital service is the responsibility of the entire team, but the final responsibility lies with the service manager.

Make sure your service manager works with the rest of your team. Otherwise you risk your team not understanding what measures they need to put in place to guarantee quality.

Your team will need to consider quality when writing user stories, and allow time and resources to be used for:

* testing what is being built
* doing the basics, so assumptions about accessibility can be regularly tested
* taking the time to consider failure scenarios and how to respond

Using specialist skills and facilities is a good way to make sure testing is thorough, and is typically used for:

* [penetration testing](/service-manual/operations/penetration-testing.html) - having people outside of the team carry out this task helps challenge assumptions and identify weaknesses, giving a valuable outside perspective
* hiring quality assurance specialists -- they can make sure all quality-related activities are coordinated, and that your team gets the training and resources needed to make a high-quality service by:
    * having a clear remit and working with your team to build quality into everything they do, instead of just adding a gate review to the development process
    * being a short-term role, leaving your team capable of managing quality as part of their standard development and iteration of your service

