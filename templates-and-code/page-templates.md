---
layout: gsdm
title: Page templates
subtitle: HTML page templates for GOV.UK
audience:
  primary: designer, developer
  secondary: tech-arch,
type: resource
section: guidance
status: draft
---

* TOC
{:toc}
    
# Page templates

All pages on GOV.UK need to feel like they're part of the same service. In particular, the header and footer should be consistent across every page on the website.
This guide explains how to use GOV.UK template code and copy.

## The template code

Find the most up-to-date version of the template code in the Government Digital Service's ['Static' repository on GitHub](https://github.com/alphagov/static/tree/master/app/views/root). 

The [_base.html.erb](https://github.com/alphagov/static/blob/master/app/views/root/_base.html.erb) file contains the default template code. Ruby users can also make use of our [Prototyping app](https://github.com/alphagov/prototyping) to quickly mock up static prototypes using GOV.UK templates and assets.

## Template assets

The template code contains direct links to CSS and JavaScript assets hosted on the GOV.UK domain. We recommend that you leave these links as is whilst you develop your service (as opposed to linking to your own copies of these files). That way, you'll always be using the latest version of the assets and when they change you can identify and resolve any conflicts immediately.

Once the service is ready for production we'll need to decide whether you should continue to link to the assets in this way or whether you should now use your own copies of them. Please contact the GDS team to discuss this at the appropriate time.

## Keeping template code and assets up-to-date

GOV.UK is continuously being improved, which means that template and asset code changes regularly. All services on GOV.UK are expected to keep their templates and assets up-to-date. How you do this will depend on how you implement the templates and where your service is hosted. Please contact the GDS team to discuss options.

# Template elements

## The site header

This is the black strip across the top of every page on GOV.UK. It contains the GOV.UK logo and the search box.

Please don't add any other elements to the site header.

This element is mandatory.

### The GOV.UK logo

This must always link back to the [GOV.UK home page](http://gov.uk).

### The search box

In most cases the default action for the search box should be to perform a search on the whole of GOV.UK. If your service is to be hosted in the same environment as GOV.UK the default template code will do this.

If your service is to be hosted in a different environment you'll need to change the action attribute of the form as follows: `action="http://gov.uk/search"`.

In some cases we might decide that the scope of the search should be reduced to just pages from your service. In these cases the scope should be clearly indicated in the placeholder attribute and label of the search box.

In order for pages from your service to appear within GOV.UK search results they will need to be indexed. It doesn't always make sense to do this. For example, there's no point indexing a page that can only be reached via a login or as part of a transaction. If we decide that your pages *should* be indexed you'll need to contact the GDS team to discuss the best way of going about this.

## Breadcrumb navigation

All content pages on GOV.UK are given a category and subcategory. These are used to create category pages and breadcrumb navigation, which appears immediately below the site header. Breadcrumb navigation allows people to broaden their search for information from a specific page to all pages within the same category or subcategory. It also helps them to orientate themselves within the website.

This element is optional. Not all pages require breadcrumb navigation. It depends if the actions that it supports are genuine use cases for a specific page. For example, a page that's part-way through a multi-page transaction would not benefit from breadcrumbs because the user's main goal is completing the transaction, not browsing similar content.

If we decide that your page *should* have breadcrumb navigation you'll need to contact us to establish which categories should be assigned to your service. This will determine what the breadcrumb links are and on which category page the service appears.

## The main content area

This is where all the page content lives. All content within this area should be sympathetic to the GOV.UK visual style and use the available design patterns where possible.

There are two main variants of this element: A narrower version with a 'related links' panel on the right, for content pages, and a wider version without a panel for transaction and landing pages.

## The page footer

This is the grey panel at the bottom of every page. It contains global site navigation and a copyright message.

Please don't add any other elements to the footer. This element is mandatory.

From time to time we may update the links in the footer. It's important that your web pages are updated promptly when this happens. Contact the GDS team to discuss the best way of ensuring that this happens.

## Alpha and Beta templates

[The GOV.UK Beta warning](https://github.com/alphagov/static/blob/master/app/assets/javascripts/welcome.js) served users an interstitial when first arriving on GOV.UK during its public beta.

If you are running a public beta of a service then it is highly recommended that you do the same thing. Users of government services might not always be used to seeing incomplete or trial version of websites in the open, and it's important to make them aware that they may not be viewing the 'canonical' information they need.

The message read:

>Welcome to GOV.UK. From 17 October this website will replace Directgov and Business Link as the best place to find government services and information.  

>Until then, you can explore the website by using this experimental trial 'beta' version.

>PLEASE BE AWARE: this is a test website. It may contain inaccuracies or be misleading. <a href='http://www.direct.gov.uk'>Directgov</a> and <a href='http://businesslink.gov.uk'>Business Link</a> remain the official websites for government information and services.

Users were then prompted to accept that they had read the warning, which set a cookie so the message was not displayed repeatedly. 
