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

When building your service, you'll need to make sure that appropriate steps are
taken to guarantee its security.

Information security is a topic both broad and deep, drawing from fields ranging from economics and psychology through to mathematics and probability.

This document doesn’t provide a thorough review of the field. Instead, it aims to:

* provide you with the briefest of introductions to information security
* explain the communities and processes that exist to help you build world-class secure services

## What's information security?

The term information security refers to the theory and practice of defending
data or information systems against:

* unauthorised or unintended access
* destruction
* disruption
* tampering

### Concepts

There are 3 main concepts that security professionals frequently refer to:

- confidentiality - the assurance that information is not disclosed to individuals or systems that
  are not authorised to receive it
- integrity - the assurance that information can't be modified by those who are not
  authorised to modify it, or that any such modifications will not pass
  undetected
- availability - the assurance that information is available when it's needed, and that mishap or
  malice cannot affect the ability of systems to provide information when requested

In government, much is made of these 3 main "concepts of information
security," as will be explained below. Security systems typically attempt to
address one or more of these concerns through:

- physical controls, eg walls, locked doors, guards
- procedural controls, eg managerial oversight, staff training, defined emergency
  response processes
- regulatory controls, eg legislation, policy, rules of conduct
- technical controls, eg cryptographic software, authentication and authorization
  systems, secure protocols

Not every system requires a full range of security controls.
'Completely secure systems' don't exist, and overly secure systems are often
too expensive or thoroughly inconvenient for their users.

Aim to build a service that is appropriately secure, which you can achieve by managing risks. In practice, you’ll
be guided by an assessment of the risks related to a lapse in
the confidentiality, integrity, or availability of your service.

## Information security in government

Government has a fixed set of assurance and accreditation
processes. These provide a structure and a shared vocabulary for
discussing, analysing and addressing security considerations.

If the processes work correctly:

* you'll have a clear and accurate understanding of what risks you’re accepting
* those delivering the service will know what controls they’re going to use to reduce risks

The difference between assurance and accreditation processes is that:

* assurance is the broad set of activities involved in assessing and
  managing the risks associated with a system under development
* accreditation is about a subset of the assurance work, involving a formal
  and independently verified process similar to
  [ISO27001](https://en.wikipedia.org/wiki/ISO/IEC_27001)

The important thing to note about building trustworthy and secure systems is
that it's a team game.

Involve some level of assurance in your project, no matter how small. This may be as simple as documenting the limited risks and proposing to the Senior Information Risk Owner (SIRO) that the project doesn’t require a formal accreditation.

An accreditation stage will be required for anything involving sensitive data or of interest to lots of people. This process is likely to include representatives from all of the above groups.

Don’t make assurance and accreditation a completely separate strand of work, or see it as a hurdle to be jumped over (or sidestepped). You’ll only end up with the best product by dealing with risk and making decisions based on a range of expert opinion.

## Government assurance and accreditation

The following content will introduce you to assurance and accreditation in government and will use quite a lot of acronyms. Unfortunately these are in common usage and it’s very hard to work with the existing documentation and processes without speaking the lingo.

It’s included here in the hope that they’ll provide a helpful reference that can be used when reading existing documentation.

### Roles

It’s important to understand the different roles involved with the assurance and accreditation processes. Each role requires formal training and specialist skills, so decide who plays each of these roles from the very beginning:

* [Senior Information Risk Owner](#Senior-Information-Risk-Owner) (SIRO)
* [Accreditor](#Accreditor)
* [Communications Electronics Security Group](#CESG) (CESG)
* [CESG Listed Adviser Scheme (CLAS) consultant](#CESG-Listed-Adviser-Scheme)

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

#### CESG (originally Communications Electronics Security Group)

[CESG](http://www.cesg.gov.uk/Pages/homepage.aspx) are the government agency responsible for Information Security. They can provide technical assistance or consultation on project issues.

#### CESG Listed Adviser Scheme (CLAS) consultant

CLAS consultants are part of the formal accreditation process and responsible for a lot of the formal documentation.

### Business impact levels

Business impact levels, often shortened to impact levels (IL) are a set of
numbers used to guide discussions about risk in government projects. Specifically
they are numbers between 0 and 6 for each of the 3 main concepts mentioned
above, and measure:

* confidentiality - the potential consequences of information being seen by those who should not see it
* integrity - the potential consequences of having the accuracy or completeness of information compromised
* availability - the potential consequences of information becoming inaccessible

More details about identifying these numbers can be found in this [extract from HMG IA Standard No. 1](https://www.cesg.gov.uk/publications/Documents/business_impact_tables.pdf).

### Agile design

The role holders listed above will work with the wider team to raise the appropriate concerns in the process of designing the service. Your team as a whole will need to make a range of decisions about topics like:

* what information needs to be captured
* how information is processed
* whether information is stored

It’s essential to have a close working relationship so you can make sure that:

* business impact levels and other details are kept up to date as designs evolve
* risk management plays an appropriate role as a constraint in the design process

### Good Practice Guides

The Good Practice Guides (GPG) are documents published by CESG on specific topics of interest to various types of projects. These can act as a good starting point when looking to identify risks and put in place mitigations.

Unfortunately many of these documents are Restricted. Set up a working relationship with CESG early on in your project to make sure you can access these documents. Examples include:

* GPG13 -- Protective Monitoring
* GPG8 -- Protecting External Connections to the Internet
* GPG12 -- Use of Virtualisation Products for Data Separation

### Risk Management Document Set

The Risk Management Document Set or RMADS is the result of your formal accreditation work. It’s likely to be a large set of documents presented to the SIRO for sign-off as part of go-live conversations, including:

* the Baseline Control Set (BCS)
* the system overview
* supporting evidence

### IT Health Check

The IT Health Check (ITHC) forms part of the formal Accreditation. In essence it's a penetration test carried out by a CESG approved supplier (specifically a CHECK certified individual). Read guidance on [penetration and vulnerability testing](/service-manual/operations/penetration-testing) for more details.

### Ongoing

Assurance and accreditation work is not just about getting a project to launch, but also covers the running of the resulting service. Over time, new threats may emerge, systems and processes may change, and assumptions may become invalid.

Keep your documentation up to date and carry out additional penetration tests on a regular or as-needed basis.

### Tools

Start understanding risks and get involved with the assurance and accreditation process as early on in your project as possible. This checklist is a good starting place for adding milestones to your project plan:

* identify the SIRO
* work with an accreditor to identify the business impact levels
* confirm the target business impact levels with the SIRO
* confirm with the SIRO whether a formal Accreditation is required
* procure a CLAS consultant if needed for the Accreditation work
* set up a contact at CESG who can offer assistance and some technical guidance, if possible
* produce supporting documentation, eg
  * architecture documentation
  * risks and mitigations
  * operating processes
  * references to GPGs
  * controls
* work with the CLAS consultant on completing the RMADS required
* arrange the ITHC
* present to the SIRO to get final sign-off

### Risk management

It’s important to understand that the assurance and accreditation processes and tools are all about managing the risk associated with the running service. Security is part of this, but just one part.

Nearly everything will contain risks:

* technology choice
* staffing
* processes
* access to Restricted documents
* data aggregation

Understand those risks and put in place sensible and suitable mitigations. It's unrealistic in most cases to aim for a system with no risks, and ignoring them is a recipe for future failure.

Aim for a system where the risks are known, and have your team work with risk professionals to make careful decisions about how to deal with them.

## Further reading

* [Security Engineering — Ross
  Anderson](http://www.cl.cam.ac.uk/~rja14/book.html) is a comprehensive textbook covering the theory and practice of building secure systems
* [HMG IA Standard No. 1 -- Technical Risk Assessment](https://www.cesg.gov.uk/publications/Documents/is1_risk_assessment.pdf), the CESG guide to assessing risk in information systems
* [Business Impact Level Tables](https://www.cesg.gov.uk/publications/Documents/business_impact_tables.pdf)

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
