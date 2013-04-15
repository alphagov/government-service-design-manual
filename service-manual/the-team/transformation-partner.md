---
layout: detailed-guidance
title: Selecting a Transformation Partner
subtitle: How to go about choosing a digital transformation partner
category: the-team
audience: service-managers
phases:
  - alpha
  - beta
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: The Team
    url: /service-manual/the-team
---

Many departments will find it useful to engage with a partner who can help build digital capability within the department at the same time as working with them to deliver working digital services. In selecting a supplier as a digital transformation partner for the department, these are some of the things that we'd like to see.

This is very much a superset of the things that we'll want to see from a partner. Not all departments, and not all projects, will need all of this. For instance, many projects will involve interfacing with and/or migrating from legacy systems, whereas some purely green-field projects won't need to do this. Similarly, many projects will require the partners to take responsibility for both delivery and enablement, but some will require only one or the other.

One important distinction to bear in mind is that with a transformation partner, we are buying a capability, not a thing.

## Track Record of Delivery

The partner needs a proven track record of successful delivery, and will take responsibility for the delivery of the digital service.

## Enablement

A major aim of the project should be the creation of an agile delivery capability within the department. The partner's aim should be to work alongside the department's staff during the delivery of the system, actively identifying any gaps in capability and providing training and coaching where needed. The partner's eventual intention should be to disengage from the department, leaving the department able to function on its own.

The partner will need to work on-site with the department's staff to make this possible.

## Continuous delivery

The partner should be committed to the early and frequent delivery of working software, and to the gathering of and responding to feedback from all stakeholders within the department.

Applying a [continuous delivery][] model requires the application of techniques such as continuous integration, automated testing and automated deployment. The partner should be experienced with these.

### Agile approach

The partner should have a track record with an Agile approach to software development, aligned with the [Agile Manifesto][manifesto]. This approach should be applied in an adaptive and contexual rather than adogmatic way, and should apply learnings from such methods as [Lean][], Flow and Kanban.

A cross-functional, non-siloed approach to roles within the team is preferred. For example, the QA function should be engaged as soon as each user story is picked up for development, and should work with the developers and BAs throughout the story's lifetime.

[Automated testing][] should not be restricted to functional requirements. [Performance, load, and soak testing](/service-manual/operations/load-and-performance-testing) will be required to ensure the system behaves as required, and the partner should be experienced with these forms of testing.

## Quality

[Quality][] should be "baked in" to the product. The advantages brought by both internal and external quality should be understood. Test automation is essential for a continuous delivery, and a deep understanding of the pros and cons of automated tests at various levels of the ["test pyramid"](http://martinfowler.com/bliki/TestPyramid.html) is required.

## Alignment

The partner should be aligned to the [Government][governmentstrategy] and [Departmental][] Digital Strategies, and to the [Government Service Design Manual][designmanual], and to have an ability to sell these visions within the department.

### Asking the difficult questions

The partner needs to work with the department to build an atmosphere of healthy challenge, in which all ideas are open to question; those from the department, the partner, and anywhere else. The partner should seek to do the right thing rather than merely what is asked of them, and work to achieve a consensus as to what this might be. This will need to be done with sensitivity, and a constructive approach from all parties.

## Evolutionary Architecture

The partner should have a track record with taking an evidence-based, evolutionary, domain-driven approach to developing an architecture, based on a solid understanding of system's requirements and constraints. Architecture should be developed in a hands-on fashion in concert with the team's developers, with decisions based on discovery through code.

## Specific Expertise

Across the team, we'd like to see specific expertise with a number of practices. Not all these practices are mandatory for an [Agile team](/service-manual/agile/), but we'd expect a prospective partner to be able to be familiar with these practices, to be able discuss them, and to decide upon their use pragmatically.

Note that how these capabilities are split across roles will differ from team to team. What matters is that these things can be done, not who in the team does what.

### Development practices 

* Pair programming
* Test driven development
* Automated testing
* Refactoring
* Appropriate use of design patterns
* Version control usage
    * Frequent checkins
    * Awareness of the pros and cons of branches versus feature toggles
* Appropriate use of design patterns
    
### Exploratory testing

### Business Analysis

* Collaborating with customers to develop a deep understanding of the business domain
* Requirement collection
* Domain modelling and development of a shared vocabulary
* Expressing requirements in the form of well sized stories
    * The characteristics of a good story: [INVEST][]
    
[invest]: http://xp123.com/articles/invest-in-good-stories-and-smart-tasks/ "Independent, Negotiable, Valuable, Estimable, Small, Testable"
        
### Agile project planning

* Inception and other discovery activities
* The building of an initial story list (minimum viable product), and the maintenance of the backlog
* Estimation and sizing - including an understanding of the benefits and costs of estimation
* Iteration management
* Progress tracking
* Retrospectives
    
### Project Management

* Delivery assurance
* Interfacing outside the agile team
    * Finance
    * Procurement
* Delivery assurance
* Product Ownership
    * Might do it, or just ensure it's done
        
### Continuous delivery

* Delivering working software, early and often
* Ability to talk about various tooling options, rather than being familiar with only one
* [Devops][] should be understood as an approach, not as a role
        
### Automated environment provisioning and deployment

* Provisioning tools such as [Chef][], [Puppet][] and [Vagrant][], etc
* Deployment tools such as [Fabric][], etc
* Cloud environments
    * Hosting versus [IaaS][] versus [PaaS][]
    
[puppet]: http://projects.puppetlabs.com/projects/puppet
[chef]: http://www.opscode.com/chef/
[vagrant]: http://www.vagrantup.com/
[fabric]: http://fabfile.org
        
### Enablement

* Identification of capability gaps
* Coaching
* Training
* Mentoring
    
### UX

* [XD](http://en.wikipedia.org/wiki/Experience_design "Experience Design")
* Prototyping
* Design
* Client-side coding
* User testing

## Tool and platform selection

The partner will be involved in the selection of tools and platforms for the project. They will need to be aligned with the [Government Service Design Manual][designmanual] while making these selections.

The partner should not be predisposed to building solutions based upon products for which they are also the vendor. They should be able to make judgements on the applicability to particular situations of alternatives such as:

* Mainstream versus cutting-edge. Examples might include:
    * Java versus Ruby versus Clojure
    * Spring MVC versus Rails
    * Relational database versus NoSQL data stores
* [Software as a service][SaaS] versus [Platform as a service][PaaS] versus [Infrastructure as a Service][IaaS].
* Open source versus proprietary
* Build versus Buy

Decisions should be based on evidence gathered via hands-on, code-based experimentation. A collaborative approach involving workshops and the building of prototypes with vendors is encouraged.

See also [Choosing Technology](/service-manual/making-software/choosing-technology.html).

## Integration experience

Depending on the department's existing estate, the partner may need experience in building systems capable of integrating into complex environments including significant legacy systems.

An agile approach to integration is preferred, utilising such things as:

* consumer-driven contracts
* isolation for testing
* dependency management
* collaboration with vendors

Experience of integrating with legacy systems might include experience with:

* messaging technologies such as SOAP, CORBA or EDIFACT
* mainframe systems
* J2EE, Websphere

## Operational experience

The partner should have proven experience with running day to day operations for a large consumer web service

## Legacy system migration

Depending on the department's existing estate, the partner may need to have have a proven track record with the migration of complex legacy systems. 

## Process

The partner should:

* adopt a pragmatic approach to ITIL
* show an awareness of auditing and standards (ISO27001, PCI, etc)

[Automated testing]: /service-manual/making-software/testing-in-agile
[continuous delivery]: /service-manual/agile/continuous-delivery
[departmental]: http://publications.cabinetoffice.gov.uk/digital/#departmental-strategies "Departmental Digital Strategies"
[designmanual]: https://www.gov.uk/service-manual "Government Service Design Manual"
[Devops]: /service-manual/operations/devops
[governmentstrategy]: http://publications.cabinetoffice.gov.uk/digital/strategy/ "Government Digital Strategy"
[IaaS]: http://en.wikipedia.org/wiki/Infrastructure_as_a_service "Infrastructure as a Service"
[Lean]: http://en.wikipedia.org/wiki/Lean_software_development
[manifesto]: http://agilemanifesto.org/ "The Agile Manifesto"
[PaaS]: http://en.wikipedia.org/wiki/Platform_as_a_service "Platform as a service"
[Quality]: /service-manual/agile/quality
[SaaS]: http://en.wikipedia.org/wiki/Software_as_a_service "Software as a service"
