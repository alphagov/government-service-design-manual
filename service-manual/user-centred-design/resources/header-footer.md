---
layout: detailed-guidance
title: GOV.UK header and footer
subtitle: How to make your service look like GOV.UK
category: user-centred-design
type: guide
audience:
  primary: designers
  secondary: developers
status: draft
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

{:.intro}
This guide explains when and how to use the GOV.UK header, footer, logo and typeface.

---

## When to use the GOV.UK header and footer

If your service is available from www.gov.uk, service.gov.uk or blog.gov.uk then it must use the GOV.UK header and footer. This includes the crown device, the GOV.UK logotype and the New Transport typeface.

If it's not available from one of these domains (including other gov.uk domains like data.gov.uk) then it must not use the GOV.UK header and footer.

If your service isn't yet in beta then you don't need to use the GOV.UK header and footer. The important thing is to make something and make it quickly, before testing it with real users.

---

## Header and footer code

All the code and assets you need to implement the header and footer are available via the [GOV.UK template app on GitHub](https://github.com/alphagov/govuk_template).

---

## Choosing the right header

There are different versions of the GOV.UK header for different types of page. For simple services on GOV.UK use this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-1.png" alt="Header option 1">
</div>

Remember:

* always link the GOV.UK logotype back to the GOV.UK home page
* set the coloured bar width to be the same as the content area. The default is 960px at full width
* the coloured bar is always $govuk-blue (#005ea5)


If your service is more than a few pages long, you can help users understand where they are by adding the service title, like this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-2.png" alt="Header option 2">
</div>

If you need to include contact or account management links, do this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-3.png" alt="Header option 3">
</div>

[Here's how to code the service title and navigation](https://github.com/alphagov/govuk_template#propositional-title-and-navigation)

---

## Choosing the right footer

Services on GOV.UK should use the standard GOV.UK home page footer, but with the browse links removed. You may also add a link back to the department responsible and to other language options.

Here's an example from the 'Renew a tax disc' service:


<div class="example">
  <a href="/service-manual/assets/images/header-footer/tax-disc-footer.png">
  <img src="/service-manual/assets/images/header-footer/tax-disc-footer.png" alt="An example footer from the Renew a tax disc service">
  </a>
</div>


