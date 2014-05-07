---
layout: detailed-guidance-2
title: GOV.UK templates
subtitle: GOV.UK header, footer, branding and typeface
category: user-centered-design
type: guide
audience:
  primary: designers, developers
  secondary:
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
    title: User centered design
    url: /service-manual/user-centered-design
---

{:.intro}
All pages on GOV.UK services should use the appropriate header, footer, branding and typeface.
The version you use depends on what kind of service you're making, and whether it's live or not.

---

## Alpha services

During the discovery and alpha phases of delivery you can use the GOV.UK templates, but you don't have to. The important thing is to make something and make it quickly, before testing it with real users. 

Follow the design principles and focus on the user journeys, but don't worry about making everything pixel perfect. 

[Use our HTML prototyping kit to get up and running quickly](https://github.com/tombye/express_prototype)

---

## Beta services

Once your service is in beta it should be using the templates you intend to use once it's live. The GOV.UK template code and assets are available on GitHub for a number of frameworks, including Rails and Play.

[Get the latest version of the templates on GitHub](https://github.com/alphagov/govuk_template)

---

## Live services

Live services on GOV.UK must use the GOV.UK header, footer and logotype and the New Transport typeface.
A service is considered 'on GOV.UK' if it's available at www.gov.uk/myservice, myservice.service.gov.uk or myblog.blog.gov.uk.

If your service isn’t on GOV.UK you can still follow the guidance, but your site shouldn’t identify as being part of GOV.UK and should not use the crown or the GOV.UK logotype in the header. This blog post talks about a good example of [a site that follows our design patterns](https://gds.blog.gov.uk/2013/03/18/intranets-dcms/) without using the logo.

---

## Headers

Content pages on GOV.UK use the standard header.
The GOV.UK logotype must always link back to [GOV.UK](https://www.gov.uk/).

### Standard header

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-0.png" alt="Standard header">
</div>

Transaction pages use a variant of the standard header. The search box is removed to avoid people leaving a transaction before they've completed it.

There are different versions of the transaction header.
Choose the one that most closely meets the needs of your users.


### Transaction header 1

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-1.png" alt="Header option 1">
</div>

Use this if your service is short, taking place over only a few screens.
You should also use this on the first page of your service if you are using the service title as the page heading.



### Transaction header 2

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-2.png" alt="Header option 2">
</div>

Use this if you need to remind your users of which service they are using.
The service title should link to the service start page (unless doing this harms the user experience).



### Transaction header 3

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-3.png" alt="Header option 3">
</div>

Use this if your service needs links for account management and so-on. 

---

## Footers

If you're creating content pages for GOV.UK, use the standard GOV.UK footer.
For transactional services, use the standard footer but remove all links other than the crown copyright and OGL licence notices.


