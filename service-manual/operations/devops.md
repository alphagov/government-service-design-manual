---
layout: detailed-guidance
title: Devops
subtitle: Bringing development and operations together
category: operations
type: guide
status: draft
audience:
  primary: web-ops
  secondary: service-managers, developers, tech-archs
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
---

Many large organisations have evolved to have very separate development, quality assurance and operations business units. In many cases the overhead of communications and the misaligned incentives of the different groups leads to slow delivery and a mountain of interconnected processes. In extreme cases these units may be geographically separated, work for different organisations or under completely different management structures.

Devops is a cultural and professional movement which has grown out of these frustrations. It is not a methodology or framework, rather some high minded principles and a willingness to break down silos. Specifically devops is all about [culture](#culture), [automation](#automation), [measurement](#measurement) and [sharing](#sharing).

### Culture

It is important to realise that devops requires a culture shift, towards shared ownership and collaboration across silos involved in building and managing a service. Especially for existing organisations this is by far the most important element.

### Automation

Many business processes are ripe for automation. Automation removes manual, error prone tasks allowing people to focus on the quality of the running service. Common areas that could benefit from automation are release management, provisioning, configuration management, systems integration, monitoring, orchestration and testing.

### Measurement

Using data to drive change can be hugely powerful, in particular when it is done to involve people from different groups in the quality of the end-to-end service delivery. Collecting information from different teams and being able to compare it across former silos can be transformative on its own.

### Sharing

People from different backgrounds (say development and operations) often have different, but intersecting skill sets. Encouraging sharing between groups helps to spread an understanding of the different areas of a successful service, making resolving issues about working together and not about contract negotiation.

## Why we do this

Without close collaboration between those building and testing software and those running it in production, the quality of the resulting service will be compromised. This happens in many ways, but the root cause is often functional silos; when one group owns a specific area (say quality) it’s easy for other areas to assume that area is no longer their concern. In areas like quality, release management or performance this is toxic.

High quality digital services need to be able to adapt quickly to user needs, and this means close collaboration between different groups. A shared sense of ownership of the service and the problem is also important, as is a culture of making measurable improvements to how things work.

## Good habits

Devops isn’t a project management methodology but you will hopefully adopt some of these good habits in your organisation. These aren’t unique to devops, but they do help with breaking down silos and with the above principles.

* [Cross-functional teams](/service-manual/the-team) – working to break down silos is easiest done when teams are composed of people from different functions. This helps with the team owning the end-to-end quality of service.
* [Widely shared metrics](/service-manual/measurement) – knowing what good looks like across the service is important for everyone, and sharing high and low level metrics as widely as possible is a good way of building understanding.
* Automating repetitive tasks – using software development to automate tasks from across the service is a good way of fostering a better understanding of the whole service, as well as freeing up smart people from repetitive manual tasks.
* Post-mortems – issues will happen and making sure everyone from across different teams learns from them is critical. Running post-mortems with people from different groups is a great way of spreading knowledge.
* [Regular releases](/service-manual/making-software/release-strategies.html) – releasing software is often a bottleneck in siloed organisations, as the responsibilities of the different parts of the release are often spread out across teams. Getting to a point where you can release regularly (even many times a day) requires extreme collaboration and clever automation.

## Warning signs

Like agile, the term devops is often used for marketing or promotional purposes. This leads to a few common usages which are not necessarily in keeping with the above discussion. Watch out for:

* devops tools (nearly always marketing)
* a devops team (in many cases this is just a new silo of skills and knowledge)
* devops as a job title (you wouldn’t call someone “an agile”)

## Further reading

* [What devops means to me](http://www.opscode.com/blog/2010/07/16/what-devops-means-to-me/)
* [What is this devops thing anyway?](http://www.jedi.be/blog/2010/02/12/what-is-this-devops-thing-anyway/)
* [What is devops (and the wall of confusion)](http://dev2ops.org/2010/02/what-is-devops/)
* [There’s No Such Thing as a "Devops Team"](http://continuousdelivery.com/2012/10/theres-no-such-thing-as-a-devops-team/)
* Those interested in devops are often also interested in [Configuration Management](https://www.gov.uk/service-manual/making-software/configuration-management.html), [Monitoring](https://www.gov.uk/service-manual/operations/monitoring.html) and [Release Management](https://www.gov.uk/service-manual/making-software/release-strategies.html)
