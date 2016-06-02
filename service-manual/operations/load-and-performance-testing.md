---
layout: detailed-guidance
title: Load and performance testing
subtitle: How to avoid your service collapsing under pressure
category: operations
type: guide
audience:
  primary: web-ops
  secondary: service-managers, developers, tech-archs, qa
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
exclude_from_search: true
---

History is littered with countless government projects that collapsed under load or worked so slowly it frustrated users.

As a government service, it’s important that your systems and applications are highly efficient and can deal with expected (and unexpected) levels of traffic.

This means you need to do:

* capacity planning work at the beginning
* specific load and performance testing

## Capacity planning

Capacity planning is the process of determining what amount of infrastructure and software is required to run a live system.

You also need to look at how this changes over time:

* will traffic or database load increase every month as the site grows
* are some days or months particularly spikey (self-assessment deadline or bank holidays)

This planning will help with both estimating your ongoing costs as well as setting up realistic load and performance tests.

## Test things yourself, as well as with vendors

Several companies offer products and services around load testing. In many cases it’s useful for you to work with these companies, especially for final testing or the testing of finished components.

Still, make sure you also have the capability within your development team to do more specific load and performance testing. Having this makes it possible to have a [rapid iterative development style](/service-manual/agile). Without it, your system can have scalability or performance issues that will only be caught later on -- when they are much harder to fix.

## Types of testing

To make sure your system is efficient and able to deal with high levels or traffic, carry out:

* performance testing
* load testing

Although related and tested in similar ways, load testing and performance testing are done for different reasons.

It’s useful for you to understand the subtle differences. Make sure you consider both types when testing and analysing results.

### Load testing

This is when a site or application is given the most it can handle, while still being able to function properly.

GDS test sites and applications under realistic load (traffic) to make sure that sites and applications work for the people using them.

Load testing should involve testing load in excess of your expected traffic levels. Do this so you can simulate certain types of [Denial of Service](https://en.wikipedia.org/wiki/Denial-of-service_attack) (DOS) attack, including a [Distributed Denial of Service](https://en.wikipedia.org/wiki/Denial-of-service_attack#Distributed_attack) (DDOS) attack.

### Performance testing

This is about testing stability and responsiveness.

Even if a site or application is able to scale out successfully it doesn’t mean it’s fast. Site performance is dependent on many factors like:

* the software running the site
* the networks, proxies and caches involved in serving traffic over the internet

## Further reading

* [Matt Cutts from Google on page speed](https://www.mattcutts.com/blog/site-speed/)
* [distributed Denial of Service attacks](https://en.wikipedia.org/wiki/Denial-of-service_attack)
* [capacity planning presentation](http://www.slideshare.net/jallspaw/velocity2008-capacity-management1-484676)
