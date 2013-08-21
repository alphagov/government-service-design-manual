---
layout: category-index
title: Handling email
category: domain-names
type: category-index
audience:
  primary: service-managers, web-ops, tech-archs
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
  -
    title: Where services live on the web
    url: /service-manual/domain-names
---

## Handling email

Emails to users of your service should be sent from a human-monitored email address that originates from the
domain servicename.service.gov.uk (and not the dept/agency or any other domain name). Users are interacting
with the service and that is where they will expect communications to come from.

### Getting emails to users

In order to protect users from spam email providers put in place a variety of checks. It is often a good idea
to use a trusted specialist third-party to dispatch email as they will have tools and expertise to help ensure
that you pass those checks. As a minimum you should:

* Ensure there is an MX record set up for the domain from which you send email
* Enable SPF on the sending domain.
* Consider using DKIM on the sending domain; it can provide additional guarantees about message delivery and help recipients to more easily distinguish genuine mail from forgery.

Before releasing your service you should test your email delivery. As a minimum you should use your service with
registered email addresses from a range of popular email providers and ensure that emails arrive as you expect.
