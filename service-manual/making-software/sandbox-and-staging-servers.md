---
layout: detailed-guidance
title: Sandbox and staging servers
subtitle: Working in a development environment
category: making-software
type: guide
audience:
  primary: web-ops
  secondary: developers, qa, tech-archs
status: draft
phases:
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

Everyone working on the design, development or maintenance of your service should have a clear, easily accessible place to review the latest version of the software. Those working hands-on to build the software should be able to run their own reasonable replica of the entire service.

Make sure you have a clear staging environment where changes are reviewed and tested in the setting of the entire end-to-end service before they are deployed.

## Shared environments

Make it possible for everyone working on your service to see progress and understand their work in its full context. That means having a shared environment where:

* any team member can see the current state of the service
* stories can be signed off
* individual sandboxes are used - more experimental or early stage work can take place and be reviewed

In addition to those environments, it’s helpful to have a separate staging environment where final quality assurance and testing takes place before changes are deployed to the live/production environment. Make this identical to the production environment so it can be used effectively for performance testing.

## GOV.UK development environments

Those working on GOV.UK use:

* the [Vagrant](https://www.vagrantup.com/) tool to provide all developers with a development environment built similarly to the production environment
* a preview environment that's updated by a [Continuous Integration](https://en.wikipedia.org/wiki/Continuous_integration) system whenever tests have passed on a change
* a staging environment for review of specific changes before they go to the production environment (it's updated and reviewed as part of the release process)

## Why GOV.UK does this

We set up the development environment in this way because:

* everyone on the team should be able to understand progress to date by using running software
* people working on the service should be able to understand the effect of their work by seeing it working in context
* the team should be confident that the service as a whole works before shipping those changes to the public

##Further reading
* [Regular Releases Reduce Risk](https://gds.blog.gov.uk/2012/11/02/regular-releases-reduce-risk/)
