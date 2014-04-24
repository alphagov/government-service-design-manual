---
layout: detailed-guidance
title: Government as a platform
subtitle: Services built on a shared core
category: technology
type: guide
audience:
  primary: chief-technology-officers
status: draft
phases:
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Chief technology officer
    url: /service-manual/technology
---

The government is implementing a [platform-based](/government/publications/case-study-on-action-11-build-government-as-a-platform) operating model. [Google](http://research.google.com/pubs/SoftwareSystems.html), [Amazon](http://www.infoq.com/presentations/vogels-amazon-platform), [Twitter](https://blog.twitter.com/2010/tech-behind-new-twittercom) and [Facebook](https://www.facebook.com/Engineering), among many others, have all built their success on the back of platforms. They have developed a core technology infrastructure that others have then built upon, driving the success of the platform and meeting far more users’ needs than the original provider could have done on their own.

## Government as a platform

One simple analogy for a platform is the Lego brick, which provides the basis for constructing many different products from the same components. The [GOV.UK publishing platform](https://gds.blog.gov.uk/govuk-launch-colophon/) provides the basis for many different products and services to be provided to citizens and businesses.

Businesses like Amazon iterate and improve their services based on data, not intuition or guesswork. Their services are built on common technology platforms (search, recommendations, purchase, etc). They deliver a scalable, consistently high-quality user experience using open source and open standards. Their world looks something like this:

<img src="/service-manual/assets/images/platform-diagram.png" alt="Diagram showing a simple version of Amazon's platform diagram" />

While government can learn from this model, it cannot simplistically copy it. Government has many diverse and complex legacy systems and processes, almost all of them operating in silos. As a consequence, government currently looks something like this:

<img src="/service-manual/assets/images/current-government-systems.png" alt="Diagram showing a simplified version of the current state of government systems" />

This move to a platform-based model is a significant transition. A platform provides essential technology infrastructure, including core applications that demonstrate the potential of the platform. Other organisations and developers can use the platform to innovate and build upon. The core platform provider enforces “rules of the road” (such as the open technical standards and processes to be used) to ensure consistency, and that applications based on the platform will work well together.

## Moving to the platform model

Moving to a platform model does not imply a “top-down” conceptual identification of a list of platforms that government will buy or build. Platforms will emerge out of [user needs](/service-manual/user-centred-design/user-needs.html) and [good architectural practices](/service-manual/technology/architecture.html), but only as a consequence of the practical experience of building the features required.

Platforms are defined by connectors and functionality -- characterised by open interfaces, open data standards, and commonality.

The figures below illustrate the journey required:

<img src="/service-manual/assets/images/platform-based-government.png" alt="Diagram showing the move towards platform-based government" />

<img src="/service-manual/assets/images/service-redesign.png" alt="Diagram showing service redesign using the platform model" />

## Working with organisations outside of government

A move to platforms does not mean that government has to develop everything in-house: many of government’s needs can be met by existing cost-efficient utility services. However, government can help to establish best practice in areas such as personal data privacy.

Wherever appropriate, the government should use existing external platforms, such as payments services (ranging from third party merchant acquirer services to the UK’s national payments infrastructure). Deciding to develop platforms in-house will happen only where that is the best way to meet users’ needs in the most flexible and cost-effective way.

Government will draw upon the experience of other organisations that have already made this journey. Many have created platforms that initially sit across the top of existing silos to expedite agile and effective digital service delivery.

This co-existence enables the benefits of the platform model to be realised quickly, even in a brownfield environment such as government, while the silos below the waterline are gradually reduced in importance and eventually made obsolete.

This is the approach that government will follow, ensuring that it develops a well-defined schedule for switching off legacy environments as the platform model is progressively implemented.

Further reading:

* [Tim O’Reilly’s “Government as a Platform”](http://ofps.oreilly.com/titles/9780596804350/defining_government_2_0_lessons_learned_.html)
* [Agile Architecture](http://www.agilemodeling.com/essays/agileArchitecture.htm)
