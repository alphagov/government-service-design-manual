---
layout: detailed-guidance
title: Sandbox and staging servers
subtitle: Working in a development environment
section: making-software
type: guide
audience:
  primary: developer, tech-arch
status: draft
phases:
  - beta
  - live
---
    
#Sandbox and staging servers

Everyone working on design, development or maintenance of a service should have a clear, easily accessible place to review the latest version of the software. Those working hands-on building the software for the service should be able to run their own reasonable replica of the entire service. There should be a clear staging environment where changes are reviewed and tested in the context of the entire end-to-end service before they are deployed.

##Guidance/Tool
Everyone working on a service should be able to see progress and understand their work in its full context. That means a shared environment where any team member can see the current state of the service and where stories can be signed off, and individual sandboxes where more experimental or early stage work can take place and be reviewed.

In addition to those environments it can be helpful to have a separate staging environment where final quality assurance and testing can take place before changes are deployed to the live/production environment. This should be identical to the production environment so it can be used effectively for performance testing.

##Examples
For those working on GOV.UK we use the [Vagrant](http://www.vagrantup.com/) tool to provide all developers with a development environment configured similarly to the production environment. We then have a preview environment that is updated by our Continuous Integration system whenever tests have passed on a change. There is then a staging environment for review of specific changes before they go to the production environment. It is updated and reviewed as part of the release process.

##Why we do this
* Everyone on the team should be able to understand progress to date by using running software
* People working on the service should be able to understand the impact of their work by seeing it working in context
* The team should be confident that the service as a whole works before shipping those changes to the public

##Further reading
* [Regular Releases Reduce Risk](http://digital.cabinetoffice.gov.uk/2012/11/02/regular-releases-reduce-risk/)
