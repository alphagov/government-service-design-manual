---
layout: category-index
title: Service subdomain names
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

## Subdomain names and certificates

When you are operating a service at `servicename.service.gov.uk`, you will need
testing and development environments. Given that these [must use HTTPS](https),
we suggest purchasing wildcard certificates with this naming convention:

* *.preview.servicename.service.gov.uk
* *.staging.servicename.service.gov.uk
* *.servicename.service.gov.uk

The reasons why we prefer wildcard certificates include:

* being able to take advantage of
  [Server Name Indication (SNI)](http://en.wikipedia.org/wiki/Server_Name_Indication).
  This is essentially HTTP 1.1 virtual-hosting for HTTPS. It may not be
  possible to enumerate all of the names that will be served from a domain,
  so by allowing wildcards, we can meet that need. (You may have the
  public-facing www, admin and assets, but also non-public logging, monitoring
  and other services that still require TLS)
* having a different certificate for preview versus staging versus production
  means that we can potentially restrict who has access to the certificate.
  This is better from an operational security perspective. We should not use
  the same certificate on production as in development/test environments.
