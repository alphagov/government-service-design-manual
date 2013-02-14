---
layout: default
title: Releasing software
section: guidance
subsection: Maintaining services
status: draft
---
    
# Releasing software
Constantly improving online services means releasing changes to the
underlying software. How often you want to do this will affect how you
design and build the applications and presents a number of challenges
that this guide hopes to address.

## Guidance/Tool

It is important to think about how you release changes to a running
application as early in the products development as possible. This is
because it affects how software is developed and tested and how a
product may be supported.

Being able to release software on demand is important. 6 monthly or
longer release cycles are dangerous. Not only do new features rarely see
the light of day but fixing known problems have to fit within a rigid
release schedule. 

Note that it's important to make the distinction between releasing
regularly and _the ability to_ release all the time. The application
should always be in a state where it could be released, that means quick
changes can be made when needed. As an example changes to the software
running GOV.UK are made on average 5 times per day.

In order to do that you have to consider:

* Your approach to testing. Manual testing is important, but also
  expensive and time consuming. Invest in automation
* The quality of low level code early. Approaches like continuous
  integration, where code is tested constantly, and test driven design,
  can be helpful
* Using the same tools and release processes for both the development and
  production environments. This way the software and tools will be well
  understood and will have been run thousands of times before the first
  public launch

Although tools, potentially incuding commercial tools, are required to
aid rapid releases the discussions should not start with what tools
should be used or procured but with the needs of the service and the product
team.


## Why we do this

In some organisations, people fear releasing new applications or new
versions of software. Lots of websites, especially large applications
within large traditional organisations, don’t change very often. Many
will have fixed release schedules which might mean one release every six
months or so. This means bundling up lots of changes into a single
release, which is bad in at least two ways:

* The people using your service don’t get new features and improvements
  quickly. It could be weeks or months before an improvement that only
  took a few days to finish is actually released for people to use
* By bundling up lots of new features you make the release more
  complicated. This complexity leads to lots of different ways the release
  can go wrong

The combination of complexity, risk and the infrequent nature of
releases makes for a stressful event for all involved. No wonder most
people don’t like release day!

Releasing software comes with risks, so trying to minimise those risks
is prudent. We do that in a number of ways:

* By releasing smaller chunks regularly it’s much easier to see what is
  going to change, and if something goes wrong it’s much simpler to roll
  that change back and undo it
* Doing something regularly makes the case for investing in automation
  easier. This helps remove much of the potential for human error and
  makes releases the same every time
* If you’re doing something several times a day you tend to get better at
  it

As well as reducing risk, being able to release early and often also
helps products improve quickly, by reducing a potential barrier to quick
experiments and rapid iteration.

Finally consider the following two measures of a system; mean time between
failures and mean time to recovery. A very traditional approach involves
focusing completely on reducing the time between any failures happening, by
hopefully improving the quality of the overall system. But problems will
always happen at some point, so focusing some effort on reducing the
time taken to fix problems that do occur can often be much more cost
effective as well as improve the overall system uptime.


## Further reading

* [Regular Releases Reduce Risk](http://digital.cabinetoffice.gov.uk/2012/11/02/regular-releases-reduce-risk/) Blog post about the approach to releasing software onto GOV.UK
