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

Many services will collect personal information from users. It’s very important that this information can’t be
intercepted by malicious third parties as it travels over the Internet.

Therefore, all services accessed through service.gov.uk domains (including APIs) MUST only be accessible through
secure connections. For web-based services this means HTTPS only (often referred to by the acronyms TLS or SSL,
which both refer to the protocol underpinning these secure connections). Services MUST NOT accept HTTP connections
under any circumstances.

### Strict Transport Security

Strict Transport Security or [HSTS](http://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security) is an extension
to the HTTPS protocol that tells web browsers that they should make extra efforts to verify the security of a
connection: they should assume for a specified period that all connections with this server should be via HTTPS
and shouldn't accept mixed content (where some content on a page is served via HTTPS and some via HTTP). This
provides an extra level of assurance about the integrity of the connection.

Once a service manager has verified that their HTTPS setup is working fine they SHOULD enable
[HSTS](http://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security) on the production domains (`www.`, `admin.`
and `assets.`), by setting HTTP an HTTP response header such as

    Strict-Transport-Security: max-age=1209600, includeSubDomains;

representing a commitment to HTTPS-only traffic for 14 days. Once the service manager is confident that HSTS
is configured correctly, you SHOULD increase the commitment to months or years:

    Strict-Transport-Security: max-age=31536000, includeSubDomains;

### Verification of SSL Certificates

In order to provide your service over HTTPS you will need to purchase a certificate (or certificates) from a
recognised vendor. SSL vendors vary but regardless of which service you choose to use, you’ll need to validate
with the vendor that you own the domain.

The verification methods that are commonly in use (in order of preference) are:

1. Create a file with a special code supplied by the SSL vendor on your website
2. Create a special DNS record with a code supplied by the SSL vendor
3. Sending an email to the owner of the service.gov.uk domain

The least preferred method of sending an email to the domain owner relies on the GDS Infrastructure Team
seeing the verification email and responding. If this is necessary, then first send an email to the address that
you intend to use for verification from your own email address warning that an SSL verification is needed for
your service.

The GDS Infrastructure Team can validate requests sent to the following addresses:

    hostmaster@digital.cabinet-office.gov.uk
    webops@digital.cabinet-office.gov.uk
    webmaster@digital.cabinet-office.gov.uk

They are unable to validate requests sent to any @service.gov.uk address.
