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
    url: /service-manual/index.html
  -
    title: Design and content
    url: /service-manual/design-and-content/index.html
---

{:.intro}
All pages on GOV.UK should use the GOV.UK header, footer, branding and the New Transport typeface.

### On this page:

1. [Using the templates](#using-the-templates)
2. [Get the templates](#get-the-templates)
3. [Choose a header](#choose-a-header)
4. [Choose a footer](#choose-a-footer)

---

<h2 class="heading-36" id="using-the-templates">1. Using the templates</h2>


You can use the templates if the service you're devloping is going to be on GOV.UK.
Live or beta GOV.UK services MUST use the templates.

A service is ‘on GOV.UK’ if it’s available at www.gov.uk/myservice, myservice.service.gov.uk or myblog.blog.gov.uk.

If your service isn’t going to be on GOV.UK you can still follow this guidance, but your site shouldn’t identify as being part of GOV.UK and should not use the crown or the GOV.UK logotype in the header. 

[Alpha](/service-manual/phases/alpha) GOV.UK services can use the GOV.UK templates, but they don’t have to.
The important thing is to make something and make it quickly, before testing it with real users.  


---

<h2 class="heading-36" id="get-the-templates">2. Get the templates</h2>

The GOV.UK template code and assets are available for the Play or Rails frameworks, and for the Mustache, Liquid or Jinja templating languages.

[Get the latest version of the templates on GitHub](https://github.com/alphagov/govuk_template)

---


<h2 class="heading-36" id="choose-a-header">3. Choose a header</h2>

There are different versions of the GOV.UK header for different types of page.

The basic header for services on GOV.UK is this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-1.png" alt="Header option 1">
</div>

Remember:

* always link the GOV.UK logotype back to the GOV.UK home page
* set the coloured bar width to be the same as the content area. The default is 960px at full width
* the coloured bar is always $govuk-blue (#005ea5)

If your service is more than a few pages long, you can help users understand where they are by adding a service title, like this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-2.png" alt="Header option 2">
</div>

Finally, if you need to include contact or account management links, you can do so like this:

<div class="example">
  <img src="/service-manual/assets/images/header-footer/header-pattern-3.png" alt="Header option 3">
</div>

---

<h2 class="heading-36" id="choose-a-footer">4. Choose a footer</h2>

If you're creating content pages for GOV.UK, use the standard GOV.UK footer, as used on the home page.
For transactional pages, use the standard footer but remove all links other than the crown copyright and OGL licence notices.
