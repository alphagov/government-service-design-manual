---
layout: detailed-guidance
title: Development Environments
subtitle: Early infrastructure needs for agile projects
section: making-software
type: guide
audience:
  primary: developer, tech-arch
status: draft
phases:
  - alpha
  - beta
  - live
---
    
# Development environments

The following describes the general capabilities required of any potential development environment for use by the exemplar projects to:

* Test software choices to prove they are valid
* Experiment quickly with new approaches
* Produce and test software in a production-like architecture
* Develop rapidly and iteratively
* Continuously test and monitor software during development

Although this document does not describe the capabilities and characteristics of a production environment, there is a general presumption that any production environment should enable the exemplar project development teams to:

* Deploy updates to the system rapidly and iteratively (ie at least daily)
* Continuously test and monitor software in production

_Required_ - describes the essential capabilities of the development environment without which the development team will not be able to operate.
_Desired_ - describes optional capabilities which would make a marked difference to the production of the services.

## Required

| 1.1 | Current availability  One of our key requirements is for a service that is already operational and able to onboard customers very quickly (typically within 5 working days).
| 1.2 | Internet connectivity We require virtual machines with both incoming and outgoing internet connectivity. This should also facilitate remote management.
| 1.3 | Self service provisioning Initially we have a need for a small number of virtual machines (10 or less) but that may grow over time. We want to be able to remotely provision new machines ourselves to meet our needs as they arise, without the need to phone, fax or email anyone, and therefore require a self service method of provisioning virtual machines and storage.
| 1.4 | Suitable range of virtual machine options Support for 64 bit architectures and a range of virtual machine sizes at least up to 4 cores, 16GB RAM and 300GB disk
| 1.5 | Run own operating system  We require the flexibility to run whatever operating system is deemed suitable for the project, rather than just a limited subset of those supported by a vendor.
| 1.6 | EU based data centres Given the nature of the content we would prefer to store data in the EU, and ideally within the UK, therefore we require development environments to be hosted only in EU-based data centres.
| 1.7 | Service Level Agreement A suitable SLA should be in place with the service provider (whether internal or external), with at least a 99.5% uptime guarantee.
| 1.8 | Development team access Root access to manage virtual machines (eg to install & configure software) for approved development team members.

## Desired

| 2.1 | Provisioning API  We’d like the provisioning of virtual machines, storage, load balancing, etc to be available via an API. Any API should have a suitable authentication mechanism in place, and should be accessible to development team members via the Internet (optionally through a VPN).
| 2.2 | Create virtual machine templates  To speed up provisioning we would like to be able to store virtual machine templates from which new machines can be launched.
| 2.3 | Firewall and load balancer service  If available a managed firewall and/or load balancer service may be used.
| 2.4 | Configurable private network  We require the ability to manage internal networks, each consisting of specific groups of virtual machines. This should allow for some virtual machines not to be internet accessible.
| 2.5 | Virtual Private Network We may choose to expose parts of the service via a Virtual Private Network. The infrastructure service should at a minimum not prevent this and may ideally provide a suitable managed service.

