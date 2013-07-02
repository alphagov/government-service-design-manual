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

GDS advise that teams do not build login systems.

Building a login system is no easy task. While there are numerous open source libraries that make it trivial to add login functionality to your service, the moment you add that feature you’re: 

* significantly increasing your user support overhead (people forget how to sign in, lose their passwords, etc) 
* gathering personal data that you’ll need to constantly review and protect
* adding a relatively complex interaction for users to complete

## Find alternatives

Many features that often come with using a login system can be completed in other (and potentially more useful) ways.
Saving search results, for example, doesn’t require a login but just a way of helping users remember a specific URL. Instead of having users log in, you could: 

* provide a tool to help send the URL to an email address
* have instructions on creating a bookmark in their browser
* take their email address and let them know if the search results change

The exact details will vary according to what users need from your service, but if there’s an alternative to a login system — use it.  

## When there isn't an alternative

If, after careful review and design work, there’s no option but to build a login system, you’ll need to consider:

* who you’re providing a login service for — a small number of agents (administrative users, accredited partners, etc) or for a broad range of citizens or businesses?
* if you have enough user data — enough to build trust with users or will you need to match them against other services (online or offline)?


It’s probably safe to carry on if you’re building a service for a small number of clearly identified agents. Make sure that any authentication and authorisation code written for your system is carefully separated from the application so that you can:

* specifically monitor use of the system for attempts to gain access, like    
  * unusually high numbers of failed login attempts over a short period of time
  * a sequence of failed logins on a given account over a long period of time
*	separate user data from other data you hold to avoid collecting a large amount of identifiable information
* swap to a new identity system, like the [identity assurance](/service-manual/identity-assurance) scheme, without invasive changes to the rest of your codebase

Read the [advice published by the ID Assurance team](/service-manual/identity-assurance) if you need to:

* build a system for a broad range of citizens and businesses
* do sophisticated matching with other systems so you can build confidence in the identity of your users


## Lower usage rates

The electronic license management system on Business Link required a login to complete an application. This requirement was removed in the building of a new version of the system for [GOV.UK](https://www.gov.uk/browse/business/licences),leading to a considerable increase in usage rates. There is, though, still a login system for approved users in local authorities that need to process those applications.
