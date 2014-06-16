---
layout: category-index
title: Asset hosts
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

## Asset hosts

To provide the fastest possible experience for our users it's usually a good idea to serve assets (stylesheets, JavaScript, images) from a separate domain name. That will enable the web browser to fetch more elements of the page in parallel, and also make it straightforward to remove cookies and other extraneous information from the HTTP responses for assets.

You should not use the asset host that we use for www.gov.uk to load your assets. The files provided there are only guaranteed to work for www.gov.uk and could change without notice in ways that would break other services that are using them. Instead you should host any assets you rely on within your service.

We recommend a hostname such as:

    assets.{service-name}.service.gov.uk