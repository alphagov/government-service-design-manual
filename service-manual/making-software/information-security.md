---
layout: detailed-guidance
title: Information security
subtitle: Making sure user data stays secure
category: making-software
type: guide
audience:
  primary: service-managers, web-ops, tech-archs,
  secondary: developers, delivery-managers, qa
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

When building your service, you’ll need to make sure that appropriate steps are
taken to guarantee its security.

Information security is a topic both broad and deep, drawing from fields ranging from economics and psychology through to mathematics and probability.

This document doesn’t provide a thorough review of the field. Instead, it aims to:

* provide you with the briefest of introductions to information security
* explain the communities and processes that exist to help you build world-class secure services

## What’s information security?

The term information security refers to the theory and practice of defending data or information systems against:

* unauthorised or unintended access
* destruction
* disruption
* tampering

### Concepts

There are 3 main concepts that security professionals frequently refer to:

- confidentiality - the assurance that information is not disclosed to individuals or systems that are not authorised to receive it
- integrity - the assurance that information can’t be modified by those who are not authorised to modify it, or that any such modifications will not pass undetected
- availability - the assurance that information is available when it’s needed, and that mishap or malice cannot affect the ability of systems to provide information when requested

In government, much is made of these 3 main "concepts of information security," as will be explained below. Security systems typically attempt to address one or more of these concerns through:

- physical controls, eg walls, locked doors, guards
- procedural controls, eg managerial oversight, staff training, defined emergency response processes
- regulatory controls, eg legislation, policy, rules of conduct
- technical controls, eg cryptographic software, authentication and authorization systems, secure protocols

Not every system requires a full range of security controls. ’Completely secure systems’ don’t exist, and overly secure systems are often too expensive or thoroughly inconvenient for their users.

Aim to build a service that is appropriately secure, which you can achieve by managing risks. In practice, you’ll be guided by an assessment of the risks related to a lapse in the confidentiality, integrity, or availability of your service.

## Information security in government

The [Government Security Classifications Scheme](https://www.gov.uk/government/uploads/system/uploads/attachment_data/file/251480/Government-Security-Classifications-April-2014.pdf) divides information assets into three tiers:

* OFFICIAL
* SECRET
* TOP SECRET

The vast majority of HMG information assets will be classified as OFFICIAL, including personal and sensitive personal citizen data as defined by the [Data Protection Act](http://www.legislation.gov.uk/ukpga/1998/29/contents).

Security at OFFICIAL is achieved through following good commercial practices, using well configured commodity technologies and by people taking personal responsibility and using their judgement more actively.

The [Securing Technology at OFFICIAL](https://www.gov.uk/government/collections/securing-technology-at-official) collection is a valuable resource for understanding how OFFICIAL information should be protected.

## Assurance

Assurance is the broad set of activities involved in assessing and managing the risks associated with a system under development. If your assurance processes work correctly:

* you’ll have a clear and accurate understanding of what risks you’re accepting
* those delivering the service will know what controls they’re going to use to reduce risks

The important thing to note about building trustworthy and secure systems is that it’s a team game.

Involve some level of assurance in your project, no matter how small. This may be as simple as documenting the limited risks.

Don’t make assurance a completely separate strand of work, or see it as a hurdle to be jumped over (or sidestepped). You’ll only end up with the best product by dealing with risk and making decisions based on a range of expert opinion.

It’s also important not to see assurance as a one-off activity, or as a milestone to be achieved and then forgotten about. As with every other element of your service, you should constantly review and iterate your information security practices.

## Service security

When designing or operating a service, you’re responsible for making sure that the service, as a whole, provides appropriate security for the information you receive, process and store.

Whether you’re buying a commodity service such as hosting as a component of your service or designing a service from scratch, the [Cloud Security Principles](https://www.gov.uk/government/publications/cloud-service-security-principles/cloud-service-security-principles) are a useful way of thinking about different elements of your service’s security.

These principles are designed as a common way for suppliers to explain how their services manage information security. They also provide an excellent framework for thinking about how your service should manage information security. Not all of these principles will be applicable to every service but they’re a good starting point.

The [Cloud Security Guidance](https://www.gov.uk/government/collections/cloud-security-guidance) collection contains more information on how these principles work and how security can be implemented against each principle.

## Government assurance

The following content will introduce you to assurance in government and will use quite a lot of acronyms. Unfortunately these are in common usage and it’s very hard to work with the existing documentation and processes without speaking the lingo.

It’s included here in the hope that they’ll provide a helpful reference that can be used when reading existing documentation.

### Roles

It’s important to understand the different roles involved with the assurance process. Each role requires formal training and specialist skills, so decide who plays each of these roles from the very beginning:

* [Senior Information Risk Owner](#senior-information-risk-owner-siro) (SIRO)
* [Accreditor](#accreditor)
* [Information Asset Owner](#information-asset-owner)
* [Communications Electronics Security Group](#cesg-originally-communications-electronics-security-group) (CESG)

#### Senior Information Risk Owner (SIRO)

A SIRO will generally be a senior member of the lead organisation that’s providing the service and is responsible for:

* the risk profile of the service
* identifying all of the risks
* making sure that appropriate mitigations are in place so that the risks can be accepted

#### Accreditor

More hands on than the SIRO, the accreditor or accreditors will work with the project team to help with:

* understanding the process
* identify risks
* suggest mitigations

#### Information Asset Owner

The IAO will be a senior individual involved in running the relevant business.

Their role is to understand what information is held, what is added and what is removed, how information is moved, and who has access and why. As a result they are able to understand and address risks to the information.

#### CESG (originally Communications Electronics Security Group)

[CESG](https://www.cesg.gov.uk) are the government agency responsible for Information Security. They can provide technical assistance or consultation on project issues.

### IT Health Check

The IT Health Check (ITHC) forms part of the assurance process. In essence it’s a penetration test carried out by a CESG approved supplier (specifically a CHECK certified individual). Read guidance on [penetration and vulnerability testing](https://www.gov.uk/service-manual/operations/penetration-testing) for more details.

### Ongoing
Assurance work is not just about getting a project to launch, but also covers the running of the resulting service. Over time, new threats may emerge, systems and processes may change, and assumptions may become invalid.

Keep your documentation up to date and carry out additional penetration tests on a regular or as-needed basis.

## Risk management

It’s important to understand that the assurance processes and tools are all about managing the risk associated with the running service. Security is part of this, but just one part.

Nearly everything will contain risks:

* technology choice
* staffing
* processes
* data aggregation

Understand those risks and put in place sensible and suitable mitigations. It’s unrealistic in most cases to aim for a system with no risks, and ignoring them is a recipe for future failure.

Aim for a system where the risks are known, and have your team work with risk professionals to make careful decisions about how to deal with them.
