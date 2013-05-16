---
layout: detailed-guidance
title: User accounts and logins
subtitle: How to do them and how to avoid them
category: making-software
type: guide
audience:
  primary: developers, tech-archs
  secondary: designers, service-managers
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
  -
    title: Making software
    url: /service-manual/making-software
---

Our advice is that teams **do not** build login systems.

Building a login system is a significant undertaking. While there are numerous open source libraries that make it trivial to add login functionality to your service, the moment you add that feature you're significantly increasing your user support overhead (people forget how to sign in, lose their passwords, etc), you're accruing personal data that you will need to constantly review and protect, and you're adding a relatively complex interaction for users to complete.

## Find alternatives

Many features that are often implemented using login systems can be completed in other (and potentially more useful) ways.

Saving search results, for example, doesn't require a login but just a way of helping users remember a specific URL. Instead of having them log in you could provide a tool to help send the URL to an email address or instructions on creating a bookmark in their browser. Or perhaps you could just take their email address and let them know if the search results change?

The precise details will vary according to what users need from your service, but if there's an alternative to a login system that should be preferred.

## Where there aren't alternatives

If after careful review and design work there is no option but to build a login system you will need to consider a few questions:

* are you providing a login service for a small number of agents (administrative users, accredited partners, etc) or for a broad range of citizens or businesses?
* do you already have all the data you need in order to establish trust with those users or will you need to match them against other services (online or offline)?

If building a service for a small number of clearly identified agents then it is probably safe to proceed. You should ensure that any authentication and authorisation code written for your system is carefully separated from the application in such a way that you can:

* specifically monitor use of the system for attempts to gain access, with identifiers such as unusually high number of failed login attempts over a short period of time, or a sequence of failed logins on a given account over a long period of time
* segment user data from other data you hold to avoid aggregating a large amount of identifiable information
* swap to a new identity system such as the [identity assurance](/service-manual/identity-assurance.html) scheme without invasive changes to the rest of your codebase

If you need to build a system for a broad range of citizens and businesses, or you need to do sophisticated matching with other systems in order to build trust in the identity of your users then you should explore the [advice published by the ID Assurance team](/service-manual/identity-assurance.html).

## Examples

The ELMS license application system on Business Link required a login to complete an application. In building a new version of the system for [GOV.UK](https://www.gov.uk/browse/business/licences) we removed that requirement and usage rates have increased considerably. There is still a login system for approved users in local authorities who need to process those applications.
