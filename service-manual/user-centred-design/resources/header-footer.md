---
layout: detailed-guidance
title: Header and footer
subtitle: How to implement the GOV.UK templates
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
All pages on GOV.UK services should use the GOV.UK header, footer, branding and New Transport typeface, as specified in the templates.

---

## 1. Check that your service is on GOV.UK

A service is ‘on GOV.UK’ if it’s available at www.gov.uk/myservice, myservice.service.gov.uk or myblog.blog.gov.uk.

If your service isn’t on GOV.UK you can still follow this guidance, but your site shouldn’t identify as being part of GOV.UK and should not use the crown or the GOV.UK logotype in the header. 

[Here's a good example](https://gds.blog.gov.uk/2013/03/18/intranets-dcms/) of a site that follows our guidance without using our identity.

---

## 2. Get the templates

The GOV.UK template code and assets are available for the Play or Rails frameworks, and for the Mustache, Liquid or Jinja templating languages. You can also generate the static files.

[Get the latest version of the templates on GitHub](https://github.com/alphagov/govuk_template)

---

## 3. Choose the right approach

During the [alpha](/service-manual/phases/alpha) phase of delivery the important thing is to make something and make it quickly, before testing it with real users. You can use the GOV.UK templates, but you don’t have to. 

Follow the design principles and focus on the user journeys, but don’t worry about making everything pixel perfect.

Once your service is in [beta](/service-manual/phases/beta) it must use the GOV.UK templates, including the header, footer, logotype and the New Transport typeface.

### Public alphas and betas

If your alpha or beta service is publicly available you should add a banner below the header explaining that the content and features are suject to change, and inviting user feedback.

<div class="example">
  <img src="/service-manual/assets/images/header-footer/alpha-example.png" alt="An example of an alpha service header">
</div>

[The README for the GOV.UK Frontend Toolkit](https://github.com/alphagov/govuk_frontend_toolkit#alphabeta) explains how to implement these banners.


---

## 4. Choose a header

There are different versions of the GOV.UK header for different types of page.

The standard header for information pages on GOV.UK is this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-0.png" alt="Standard header">
</div>

Remember:

* always link the GOV.UK logotype back to the GOV.UK home page
* set the coloured bar width to be the same as the content area. The default is 960px at full width
* the coloured bar is always $govuk-blue (#005ea5)

Pages that are part of a transaction should remove the search box, so people can't accidentally leave before they've completed it:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-1.png" alt="Header option 1">
</div>

If your service is more than a few pages long, you can help users understand where they are by adding the service title, like this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-2.png" alt="Header option 2">
</div>

Finally, if you need to include contact or account management links, you can do so like this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-3.png" alt="Header option 3">
</div>

[Here's how to code the service title and navigation](https://github.com/alphagov/govuk_template#propositional-title-and-navigation)


---

## 5. Choose a footer

If you're creating content pages for GOV.UK, use the standard GOV.UK footer, as used on the home page.
For transactional pages, use the standard footer but remove all links other than the crown copyright and OGL licence notices.
