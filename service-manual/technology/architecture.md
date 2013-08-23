---
layout: detailed-guidance
title: Technology architecture
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

Architecture is not just about the design and deployment of technology - to be successful you will need to understand how several elements work together.

These include:

* the high-level **business** vision and requirements that guide the design of processes and systems around the needs of users (recognise that business requirements *will* change over time, so design solutions that minimise the impact of such changes)
* an **information** model covering both public and private data, as well as wider information management requirements like business intelligence
* **technology**, the modelling of core capabilities, their structure and relationships
* **implementation**, including the development and application of guidelines, patterns, blueprints, models and communities that ensure consistent implementation and compliance across multiple programmes and projects
* throughout the project lifecycle you will need to ensure **stakeholders** stay aligned on what is needed (you will need to resolve issues when stakeholder requirements conflict with business drivers (like reduced cost), system design (like network latency) and operational issues (like security, privacy and auditing))

The term “architecture” includes both the **logical** design as well as its **physical** implementation. Logical architecture based on [user-centric service design](https://www.gov.uk/service-manual/start#a-new-way-of-doing-things) and clear [user needs](/service-manual/user-centered-design/user-needs.html) should always precede physical architecture - at a minimum you need to understand [what capabilities are required](/service-manual/making-software/choosing-technology.html#start-with-capabilities-not-implementations) before starting to narrow down on product choices.

The [technology code of practice](/service-manual/technology/code-of-practice.html) provides guidance on the consistent use of technology across government.

##Our reference model
Our reference model draws upon modern [platform-based architectures](/service-manual/technology/government-as-a-platform.html). We should use the highly scalable open platform model of internet-scale organisations and draw upon the practical experience of other governments like [Estonia](http://e-estonia.com/components/x-road).

<img src="/service-manual/assets/images/architecture-reference-model.png" alt="Diagram showing government architecture reference model" />

Note: examples are for illustration purposes only.

To support this our architecture will need to move away from tightly-coupled and proprietary models towards an open, interoperable architecture using [open standards](/service-manual/making-software/open-standards-and-licensing.html) and [open interfaces](/service-manual/making-software/apis.html), eg XML- and JSON-based data services, to maximise government’s flexibility and improve the speed of delivery of services.

We should strive to:

* avoid duplication of data ownership – to prevent synchronisation, data quality, security, and privacy problems
* limit duplication of development effort – to prevent the continual reinvention of the wheel across multiple projects dealing with common data, users and processes. Wherever possible (and tested against the “[rule of three](http://www.maheshpai.info/?p=20)”), we will use and reuse standard, interoperable components, interfaces, data formats and processes using open standards

Duplication is OK when used to optimise the user experience (but there should be clear master-slave data relationships should information itself be replicated, eg for reasons of performance).

##Products and Services

Products and Services are the things that users engage with to find information or complete transactions. As such, they should be designed around the needs of the end user rather than the service provider. They can be either Mission IT (internal facing) or Digital Public Services (external facing).

| Products Are: | Products Are NOT: |
|-----|--------|
| designed and developed in a user specific way to meet: a) a set of user needs, with success measured by the extent to which those needs are met, and b) a policy outcome | an elaboration of an internal organisational structure |
| subject to continuing assessment and improvement to further refine the quality of the service: “Am I meeting the user need?” | a product in the narrow sense of a specific application/box product, ie not Microsoft Dynamics, or Oracle e-business suite, or DropBox |
| a single point of user engagement, driven by user needs | fixed in time, but designed to evolve to meet changing needs |
| the responsibility of a [service manager](/service-manual/service-managers) (the user representative), not necessarily the owning department: reporting lines should be based around product and user need (with all the implications for governance that entails) | |
| guided by clear product roadmaps and timescales based on user need and policy | |

##Platforms

Platforms provide a set of open interfaces, protocols, data formats and tools that enable software developers to rapidly deliver products and services.

Examples of platforms are the GOV.UK publishing platform, the performance platform, and the [identity assurance](/service-manual/identity-assurance) platform - they are not things like ‘Windows Server’ or ‘Linux Server’.

Platform should be kept as simple as possible, but designed to have no barriers to scaling up and scaling out. New features should be added only when required by new product needs - platforms emerge from the demands of products rather than being driven by top-down architectural process.

More on government platforms can be found [here](/service-manual/technology/government-as-a-platform.html).

## The Legacy Estate

“Legacy” is used as shorthand to reference existing mission IT systems and components that have been designed to meet historic requirements. Often they are built using technologies that have become difficult and expensive to maintain or change.

Examples of legacy applications are likely to span benefits systems, case management systems, and customer relationship management systems. Over time the dependence on legacy platforms should be reduced as successful new platforms are developed.

To achieve this:

* new user services must take the opportunity to reduce dependency on legacy systems, and new programmes should where possible avoid nugatory expenditure on integrating with such systems
* the volume of such systems, their cost of support and business criticality should continually decrease, ie they exist to fulfil a current specific need, but should not be the focus of further investment
* departments should baseline and plan to migrate/mitigate existing systems (triage the ‘old and broke’, ‘old and works’ and ‘defunct’), ensuring value is released:
	* badly engineered systems should be contained
	* the working legacy should be modernised to reduce costs and risks and wrapped where appropriate to provide open interfaces and interoperability with platforms
	* legacy applications should be replaced / rearchitected periodically to prevent situations where the application can no longer be supported because it has grown too complex for anyone to understand or too expensive relative to the benefits delivered
	* open standards must be used for interoperability between systems and departments, with adaptors used to integrate existing host systems, programs, messages and data (from mainframes to line of business services, where those systems do not natively support open standards). You should avoid point-to-point bespoke integration between systems to reduce costs and inter-dependencies