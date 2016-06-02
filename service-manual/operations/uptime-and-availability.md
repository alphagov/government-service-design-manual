---
layout: detailed-guidance
title: Uptime and availability
subtitle: Keeping the government online
category: operations
type: guide
status: draft
audience:
  primary: web-ops, tech-archs
  secondary: service-managers
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
exclude_from_search: true
---


# Introduction

The expectation when using an online service is that it's available 24
hours a day 365 days a year. This is one of the advantages of shifting
services online, people can transact with you when it suits them, not
just when your office or call centre is open. But achieving good uptime
for a service takes planning and good design.

# Not just on or off

It is important to build services that wherever possible can be turned
off (or fail) gracefully. The possible states for a service should not be
limited to on or off. Individual components could be designed to fall
back to minimal functions. A read-only mode could be introduced where
information can be looked at but not changed. If your service is reliant
on a third-party service, you could queue information for later
processing if that service were to become unavailable.

Remember to build in redundancy and avoid single points of failure. You
can avoid the problem of a web server failing by having more than one
and load-balancing between them, or minimise disruption if a database
crashes by using database systems which spread data and queries across a
cluster.

# People

A critical part of maintaining a service is the people running it
day-to-day. People may be involved in responding to customer issues or
in more technical activities like resolving minor outages before they
cause a real problem.

## Weekends and evenings

Don't forget that people will be using your service when you're not at
work. Failures during out-of-hours periods can go unnoticed for long
periods of time if no-one is responsible for them. Make sure you
consider evenings and weekends and also whether an on-call system is
appropriate, or whether you need a dedicated 24/7 support capability.

## Trade-offs

Ultimately you may make trade-offs, eg sacrificing certainty around uptime
for lower cost. Try to do this openly and communicate
decisions made. Many problems regarding uptime come from mismanaged
expectations.

# Things to watch out for

## Underlying infrastructure availability

Remember the availability of the service is dependent on the
availability of lots of systems, potentially across multiple suppliers
not all of whom you have a direct relationship with. Maybe your
application (maintained by your development team) relies on an
application server or database (provided by another team in your
department) which runs on an Infrastructure as a Service platform
(provided via contract with a commercial supplier) which relies on
network connectivity and power from utilities (which you have no direct
contract with). This can become quite complex in a service-oriented
architecture where you rely on other software applications and
interfaces.

Understand your dependencies and understand their dependencies and all
the intended uptime expectations.

## Scheduled maintenance

Sometimes uptime can be quoted to ignore pre-arranged maintenance
periods. So for example a service may claim 100% even though it shuts
down every Monday evening for maintenance. Hiding uptime problems behind
lots of maintenance periods should be avoided. However, it may be fine
to quote both planned and unplanned downtime where services genuinely do
need scheduled maintenance.

## Penalties

Some suppliers will offer recompense for missing uptime guarantees or
service level agreement response times. Although this can obviously be
useful consider whether the money or service credits offered really
offset the difficulties faced by people using the service. If you're
regularly getting credits for uptime problems consider whether you are
really getting the offered uptime or SLA from your supplier.

## Contracts

The impact that contracts with suppliers (of products or services) can
have on availability should not be underestimated. Comprehensive
guidance on contracts is out of the scope of this guide but ensure
contract terms, service level agreements and uptime guarantees and
understood by the team responsible for the service. And that at the very
least the service manager is involved with contracts for underlying
systems.
