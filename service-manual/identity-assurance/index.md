---
layout: detailed-guidance
title: Identity assurance
subtitle: Secure access to public services
category: identity-assurance
type: category-index
audience:
  primary: service-managers, designers, developers, tech-archs, performance-analysts, user-researchers, content-designers
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
---

Many public services require the user to prove their identity as an early and necessary step to accessing government services.  

The Identity Assurance Service will provide a simple, secure way for people to access public services on GOV.UK.

The service is currently being prototyped. The IDAP team is working closely with several government departments, and a beta for a small number of government services will launch from late 2013.

##The Identity Assurance Service

The Identity Assurance Service, developed by the Identity Assurance Programme (IDAP) within the Government Digital Service (GDS) will be the default service for all government departments providing public digital services which require identity assurance.

You should be working closely with IDAP to ensure that your identity assurance solution is aligned with IDAP’s cross-government platform.

##How does it work?

IDAP is implementing an Identity Assurance Service that allows a user to authenticate themselves via an identity provider (IdP). IdPs are private sector organisations certified against government standards. They will perform checks to confirm the user’s identity and then assert that identity to the government department via the Identity Assurance Service platform.

In practice, this will mean that users of services will be prompted to log in using a third party while completing a transaction with government. If they log in successfully, and the IdP can assert that user’s identity, then the user will be able to complete that transaction.

The Identity Assurance Service is about ‘who the user is’ not 'what the user is entitled to’. The government department is responsible for determining entitlement. 

##Levels of assurance

The government department determines the level of assurance (LoA) they require for the digital service.

**LoA 1** is used when a government department needs to know that it is the same user returning to the service but does not need to know who that user is.

**LoA 2** is used when a government department needs to know who the user is and that that they are a real person to a sufficient level of confidence to be offered in support of civil proceedings

**LoA 3** is used when a government department requires to know who the user is and that that they are a real person to a sufficient level of confidence to be offered in support of criminal proceedings.

##Further reading

* [Good Practice Guide 43:](https://www.gov.uk/government/publications/identity-assurance-enabling-trusted-transactions) Requirements for secure delivery of online public services
* [Good Practice Guide 44:](https://www.gov.uk/government/publications/identity-assurance-enabling-trusted-transactions) Authentication Credentials in Support of HMG Online Services 
* [Good Practice Guide 45:](https://www.gov.uk/government/publications/identity-assurance-enabling-trusted-transactions) Identity Proofing and Verifying the Identity of an Individual in Support of HMG Online Services
* [Identity assurance: enabling trusted transactions](https://www.gov.uk/government/publications/identity-assurance-enabling-trusted-transactions)
* [Blog posts from the IDA team](http://digital.cabinetoffice.gov.uk/?s=identity+assurance)