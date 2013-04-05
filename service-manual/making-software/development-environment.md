---
layout: detailed-guidance
title: Development Environments
subtitle: Early infrastructure needs for agile projects
category: making-software
type: guide
audience:
  primary: designers, delivery-managers, web-ops, developers, tech-archs
  secondary: service-managers, content-designers, performance-analysts
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

The following describes the general capabilities required of any potential development environment for use by the exemplar projects to:

* test software choices to prove they are valid
* experiment quickly with new approaches
* produce and test software in a production-like architecture
* develop rapidly and iteratively
* continuously test and monitor software during development

Although this document does not describe the capabilities and characteristics of a production environment, there is a general presumption that any production environment should enable the exemplar project development teams to:

* deploy updates to the system rapidly and iteratively (ie at least daily)
* continuously test and monitor software in production


## Required

The essential capabilities of the development environment without which the development team will not be able to operate.

| 1.1 | **Current availability**: one of our key requirements is for a service that is already operational and able to onboard customers very quickly (typically within 5 working days).
| 1.2 | **Internet connectivity**: we require virtual machines with both incoming and outgoing internet connectivity. This should also facilitate remote management.
| 1.3 | **Self service provisioning**: initially we have a need for a small number of virtual machines (10 or less) but that may grow over time. We want to be able to remotely provision new machines ourselves to meet our needs as they arise, without the need to phone, fax or email anyone, and therefore require a self service method of provisioning virtual machines and storage.
| 1.4 | **Suitable range of virtual machine options**: support for 64 bit architectures and a range of virtual machine sizes at least up to 4 cores, 16GB RAM and 300GB disk
| 1.5 | **Run own operating system**: we require the flexibility to run whatever operating system is deemed suitable for the project, rather than just a limited subset of those supported by a vendor.
| 1.6 | **EU based data centres**: given the nature of the content we would prefer to store data in the EU, and ideally within the UK, therefore we require development environments to be hosted only in EU-based data centres.
| 1.7 | **Service Level Agreement**: a suitable SLA should be in place with the service provider (whether internal or external), with at least a 99.5% uptime guarantee.
| 1.8 | **Development team access**: root access to manage virtual machines (eg to install & configure software) for approved development team members.

## Desired

Optional capabilities which would make a marked difference to the production of the services.

| 2.1 | **Provisioning API**: We’d like the provisioning of virtual machines, storage, load balancing, etc to be available via an API. Any API should have a suitable authentication mechanism in place, and should be accessible to development team members via the Internet (optionally through a VPN).
| 2.2 | **Create virtual machine templates**: To speed up provisioning we would like to be able to store virtual machine templates from which new machines can be launched.
| 2.3 | **Firewall and load balancer service**: If available a managed firewall and/or load balancer service may be used.
| 2.4 | **Configurable private network**: We require the ability to manage internal networks, each consisting of specific groups of virtual machines. This should allow for some virtual machines not to be internet accessible.
| 2.5 | **Virtual Private Network**: We may choose to expose parts of the service via a Virtual Private Network. The infrastructure service should at a minimum not prevent this and may ideally provide a suitable managed service.

