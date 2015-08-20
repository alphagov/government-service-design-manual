---
layout: detailed-guidance
title: DevOps
subtitle: Bringing development and operations together
category: operations
type: guide
status: draft
audience:
  primary: web-ops, developers
  secondary: service-managers, tech-archs
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
---

DevOps is a cultural and professional movement in response to the mistakes commonly made by large organisations. Often organisations will have very separate units for:

* development
* quality assurance
* operations business

In extreme cases these units may be:

* based in different locations
* work for different organisations
* under completely different management structures

Communication costs between these units, and their individual incentives, leads to slow delivery and a mountain of interconnected processes.

This is what DevOps aims to correct. It is not a methodology or framework, but a set of principles and a willingness to break down silos. Specifically DevOps is all about:

* [culture](#culture)
* [automation](#automation)
* [measurement](#measurement)
* [sharing](#sharing)

### Culture

DevOps needs a change in attitude so shared ownership and collaboration are the common working practices in building and managing a service. This culture change is especially important for established organisations.

### Automation

Many business processes are ready to be automated. Automation removes manual, error-prone tasks -- allowing people to concentrate on the quality of the service. Common areas that benefit from automation are:

* release management (releasing software)
* provisioning
* configuration management
* systems integration
* monitoring
* orchestration (the arrangement and maintenance of complex computer systems)
* testing

### Measurement

Data can be incredibly powerful for implementing change, especially when it’s used to get people from different groups involved in the quality of the end-to-end service delivery. Collecting information from different teams and being able to compare it across former silos can implement change on its own.

### Sharing

People from different backgrounds (ie development and operations) often have different, but overlapping skill sets. Sharing between groups will spread an understanding of the different areas behind a successful service, so encourage it. Resolving issues will then be more about working together and not negotiating contracts.

## Why DevOps

The quality of your service will be compromised if teams can't work together, specifically:

* those who build and test software
* those that run it in production

The root cause is often functional silos; when one group owns a specific area (say quality) it’s easy for other areas to assume that it’s no longer their concern.

This attitude is toxic, especially in areas such as:

* quality
* release management
* performance

High quality digital services need to be able to adapt quickly to user needs, and this can only happen with close collaboration between different groups.

Make sure the groups in your team:

* have a shared sense of ownership of the service
* have a shared sense of the problem
* develop a culture of making measurable improvements to how things work

## Good habits

DevOps isn’t a project management methodology, but use these good habits in your organisation. While not unique to DevOps, they help with breaking down silos when used with the above principles:

* [cross-functional teams](/service-manual/the-team) -- make sure your teams are made up of people from different functions (this helps with the team owning the end-to-end quality of service and makes it easier to break down silos)
* [widely shared metrics](/service-manual/measurement) -- it’s important for everyone to know what ‘good’ looks like so share high and low level metrics as widely as possible as it builds understanding
* automating repetitive tasks -- use software development to automate tasks across the service as it:
  * encourages a better understanding of the whole service
  * frees up smart people from doing repetitive manual tasks
* post-mortems -- issues will happen so it’s critical that everyone across different teams learns from them; running post-mortems (an analysis session after an event) with people from different groups is a great way of spreading knowledge
* [regular releases](/service-manual/making-software/release-strategies) -- the capacity for releasing software is often limited in siloed organisations, because the responsibilities of the different parts of the release are often spread out across teams -- getting to a point where you can release regularly (even many times a day) requires extreme collaboration and clever automation

## Warning signs

Like agile, the term DevOps is often used for marketing or promotional purposes. This leads to a few common usages, which aren’t necessarily in keeping with what’s been said here. Watch out for:

* DevOps tools (nearly always marketing)
* a DevOps team (in many cases this is just a new silo of skills and knowledge)
* DevOps as a job title (you wouldn’t call someone “an agile”)

## Related guides in the Service Manual

Those interested in DevOps are often also interested in:

* [Configuration Management](https://www.gov.uk/service-manual/making-software/configuration-management.html)
* [Monitoring](https://www.gov.uk/service-manual/operations/monitoring.html)
* [Release Management](https://www.gov.uk/service-manual/making-software/release-strategies.html)

## Further reading

* [What DevOps Means to Me](https://www.chef.io/blog/2010/07/16/what-devops-means-to-me/)
* [What is this DevOps thing anyway?](http://www.jedi.be/blog/2010/02/12/what-is-this-devops-thing-anyway/)
* [What is DevOps? (and the wall of confusion)](http://dev2ops.org/2010/02/what-is-devops/)
* [There’s no such thing as a "DevOps Team"](http://continuousdelivery.com/2012/10/theres-no-such-thing-as-a-devops-team/)
