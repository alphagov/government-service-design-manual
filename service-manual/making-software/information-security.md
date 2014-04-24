---
layout: detailed-guidance
title: Information security
subtitle: Ensuring user data stays secure
category: making-software
type: guide
audience:
  primary: service-managers, web-ops, developers, tech-archs,
  secondary: delivery-managers, qa
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

When building your service, you'll need to ensure that appropriate steps are
taken to ensure its security. Information security is a topic both broad and
deep, drawing from fields ranging from economics and psychology through to
mathematics and probability.

This document cannot claim to provide a
thorough review of the field. Instead, it aims to provide you with a brief
introduction to information security, and will explain the communities and
processes to help you build world-class secure services.

## Introduction to information security

The term information security refers to the theory and practice of defending
data or information systems against unauthorised or unintended access,
destruction, disruption or tampering. Security professionals frequently refer to
3 main concepts:

Confidentiality
: the assurance that information is not disclosed to individuals or systems that
  are not authorised to receive it.

Integrity
: the assurance that information can't be modified by those who are not
  authorised to modify it, or that any such modifications will not pass
  undetected.

Availability
: the assurance that information is available when it's needed, specifically
  that mishap or malice cannot affect the ability of systems to provide the
  information when requested.

In government, much is made of these 3 main "concepts of information
security," as will be explained below. Security systems typically attempt to
address one or more of these concerns through:

- physical controls: walls, locked doors, guards
- procedural controls: managerial oversight, staff training, defined emergency
  response processes
- regulatory controls: legislation, policy, rules of conduct
- technical controls: cryptographic software, authentication and authorisation
  systems, secure protocols

Not every system requires a full battery of security controls. Indeed,
'completely secure systems' don't exist, and overly secure systems are often
prohibitively expensive or thoroughly inconvenient for their users.

You should aim to build services that are *appropriately secure*, and in practice you will
be guided by an assessment of the risks associated with a lapse in
the confidentiality, integrity, or availability of your service.

## Information security in government

Within government, there's an established set of assurance and accreditation
processes. These provide a structure and a shared language within which to
discuss, analyse and address security considerations. If the processes work
correctly, managers should have a clear and accurate understanding of what risks
they're accepting, and those providing the service should know what controls
they're going to employ to mitigate those risks.

> **Assurance** is the broad set of activities involved in assessing and
> managing the risks associated with the system under development, while
> **accreditation** refers to a subset of the assurance work, involving a formal
> and independently verified process similar to
> [ISO27001](http://en.wikipedia.org/wiki/ISO/IEC_27001).

The important thing to note about building trustworthy and secure systems is
that it's a team game. Assurance and accreditation should not be a completely
separate strand of work, or be seen as a hurdle to be jumped over (or
sidestepped). Only by engaging with risk and making decisions based on a range
of expert opinion will you end up with the best product.

The rest of this document will introduce you to assurance and accreditation in
government. The content will use quite a lot of acronyms; unfortunately these
are in common usage and it's very hard to engage with the existing documentation
and processes without speaking the lingo. We include them here in the hope that
they'll provide a helpful reference which can be used when reading existing
documentation.

## Roles

It's important to understand the different roles involved within the process
detailed below. One of the first things you should do on your project is to
establish who plays each of these roles. Note that all of them require formal
training and specialist skills.

Senior Information Risk Owner (SIRO)
: Generally a senior member of the lead organisation providing the service.
  Ultimately responsible for the risk profile of the service, it is the SIRO's
  job to know whether all the risks been identified, whether there appropriate
  mitigations in place so that the risks can be accepted, and so on.

Accreditor
: More hands on than the SIRO, the accreditor or accreditors will work with the
  project team to help with understanding the process, identify risks and
  suggest mitigations.

CESG Listed Adviser Scheme (CLAS) Consultant
: Part of the formal accreditation process. Responsible for much of the formal
  documentation.

CESG (originally Communications-Electronics Security Group)
: [CESG](https://www.cesg.gov.uk/) are the government agency
  responsible for information security. They can provide technical assistance or
  consultation on project issues.

All projects, however small, should involve some level of assurance. This may be
as simple as documenting the limited risks and proposing to the SIRO that the
project does not require a formal accreditation. For anything involving
sensitive data or of interest to lots of people an accreditation stage will be
required, and this process is likely to include representatives from all of the
above groups.

## Business Impact Levels

Business Impact Levels, often shortened to Impact Levels (IL) are a set of
numbers used to guide discussions of risk in government projects. Specifically
they are numbers between 0 and 6 for each of the 3 main concepts mentioned
above, and measure:

* for confidentiality: the potential impact if the information is seen by those who should not see it
* for integrity: the potential impact if the accuracy or completeness of the information is compromised
* for availability: the potential impact if the information becomes inaccessible

More details about identifying these numbers can be found in this [extract from HMG IA Standard No. 1](http://www.cesg.gov.uk/publications/Documents/business_impact_tables.pdf).

## Agile Design

The role-holders listed above will work with the wider team to bring the appropriate concerns to bear in the process of designing the service. The team as a whole will need to make a range of decisions about topics such as what information needs to be captured, how it's processed, whether it's stored, and so on which will have a direct impact on the assurance/accreditation process.

A close working relationship will be essential to make sure that business impact levels and other details are kept up to date as designs evolve and that risk management plays an appropriate role as a constraint in the design process.

## Good Practice Guides (GPG)

The Good Practice Guides (GPG) are documents published by CESG on specific topics of interest to various types of projects. These can act as a good starting point when looking to identify risks and put in place mitigations.

Unfortunately many of these documents are Restricted. It's advisable to establish a working relationship with CESG early on in the project to make sure you can access these documents. Examples include:

* GPG13 -- Protective monitoring
* GPG8 -- Protecting External Connections to the Internet
* GPG12 -- Use of Virtualisation Products for Data Separation

## Risk Management Document Set (RMADS)

The Risk Management Document Set or RMADS are the result of the formal accreditation work. This is likely a large set of documents, including the Baseline Control Set (BCS), system overview and supporting evidence, presented to the SIRO for sign-off as part of go-live conversations.

## IT Health Check (ITHC)

The IT Health Check (ITHC) forms part of the formal accreditation. In essence it's a penetration test carried out by a CESG approved supplier (specifically a CHECK certified individual). Read the guide about [penetration and vulnerability testing](/service-manual/operations/penetration-testing.html) for more details.

## Ongoing

The assurance and accreditation work described above is not just about getting a project to launch. It also covers the running of the resulting service. Over time, new threats may emerge, systems and processes may change, and assumptions may become invalid.

Documentation should be kept up-to-date and additional penetration tests organised on a regular or as-needed basis.

## Tools

It is important to start understanding risks and engaging with the assurance and accreditation process as early in a project as possible. This checklist is a good starting place for milestones to add to a project plan:

1. Identify the SIRO
2. Work with an accreditor to identify the Business Impact Levels
3. Confirm with the SIRO the target Business Impact Levels
4. Confirm with the SIRO whether a formal accreditation is required
5. Procure a CLAS consultant if needed for the accreditation work
6. If possible establish a contact at CESG who can offer assistance and some technical oversight
7. Produce supporting documentation; for example architecture documentation, risks and mitigations, operating processes, references to GPGs, controls
8. Work with the CLAS consultant on completing the RMADS if required
9. Arrange the ITHC
10. Present to the SIRO to get final sign-off

## Risk management

It's important to understand the assurance and accreditation processes and tools are all about managing the risk associated with the running service. Security is part of this, but just one part.

Nearly everything brings with it risks: technology choice, staffing, processes, access to Restricted documents, data aggregation, etc. It's important to understand those risks and put in place sensible and suitable mitigations. It's unrealistic in most cases to aim for a system with no risks, and ignoring them is a recipe for future failure.

The aim is a system where the risks are known and the team, working with risk professionals, have made careful decisions about how to deal with them.

## Further reading

* [Security Engineering — Ross
  Anderson](http://www.cl.cam.ac.uk/~rja14/book.html): a comprehensive textbook covering the theory and practice of building secure systems
* [HMG IA Standard No. 1 -- Technical Risk Assessment](http://www.cesg.gov.uk/publications/Documents/is1_risk_assessment.pdf): the CESG guide to assessing risk in information systems
* [Business Impact Level Tables](http://www.cesg.gov.uk/publications/Documents/business_impact_tables.pdf): Business Impact Level Tables

*[BCS]: Baseline Control Set
*[CESG]: UK National Technical Authority for information assurance
*[CHECK]: A CESG scheme to accredit 3rd party penetration testers
*[CLAS]: CESG Listed Adviser Scheme
*[GPG]: Good Practice Guide
*[GPGs]: Good Practice Guides
*[IA]: information assurance
*[IL]: Impact Level
*[ITHC]: IT Health Check
*[RMADS]: Risk Management Document Set
*[SIRO]: Senior Information Risk Officer
