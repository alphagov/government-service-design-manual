---
layout: detailed-guidance
title: Releasing software
subtitle: How regular releases can reduce risk
category: making-software
type: guide
audience: 
  primary: web-ops, developers, tech-archs
  secondary: service-managers, designers, delivery-managers
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
    title: Making software
    url: /service-manual/making-software
---

Constantly improving online services means releasing updates to the underlying software. How often you want to do this will affect how you design and build applications, and will present a number of challenges that this guide hopes to address.

## Releasing software on demand

As early as possible in your product’s development, think about how you release updates to a running application. This is because it affects how software is developed and tested, and how your product may be supported.

Being able to release software on demand is important. Release cycles of 6 months or longer are dangerous. Not only do new features rarely see the light of day, but known problems will have to be fixed within a rigid release schedule.

It’s important to make the distinction between:

* releasing regularly
* having the ability to release all the time

Your application should always be in a state where it could be released, meaning quick changes can be made when needed. As an example, changes to the software running GOV.UK are made on average 5 times per day.

To be able to do this, you have to consider:

* [your approach to testing](/service-manual/making-software/code-testing.html)
* the quality of low-level code -- approaches like test-driven design and continuous integration (where code is tested constantly) can be helpful
* using the same tools and release processes for both the [development and production environments](/service-manual/making-software/development-environment.html) - this way the software and tools will be well understood and will have been run thousands of times before the first public launch

Although you need tools (potentially including commercial tools) to help with rapid releases, don’t start your discussions around what tools should be used or procured! Start discussions based on the needs of the service and the product team.

## Why government does this

In some organisations, people fear releasing new applications or new versions of software. Lots of websites, especially large applications within large traditional organisations, don’t change very often. Many will have fixed release schedules which might mean 1 release every 6 months or so. This means bundling up lots of changes into a single release, which is bad in at least 2 ways:

* the people using your service don’t get new features and improvements quickly
* bundling up lots of new features makes the release more complicated

It could be weeks or months before an improvement is actually released for people to use, even though it only took a few days to finish. This complexity means there are lots of different ways the release can go wrong.

The combination of complexity, risk, and the infrequent nature of releases make for a stressful event for all involved. No wonder most people don’t like release day!

## Regular releases reduce risk

Releasing software comes with risks, so trying to minimise those risks is a good idea. GDS does this in a number of ways, and the benefit of this is that:

* releasing smaller chunks regularly makes it much easier to see what’s going to change, and if something goes wrong it’s much simpler to undo
* doing something regularly makes the case for investing in automation easier, removing much of the potential for human error and making each release the same
* if you’re doing something several times a day you tend to get better at it

As well as reducing risk, being able to release early and often also helps products improve quickly, as it removes barriers to quick experiments and rapid iteration.

Finally, consider the following 2 measures of a system:

* mean time between failures
* mean time to recovery

A very traditional approach involves working to reduce the time between any failures happening, hopefully improving the quality of the overall system.

But problems will always happen at some point! Using your efforts to reduce the time taken to fix problems can often be much more cost effective, as well as improving the overall system uptime.

## Further reading

* [Regular Releases Reduce Risk](https://gds.blog.gov.uk/2012/11/02/regular-releases-reduce-risk/) blog post about the approach to releasing software onto GOV.UK
