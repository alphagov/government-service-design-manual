---
layout: detailed-guidance-2
title: Page templates
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
All pages on GOV.UK services should use the GOV.UK header, footer, typeface and branding.


## Template code

### Building services

The [latest version of the template code](https://github.com/alphagov/govuk_template) is available as a Ruby gem, or for the Play, Mustache and Liquid frameworks. The templates include assets and code for the site headers, footers and New Transport typeface.

### Prototyping

If you need to quickly prototype a service, you can [use this prototyping tool](https://github.com/tombye/express_prototype). 
It uses Node.js and Express to give you easy access to the templates, plus helpers for styling the [GOV.UK elements](http://govuk-elements.herokuapp.com/).



## Site header

Remember:

* only use the GOV.UK logotype if your service is [currently on GOV.UK](/service-manual/user-centered-design/what-should-service-look-like)
* always link the logotype back to [GOV.UK](https://www.gov.uk/)
* don't include the GOV.UK search box in transactions, as it will encourage people to leave the transaction mid-way

Here are three versions of the site header. Choose the one that most closely meets the needs of your users.

### Option 1

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-1.png" alt="Header option 1">
</div>

Use this header if your service is short, taking place over only a few screens.

You should also use this on the first page of your service if you are using the service title as the page heading.



### Option 2: Service title

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-2.png" alt="Header option 2">
</div>

Use this header if you find that your users need reminding of which service they are accessing.

The service title should link to the service start page (unless doing this harms the user experience).



### Option 3: Service title and navigation

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-3.png" alt="Header option 3">
</div>

Use this header if your service needs internal navigation. 



## Site footer

For information services, use the standard GOV.UK footer, as used on the GOV.UK home page.

For transactional services, use the standard footer but remove all links other than the crown copyright and OGL licence notices.


