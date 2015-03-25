---
layout: category-index
title: How service.gov.uk domain names are managed
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

## How service.gov.uk domain names are managed

Once [a domain name for a service has been agreed](/service-manual/domain-names/setting-up.html), the
GDS Infrastructure Team allocates service.gov.uk domain names by delegating control to the team running
the service. The team should provide the details of nameservers and GDS will delegate
{service-name}.service.gov.uk to those nameservers. At that point the service team will be able to
manage the detailed configuration of which servers the domain name points to, and so on.

## How is this different to how direct.gov.uk domains have been managed?

The domain direct.gov.uk was managed by the DirectGov Service Desk directly (and inherited by GDS).
All domains were administered by DirectGov and each subdomain did not have its own DNS servers.

This meant that services would need to give details of the names and IP addresses of their webservers
to DirectGov, who would directly create the necessary A and CNAME records. The DirectGov Service Desk
were also a critical part of any DNS change.

The new model for service.gov.uk domains requires that you provide your own DNS servers which can
respond to requests for {something}.{servicename}.service.gov.uk. It is not necessary that you run
and manage your own separate DNS servers:

1. There are service providers on GCloud who will provide DNS Hosting with a web-interface for you
   to manage your DNS records.
2. Many infrastructure hosting providers also provide DNS services for the use of their customers.
3. As a last resort (or for larger services) you may wish to ask the supplier to provide DNS services
   by any means they consider reasonable (including utilising options 1 and 2 above).

## Why do we do this?

It is likely that as you build and iterate your service you will need to make a number of changes to its
configuration. You might introduce load balancers or content delivery networks, you may move hosting
providers, or your provider might need to change the IP addresses that your servers are assigned. In any
of those cases it is important that your team is able to quickly respond to the situation and make any
relevant changes with as few intermediaries as possible. By delegating control GDS ensures that control
is in the hands of the service team and not blocked by a central authority.
