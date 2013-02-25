---
layout: detailed-guidance
title: Page templates
subtitle: How to use the elements that make up the GOV.UK page templates
audience:
  primary: designer, developer
  secondary: tech-arch,
type: resource
section: guidance
status: draft
---

All new government services should have a consistent appearance. The best way 
to do this is to use the same [template code and assets](shared-assets-libraries).
This guide describes each element in the page template and explains how it should
be used.

{% include _toc.md %}

## Site header

{:.left}
![Image of site header](/assets/images/page-template-elements/header.png)

This is the strip across the top of every page on GOV.UK. It contains the
GOV.UK logo and the search box. This element is mandatory and should not be added to.

### GOV.UK logo

{:.left}
![Image of GOV>UK logo](/assets/images/page-template-elements/logo.png)

This must always link back to the [GOV.UK home page](http://gov.uk).

### Search box

{:.left}
![Image of search box](/assets/images/page-template-elements/search.png)

In most cases the default action for the search box should be to perform a
search on the whole of GOV.UK. If your service is to be hosted in the same
environment as GOV.UK the default template code will do this.

If your service is to be hosted in a different environment you'll need to change
the action attribute of the form as follows: `action="http://gov.uk/search"`.

In some cases we might decide that the scope of the search should be reduced to
just pages from your service. In these cases the scope should be clearly
indicated in the placeholder attribute and label of the search box.

In order for pages from your service to appear within GOV.UK search results they
will need to be indexed. It doesn't always make sense to do this. For example,
there's no point indexing a page that can only be reached via a login or as part
of a transaction. If we decide that your pages *should* be indexed you'll need
to contact the GDS team to discuss the best way of going about this.

## Breadcrumb navigation

{:.left}
![Image of breadcrumb navigation](/assets/images/page-template-elements/breadcrumb.png)

All content pages on GOV.UK are given a category and subcategory. These are used
to create category pages and breadcrumb navigation, which appears immediately
below the site header. Breadcrumb navigation allows people to broaden their
search for information from a specific page to all pages within the same
category or subcategory. It also helps them to orientate themselves within the
website.

This element is optional. Not all pages require breadcrumb navigation. It
depends if the actions that it supports are genuine use cases for a specific
page. For example, a page that's part-way through a multi-page transaction would
not benefit from breadcrumbs because the user's main goal is completing the
transaction, not browsing similar content.

If we decide that your page *should* have breadcrumb navigation you'll need to
contact us to establish which categories should be assigned to your service.
This will determine what the breadcrumb links are and on which category page the
service appears.

## Main content area

This is where all the page content lives. All content within this area should be
consistent with the GOV.UK visual style and use the available 
[design patterns](/design-patterns).

There are two main variants of this element: A narrower version with a 'related
links' panel on the right, for content pages, and a wider version without a
panel for transaction and landing pages.

## Page footer

{:.left}
![Image of page footer](/assets/images/page-template-elements/footer.png)

This is the grey panel at the bottom of every page. It contains global site
navigation and a copyright message. This element is mandatory and should not be 
added to.

From time to time we may update the links in the footer. It's important that
your web pages are updated promptly when this happens. Contact the GDS team to
discuss the best way of ensuring that this happens.

## Alpha and Beta templates

[The GOV.UK Beta warning](https://github.com/alphagov/static/blob/master/app/ass
[ets/javascripts/welcome.js) served users an interstitial when first arriving on
[GOV.UK](www.gov.uk) during its public beta.

If you are running a public beta of a service then it is highly recommended that
you do the same thing. Users of government services might not always be used to
seeing incomplete or trial version of websites in the open, and it's important
to make them aware that they may not be viewing the 'canonical' information they
need.

The message read:

>Welcome to GOV.UK. From 17 October this website will replace Directgov and
>Business Link as the best place to find government services and information.

>Until then, you can explore the website by using this experimental trial 'beta'
>version.

>PLEASE BE AWARE: this is a test website. It may contain inaccuracies or be
>misleading. <a href='http://www.direct.gov.uk'Directgov</a and <a
>href='http://businesslink.gov.uk'Business Link</a remain the official websites
>for government information and services.

Users were then prompted to accept that they had read the warning, which set a
cookie so the message was not displayed repeatedly.
