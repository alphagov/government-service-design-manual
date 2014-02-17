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

Our advice is that teams do not build login systems.

Building a login system is a significant undertaking. While there are numerous
open source libraries that make it trivial to add login functionality to your
service, there are significant downsides:

* significantly increasing your user support overhead (people forget how to
  sign in, lose their passwords, etc)
* gathering personal data that you will need to constantly review and protect
* adding a relatively complex interaction for users to complete

## Find alternatives

Many features that are often implemented using login systems can be completed
in other (and potentially more useful) ways.

Saving search results, for example, doesn't require a login but just a way of
helping users remember a specific URL. Instead of having them log in, you could
provide a tool to help send the URL to an email address or instructions on
creating a bookmark in their browser. Or perhaps you could just take their
email address and let them know if the search results change?

The precise details will vary according to what users need from your service,
but if there's an alternative to a login system, that should be preferred.

## Where there aren't alternatives

If after careful review and design work there is no option but to build a
login system you will need to consider a few questions:

* are you providing a login service for a small number of agents (administrative users, accredited partners, etc) or for a broad range of citizens or businesses?
* do you already have all the data you need in order to establish trust with those users or will you need to match them against other services (online or offline)?

If building a service for a small number of clearly identified agents then
it's probably safe to proceed. You should ensure that any authentication
and authorisation code written for your system is carefully separated from
the application in such a way that you can:

* specifically monitor use of the system for attempts to gain access, with identifiers such as unusually high number of failed login attempts over a short period of time, or a sequence of failed logins on a given account over a long period of time
* segment user data from other data you hold to avoid aggregating a large amount of identifiable information
* swap to a new identity system such as the [identity assurance](/service-manual/identity-assurance) scheme without invasive changes to the rest of your codebase

If you need to build a system for a broad range of citizens and businesses,
or you need to do sophisticated matching with other systems in order to build
trust in the identity of your users then you should explore the
[advice published by the ID Assurance team](/service-manual/identity-assurance).

## Credentials

You should help your users to pick strong, secure passwords or phrases and
consider whether it's appropriate to require 2-factor authentication for
extra security.

As a minimum passphrases should be eight characters long and include a mix
of letters, numbers and symbols, but ideally they'll be longer than that. We
refer to passphrases as a phrase is usually easier to remember but harder to
guess than a short collection of symbols or a single word.

> For some admin systems on GOV.UK we use the
  [zxcvbn](https://github.com/lowe/zxcvbn)
  library that measures how hard it would be for a computer to crack a
  passphrase using brute-force methods. That library is used to validate
  new passphrases and insist on strong passphrases. It's only one measure,
  but it increases our confidence that our users are picking good
  passphrases.

All new government services should be served over
[HTTPS](/service-manual/domain-names/https.html) to ensure the
communication between the user and the service is encrypted. This is
especially important when logging in.

## Examples

The Electronic License Management System (ELMS) license application system
on Business Link required a login to complete an application. In building
a new version of the system for
[GOV.UK](https://www.gov.uk/browse/business/licences) we removed that
requirement and usage rates have increased considerably. There's still a
login system for approved users in local authorities who need to process
those applications.

## Also see

* [Wikipedia on Password policies](http://en.wikipedia.org/wiki/Password_policy#Password_length_and_formation)
* [CESG Good Practice Guide 44: authentication credentials in support of HMG online services](https://www.gov.uk/government/publications/identity-assurance-enabling-trusted-transactions)
* [xkcd cartoon explaining password vs. passphrase](https://xkcd.com/936/)
