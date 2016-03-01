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

* ensure there is a [mail exchanger (MX) record](https://en.wikipedia.org/wiki/MX_record) set up for the domain from which you send email
* implement the [Domain-based Message Authentication, Reporting and Conformance (DMARC)](https://www.gov.uk/government/publications/email-security-standards/domain-based-message-authentication-reporting-and-conformance-dmarc) standard to improve the likelihood that your email will be delivered correctly

Before releasing your service you should test your email delivery. As a minimum you should use your service with
registered email addresses from a range of popular email providers and ensure that emails arrive as you expect.

For more on good practice on email in government, see the [Common Technology Services email blueprint](https://www.gov.uk/guidance/common-technology-services-cts-secure-email-blueprint)
