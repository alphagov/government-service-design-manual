---
layout: category-index
title: Getting a domain name and start/end page
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

## Getting a domain name and start/end page

To make sure that the right user journey (appropriate start/end pages, clear domain names) are
set up for a new service it's important that you engage with the GOV.UK team within GDS early.

You can start the service domain name process by emailing [gdsapprovals@digital.cabinet-office.gov.uk](mailto:gdsapprovals@digital.cabinet-office.gov.uk)
who will discuss with you the best name for your Service Domain and the start pages on GOV.UK.

While you are waiting for this process, you should start looking at where you will host the
DNS, as GDS will delegate control of your Service Domain to you. You should
[read more about delegating DNS for Service Domains here](/service-manual/domain-names/how-they-work)
which will prepare you for the next step.

Once your service has an agreed name, you will need to supply the following information to your
GDS contact who will put you in touch with the GDS Infrastructure Team to arrange delegation.

1. Service name, eg {service-name}.service.gov.uk
2. DNS servers to delegate to (ask your technical team to see [the guidance on DNS delegation](/service-manual/domain-names/how-they-work))
3. Date you need it by (at least 5 working days' notice)

If you are intending to use the "Managed DNS" product offered by Dyn, then you will need to give as much notice as possible, as Dyn is currently
used to manage the main service.gov.uk DNS domain. Dyn requires a signed letter of authorisation from GDS and work on their systems in order to
manage sub-zones of service.gov.uk via other customer accounts -- this may take additional time to arrange.
