---
layout: detailed-guidance
title: Ensuring Quality
subtitle: So good that people prefer to use them
category: agile
type: guide
audience:
  primary: service-managers, designers, developers, performance-analysts, user-researchers, content-designers
  secondary: 
status: draft
phases:
  - discovery
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Agile
    url: /service-manual/agile
---

Quality is at the heart of building services that are so good that people prefer to use them. It is the responsibility of every member of the team and the quality of a system is defined by the people who create it.

If the system being produced is lacking in quality, then it should be evident to everyone involved, and every person on the project should be taking action to increase quality and fix issues.

## Defining Quality

Quality will mean different things to different members of the team, fundamentally it is about the end-to-end experience of the service's users, from initial interaction to completion of the transaction.

That may include:

* the accessibility of the service to the broadest possible range of users across an appropriate range of devices
* the simplicity of the interactions involved
* the ability of the service provider to quickly provide appropriate support
* whether the storage and handling of any data is proportionate and secure
* the robustness of the underlying software and infrastructure
 * is the code well tested and largely bug-free?
 * does it respond quickly when placed under reasonable amounts of load?
 * can it be scaled rapidly to handle unusual amounts of traffic
* the ability of the team to quickly add or modify features to respond to changing requirements or contexts


### A note on "Technical Debt"

It is common to talk of "technical debt" in software. Definitions of technical debt vary, but generally it means compromises made in the development of an application or system favouring speed of delivery over a clean, extensible technical design.
It is impossible to develop software without accruing some technical debt, but you should ensure your team have some way of sharing an understanding of the technical debt in their system. A large amount of technical debt slows down future development and having a clear picture will let you prioritise work to reduce it and thereby ensure your ability to iterate rapidly in future.

## Testing

In order to ensure quality a range of testing must be carried out. You can only know if your service meets the above criteria and any other definition of quality if you have tested it under normal and unusual conditions. For an example of testing under unusual stresses, it is worth reviewing Dylan Robert's book ["Learning From First Responders"](http://oreil.ly/163ybtz) about the Obama 2012 campaign's preparations to ensure their software was highly available during the final days of that Presidential campaign.

For further guidance on testing, see the following sections of this manual:
* [Testing in an agile environment](https://www.gov.uk/service-manual/making-software/testing-in-agile.html)
* [Accessibility testing](https://www.gov.uk/service-manual/making-software/accessibility-testing.html)
* [Testing code](https://www.gov.uk/service-manual/making-software/code-testing.html)
* [Load and performance testing](https://www.gov.uk/service-manual/operations/load-and-performance-testing.html)
* [Vulnerability and penetration testing](https://www.gov.uk/service-manual/operations/penetration-testing.html)

## Team Roles and QA Specialists

The quality of any digital service is the responsibility of the entire team, but the final responsibility lies with the Service Manager. It is essential that the Service Manager works with the team to understand the measures that they need to put in place to ensure quality.

That will include making sure that the team are considering quality when writing development stories and allowing time and resources to test what they're building, doing the groundwork to ensure that assumptions about the accessibility of a service are regularly tested, taking the time to consider failure scenarios and how they will respond, and so on.

At times there will be a need for specialist skills and facilities to ensure good testing. For example, there should always be some [Penetration testing](https://www.gov.uk/service-manual/operations/penetration-testing.html) undertaken by people outside the team as an outside perspective helps challenge assumptions and identify weaknesses the team are not in a position to see clearly.

Some teams may seek to employ a Quality Assurance specialist. That role can be helpful to ensure all quality-related activities are co-ordinated and to make sure that the team is getting the training and resources they need to make a high quality service.

A QA specialist should have a clear remit and be expected to work with the team to build quality into everything they do rather than simply adding a review gate to the development process. We encourage teams employing a QA specialist to consider such a role a short-term responsibility with the intention that they quickly leave the team able to manage quality as part of their standard development and iteration of the service.