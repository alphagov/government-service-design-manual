---
layout: gsdm
title: Load and performance testing
section: guidance
subsection: Technical Architecture
status: draft
---
    
# Load and performance testing

History is litered with countless Government projects which collapsed under
load or which worked slowly enough to frustrate users.

As a Government service it is important that your systems and applications are
performant and can deal with exected (and unexpected) levels of traffic. This
means doing some capacity planning work up front but it also means doing
specific load and performance testing. 

## Guidance/Tool

### Capaicity planning

Capacity planning is the process of determining what amount of
infrastructure and software is required to run a live system.
Importantly you should also look at how this changes over time. Will
traffic or database load increase every month as the site grows, or are
some days or months (self assessment deadline or bank holidays)
particularly spikey. The output should help with both estimating ongoing
costs as well as setting up realistic load and performance tests.

### Test yourself as well as with vendors

There are several companies that offer products and services around load
testing. In many cases working with these can be useful, especially for
final testing or testing of mature components. However it's important to
also have some capability within your development team to do more ad-hoc
load and performance testing. This is important to allow for an
rapid iterative development style, otherwise scalability or performance
issues can be introduced and caught late, when they are much harder to
fix.

## Why we do this

Although related, and tested in similar ways, load testing and
performance testing are done for different reasons. It's worth
understanding the subtle difference and making sure you consider both
when testing and alalysing results.

### Load testing

We test sites and applications under realistic load to ensure that, when
launched, they work for the people using them. This should involve
testing in load in excess of expected traffic in order to simulate
certain types of Denial of Service (DoS) attack, including a Distributed
Denial of Service (DDoS) attack.

### Performance testing

Even if a site or application is able to scale out successfully it
doesn't mean it is fast. Site performance is a factor of many things,
from the software running the site to the networks, proxies and caches
involved in serving traffic over the internet. Fast sites are generally
much more effective to the extent that Google now includes performance
in it's algorithms for determining which sites to feature in search
results.

## Further reading

[Matt Cutts from Google on page speed](http://www.mattcutts.com/blog/site-speed/)

[Distributed Denial of Service attacks](http://en.wikipedia.org/wiki/Denial-of-service_attack)

[Capacity Planning Presentation](http://www.slideshare.net/jallspaw/velocity2008-capacity-management1-484676)
