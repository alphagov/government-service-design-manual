---
layout: detailed-guidance
title: Development environments
subtitle: Early infrastructure needs for agile projects
category: making-software
type: guide
audience:
  primary: designers, delivery-managers, web-ops, developers
  secondary: service-managers, content-designers, performance-analysts, tech-archs
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
exclude_from_search: true
---

As software developers, the environments we use every day matter greatly. Below are a set of guidelines for development environments to enable the [exemplar projects][] (service transformations committed to in the [Government Digital Strategy](/government/publications/government-digital-strategy)) to:

* test software choices to prove they are valid
* experiment quickly with new approaches
* produce and test software in a production-like architecture
* develop rapidly and iteratively
* continuously test and monitor software during development

Although this document doesn't describe the capabilities and characteristics of a production environment, there's a general presumption that any production environment should enable the exemplar project development teams to:

* deploy updates to the system rapidly and iteratively (ie at least daily)
* continuously test and monitor software in production

[exemplar projects]: https://www.gov.uk/transformation

## Required

The essential capabilities of the development environment without which the development team will not be able to operate.

### Current availability

A service that's already operational and able to onboard customers very quickly (typically within 5 working days).

### Internet connectivity

Both incoming and outgoing internet connectivity. This should also help remote management.

### Self-service provisioning

We should be able to remotely provision new machines ourselves to meet our needs as they arise, without the need to phone, fax or email anyone, and therefore require a self-service method of provisioning virtual machines and storage.

### Suitable range of virtual machine options

Support for 64 bit architectures and a range of virtual machine sizes at least up to 4 cores, 14GB RAM and 300GB disk.

### EU-based data centres

We would prefer to store data in the EU, and ideally within the UK, therefore we require development environments to be hosted only in EU-based data centres.

### Service Level Agreement (SLA)

A suitable SLA should be in place with the service provider (whether internal or external), with at least a 99.5% uptime guarantee.

### Development team access

Approved development team members should have access to manage virtual machines (eg to install and configure software).

## Desired

Optional capabilities which would make a marked difference to the production of the services.

### Provisioning API

The provisioning of virtual machines, storage, load balancing etc to be available via an application programming interface (API). Any API should have a suitable authentication mechanism in place, and should be accessible to development team members via the internet (optionally through a virtual private network (VPN)).

### Run own operating system

The flexibility to run whatever operating system is deemed suitable for the project, rather than just a limited subset of those supported by a vendor.

### Create virtual machine templates

To speed up provisioning we would like to be able to store virtual machine templates from which new machines can be launched.

### Firewall and load balancer service

If available a managed firewall and/or [load balancer](https://en.wikipedia.org/wiki/Load_balancing_(computing)) service may be used.

### Configurable private network

We require the ability to manage internal networks, each consisting of specific groups of virtual machines. This should allow for some virtual machines to not be internet accessible.

### Virtual Private Network

We may choose to expose parts of the service via a VPN. The infrastructure service should at a minimum not prevent this and may ideally provide a suitable managed service.

*[SLA]: Service Level Agreement
*[API]: application programming interface
*[VPN]: virtual private network
