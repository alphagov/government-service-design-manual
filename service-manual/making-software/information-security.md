---
layout: detailed-guidance
title: Information security
subtitle: Ensuring user data stays secure
category: making-software
type: guide
audience:
  primary: service-manager, developer, tech-arch
status: draft
phases:
  - alpha
  - beta
  - live
---

It goes without saying that security of Government services is incredibly
important. The Government Information security community and processes
exist to help service managers both meet their obligations to those
processess, and more importantly to help build world class services.

The assurance and accreditation processes exist to provide a structure for those
activities with a shared language to allow risks, mitigations (and opportunities)
to be clearly understood by everyone throughout the organisation delivering
a service.

The important thing to note about building trustworthy and secure systems is
that it's a team game. assurance and accreditation should not be a completely
separate strand of work, or seen as a hurdle to be got over (or around). Only
by engaging with risk and making decisions based on a range of expert opinion
will you end up with the best product.

The following uses quite a lot of acronyms. Unfortunately these are in common
usage and it's very hard to engage with the existing documentation and
processes without speaking the lingo. All the acronymns should be explained
before being used.

## Assurance and accreditation

### Roles

It's important to understand the different roles involved within the process
detailed below. One of the first things you should do on your project is to
establish who plays each of these roles. Note that all of them require formal
training and specialist skills.

* Senior Information Risk Owner (SIRO) - generally a senior member of the lead organisation providing the service. Ultimately responsible for the risk profile of the service: have all the risks been identified, are there appropriate mitigations in place so that the risks can be accepted, and so on.
* Accreditor - more hands on than the SIRO, the accreditor or accreditors will work with the project team to help with understanding the process, identify risks and suggest mitigations.
* CLAS Consultant - part of the formal Accreditation process. Responsible for much of the formal documentation.
* [CESG](http://www.cesg.gov.uk/Pages/homepage.aspx) - the government agency responsible for Information Security. Can provide technical assistance or consultation on project issues.

### Assurance vs accreditation

Assurance is the wider set of activities involved in assessing and managing the
risks associated with the system under development. Accreditation is a subset
of the assurance work, involing a formal and externally verified process simiar
to [ISO27001](http://en.wikipedia.org/wiki/ISO/IEC_27001).

All projects, however small, should involve some level of assurance. This may
be as simple as documenting the limited risks and proposing to the SIRO that
the project does not require a formal accreditation. For anything involving
sensitive data or of interest to lots of people an accreditation stage will
be required.

### Business Impact Levels

Business Impact Levels, often shorted to Impact Levels (IL) are a set of
numbers used to guide discussions of risk in Government projects. Specifically
they are numbers between 0 and 6 for each of the three following critera:

* Confidentiality - the potential impact if the information is seen by those who should not see it,
* Integrity - the potential impact if the accuracy or completeness of the information is compromised,
* Availability - the potential impact if the information becomes inaccessible.

More details about identifying these numbers can be found in this [extract from HMG IA Standard No. 1](http://www.cesg.gov.uk/publications/Documents/business_impact_tables.pdf).

### Agile Design

The role-holders listed above will work with the wider team to bring the appropriate concerns
to bear in the process of designing the service. The team as a whole will need to
make a range of decisions about topics such as what information needs to be captured,
how it is processed, whether it is stored, and so on which will have a direct impact
on the assurance/accreditation process.

A close working relationship will be essential to ensure that business impact levels and
other details are kept up to date as designs evolve and that risk management plays an
appropriate role as a constraint in the design process.

### Good Practice Guides (GPG)

The Good Practice Guides (GPG) are documents published by CESG on specific
topics of interest to various types of projects. These can act as a good
starting point when looking to identify risks and put in place mitigations.
For example

Unfortunately many of these documents are Restricted. It is advisable to
establish a working relationship with CESG early on in the project
to make sure you can access these documents. Examples include:

* GPG13 - Protective monitoring
* GPG8 - Protecting External Connections to the Internet
* GPG12 - Use of Virtualisation Products for Data Separation

### Risk Management Document Set (RMADS)

The Risk Management Document Set or RMADS are the result of the formal
accreditation work. This is likely a large set of documents, including the
Baseline Control Set (BCS), system overview and supporting evidence,
presented to the SIRO for sign-off as part of go-live conversations.

### IT Health Check (ITHC)

The IT Health Check (ITHC) forms part of the formal Accreditation. In essense
it is a penetration test carried out by a CESG approved supplier (specifically
a CHECK certified individual). Read the guide about
penetration and vulnerability testing for more details.

### Ongoing

The Assurance and accreditation work described above is not just about getting
a project to launch. It also covers the running of the resulting service. New
threats may emerge or systems and processes change over time. Documentation
should be kept up-to-date and additional penetration tests organised on a
regular or as-needed basis.

## Tools

It is important to start understanding risks and engaging with the Assurance and
Accreditation process as early in a project as possible. The following is a
good starting place for milestones to add to a project plan:

1. Identify the SIRO
2. Work with an accreditor to identify the Business Impact Levels
3. Confirm with the SIRO the target Business Impact Levels
4. Confirm with the SIRO whether a formal Accreditation is required
5. Procure a CLAS consultant if needed for the Accreditation work
6. If possible establish a contact at CESG who can offer assistance and some technical oversight
7. Produce supporting documentation; for example architecture documentation, risks and mitigations, operating processes, references to GPGs, controls
8. Work with the CLAS consultant on completing the RMADS is required
9. Arrange the ITHC
10. Present to the SIRO to get final sign-off

## Why we do this

### Risk management

It's important to understand the Assurance and accreditation processes and tools
are all about managing the risk associated with the running service. Security is
part of this but just one part. Nearly everything brings with it risks:
technology choice, staffing, processes, access to restricted documents, data
agregation, etc. It is important to understand those risks and put in place
sensible and suitable mitigations. It is unrealistic in most cases to aim for a
system with no risks, and ignoring them is a recipe for future exploits. The
aim is a system where the risks are known and the team, working with risk
professionals, have made careful decisions about how to deal with them.

## Further reading

* [HMG IA Standard No. 1 - Technical Risk Assesment](http://www.cesg.gov.uk/publications/Documents/is1_risk_assessment.pdf)
* [G-Cloud Impact Level Guidance](http://gcloud.civilservice.gov.uk/2012/03/09/so-what-is-il3-a-short-guide-to-business-impact-levels/)
* [Business Impact Level Tables](http://www.cesg.gov.uk/publications/Documents/business_impact_tables.pdf)
