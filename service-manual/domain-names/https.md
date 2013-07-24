---
layout: category-index
title: HTTPS
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

## HTTPS

Many services will collect personal information from users. It’s very important that this information can’t be intercepted by malicious third parties as it travels over the Internet.

Therefore, all services accessed through service.gov.uk domains (including APIs) MUST only be accessible through secure connections. For web-based services this means HTTPS only (often referred to by the acronyms TLS or SSL, which both refer to the protocol underpinning these secure connections). Services MUST NOT accept HTTP connections under any circumstances.

### Strict Transport Security

Once a service manager has verified that their HTTPS setup is working fine they SHOULD enable [HSTS](http://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security) on the production domains (`www.`, `admin.` and `assets.`), by setting HTTP an HTTP response header such as

    Strict-Transport-Security: max-age=1209600, includeSubDomains;

representing a commitment to HTTPS-only traffic for 14 days. Once the service manager is confident that HSTS is configured correctly, you SHOULD increase the commitment to months or years:

    Strict-Transport-Security: max-age=31536000, includeSubDomains;


### Verification

Ssl vendors vary but regardless of which service you choose to use, you’ll need to validate your purchase with the vendor. The simplest method is usually to request that an email be sent to the registrant email address listed in whois. Alternatively, your vendor may offer you a list of pre-approved email addresses to choose from. The GDS Infrastructure team can validate requests sent to the following addresses:

    hostmaster@digital.cabinet-office.gov.uk
    webops@digital.cabinet-office.gov.uk
    webmaster@digital.cabinet-office.gov.uk
