---
layout: detailed-guidance
title: Identity assurance guidance
subtitle: For government service providers
category: identity-assurance
type: guide
audience:
  primary: service-managers, designers, delivery-managers, web-ops, developers, tech-archs, performance-analysts, user-researchers, qa, content-designers
  secondary:
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

## Introduction

Users of digital government services need to be able to sign in securely and conveniently to access services and records, and be confident that their data is secure and their privacy protected.

Identity assurance is the way that users of digital services can sign in securely to digital services, and digital services can be confident that users are who they say they are. A good background description can be found [on this blog post](https://gds.blog.gov.uk/2014/01/23/what-is-identity-assurance/).

Service providers can be assured they’re providing their service to the right individual or business after matching an assured identity to their own records.

This guidance is for government digital service providers that need to use the identity assurance service. Service providers are referred to as ‘relying parties’ to avoid confusion between those providing the government service to the user and those providing the identity service to the user.

The guidance gives an overview of the identity assurance service and explains what relying parties will need to do to prepare to start using the identity assurance service.

Relying parties that need to use the identity assurance service will need to:

- assess the level and type of identity assurance they require
- build a service to request and consume Security Assertion Markup Language (SAML) identity assertions
- build a matching service to match assured identities to their own records

## How the identity assurance service works

### 1. User is directed to the hub to sign in

When users try to use a service provided by a relying party, they will be directed to the identity assurance hub at the point in the user journey where the relying party needs to know who they are.

The hub manages communications between users, relying parties and identity providers. It makes sure that users can assert their identities securely and safely, and that relying parties can be confident that users are who they say they are.

### 2. Register with an identity provider, if not done previously

Users will be guided through selecting an identity provider from a list of organisations contracted by the Cabinet Office to provide the service.

Identity providers are private sector organisations that are certified as meeting relevant security and service standards. They carry out checks to confirm the user’s identity, using methods designed to meet published standards, and then communicate that identity to the relevant government department through the [hub](https://identityassurance.blog.gov.uk/2013/10/30/a-hub-is-born/).

### 3. Relying party matches the assured identity to its own records about the user

The user returns to the hub once this is complete so they can be matched up to the right record in the relying party using a limited set of information that has been verified by the identity provider:

* name
* address
* date of birth
* gender

The relying parties use this limited set of identity information (the matching dataset) to make sure they're dealing with the right person. This is known as matching. This involves looking for the correct record in the relying party’s database which matches the assured identity provided by the user.

The hub sends the matching dataset to the relying party’s matching service. Relying parties are responsible for:

* the process that decides which relying party record to match against
* managing the risk of incorrect matching
* handling the case of multiple (many records found) or no matches (user not found)

Once matching has taken place, the user is directed to the correct record within the relying party’s service. The identity assurance team can advise services on how to develop their matching service.

### 4. User is given access to their account within the service they want to use

Once a user has registered with an identity provider, they will be able to sign in to all digital services that use the identity assurance service and require the same level of assurance, using credentials provided to them by their identity provider.

The first government services to start using the identity assurance service will connect to the hub in early 2014. Each other service will connect at a time agreed with the programme team. Relying parties that have identity assurance requirements should contact the identity assurance programme team to start planning their connection to the service.

Read the [blog posts from the identity assurance programme team](https://identityassurance.blog.gov.uk/) for more information and updates about the development of the identity assurance service,

## How to assess the identity assurance requirements of a digital service

Relying parties can assess their identity assurance needs by using [guidance published by the Cabinet Office](https://www.gov.uk/government/collections/identity-assurance-enabling-trusted-transactions). Each relying party must carry out a risk assessment of their digital service, using [Good Practice Guide 43: Requirements for secure delivery of online public services][gpg43].

The identity’s ‘level of assurance’ is the degree of confidence the relying party requires that a user is who they say they are:

* level of assurance 1 is used when a relying party needs to know that it is the same user returning to the service but does not need to know who that user is
* level of assurance 2 is used when a relying party needs to know on the balance of probabilities who the user is and that that they are a real person
* level of assurance 3 is used when a relying party needs to know beyond reasonable doubt who the user is and that that they are a real person
* level of assurance 4 is as level of assurance 3, but with a biometric profile captured at the point of registration. This level is not within the scope of this stage in the identity assurance programme

The different levels of assurance available and how they are reached are defined by the Cabinet Office's good practice guides. Each relying party must determine the level of assurance required for each digital service:

* [Good Practice Guide 45: Identity Proofing and Verification of an Individual][gpg45] explains the requirements for achieving the levels of assurance required to register an individual identity
* [Good Practice Guide 44: Authentication Credentials in Support of HMG Online Services][gpg44] explains the level of authentication strength required to authenticate an individual when they return to sign in using their assured identity
* [Good Practice Guide 46: Organisation Identity][gpg46] explains the level of identity needed by organisations for proofing and verification

[gpg43]: https://www.gov.uk/government/publications/requirements-for-secure-delivery-of-online-public-services
[gpg44]: https://www.gov.uk/government/publications/authentication-credentials-for-online-government-services
[gpg45]: https://www.gov.uk/government/publications/identity-proofing-and-verification-of-an-individual
[gpg46]: https://www.gov.uk/government/publications/identity-assurance-organisation-identity

## How to work with the identity assurance programme to prepare to use identity assurance

Each relying party will need to go through the programme’s induction process. This process consists of 4 structured stages:

1. Introduction - identify a need for identity assurance.
2. Needs Analysis - identify the service’s identity assurance requirements (levels and types), users and dependencies. Agree a provisional date for connection to hub.
3. Planning - complete a readiness assessment. Agree a detailed plan setting out what the service and GDS will each do and when. Get internal approvals within GDS and the relying party's organisation.
4. Connect to hub - agree a memorandum of understanding between the Identity Assurance Programme and the department responsible for the service.

Government providers of digital services that want to find out more about identity assurance, with a view to becoming relying parties, should contact [Janet Hughes](mailto:janet.hughes@digital.cabinet-office.gov.uk).

## How to match assured identities to service records

Each relying party will need to set up a matching service that can accept digital assertions of identity and match them to their internally held records. This process makes sure they can use the identity assertions provided by the identity assurance service.

Matching assured identities to the records held by relying parties is likely to be complicated - each relying party will need to establish the most efficient and effective way of doing this depending on the quality of its data and how it’s stored and managed.

Relying parties will need to be able to do a risk-based match to ascertain which locally held account to match against. For example, relying parties need to be able to handle users who have proved their identities, but the service can't match with sufficient confidence to any particular record or even multiple records.

GDS has developed a matching service adapter that will help some relying parties to convert their communication with the hub to the appropriate language.

The service provider will be supplied with details of the matching service adapter during the early stages of the induction process.
