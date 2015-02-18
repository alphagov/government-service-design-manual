---
layout: detailed-guidance
title: Cookies
subtitle: When to use them, and when to tell users about them
category: making-software
type: guide
audience:
  primary:
  secondary: service-managers, designers, developers, tech-archs
status: draft
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
---


This short guide about cookies covers:

* what to keep in mind when including cookies into your services
* how and why you must notify users about cookies on your service

## Cookies explained

Cookies are small data files that are sent from a website and stored on a user's computer. They are used to store information that can be retrieved later in the visit or in future visits to the website.

Many uses of cookies are harmless, but sometimes they are used to track users and their browsing habits across multiple websites and target them with relevant advertising.

## Privacy and Electronic Communications Regulations

In May 2011, the [Privacy and Electronic Communications Regulations](https://ico.org.uk/for-organisations/guide-to-pecr/) were updated to require website operators to gain consent before storing or retrieving data from a user's computer (or other device). This change directly affects the use of cookies and other similar technologies, like HTML5 local storage.

Before using cookies, your website operator needs to either:

- get explicit, informed consent from the user before storing cookies on a user's computer
- be satisfied that the user understands that their actions will result in cookies being stored (implied consent)
- be satisfied that the cookie is "absolutely essential" to the operation of the website (eg cookies used for operating a shopping cart)

Responsibility for [complying with these regulations](https://ico.org.uk/for-organisations/guide-to-pecr/cookies/) lies with your website operator.

## Using cookies

This guide covers how to use cookies on government services, but the principles also apply to other technologies, like HTML5 local storage.

Minimise the use of cookies in your service, storing the smallest amount of information that you need for as short a time as necessary to provide a good service to users.

If your service needs to store cookies, you need to make sure that this can be explained simply and clearly, in a way that the majority of users can understand.

You must notify users that cookies are being stored.

## Types of cookies

This information covers 3 types of cookies:

* [first party cookies](#first-party-cookies)
* [third party cookies](#third-party-cookies)
* [exempt cookies](#exempt-cookies)

### First party cookies

These cookies are set by the website that the user is currently viewing. They’re under the control of the website operator and can only be accessed by the website. Data stored in the cookie is not shared with other websites.

Examples of first party cookies include:

- cookies for storing logged in status
- cookies for storing user preferences
- some types of analytics cookies

First party cookies are minimally intrusive as the website owner has complete control over:

* what data is stored within them
* how long the data is stored for
* what the data is used for

### Third party cookies

These cookies are set by external services used on the website and are:

* under the control of the third party service, not the website operator
* can be accessed on any website that makes use of the service
* can be used to track a user from one site to another

Examples of third party cookies include:

- cookies from social media sharing services
- cookies from advertising campaign management services
- cookies from embedded document sharing services (eg dropbox)
- cookies from some analytics services

Third party cookies are intrusive as the website owner usually has no control over what data is collected or how it is used.

### Exempt cookies

These cookies do not need to gain consent from a user. Exempt cookies have many uses, like:

* being used for [load balancing](https://en.wikipedia.org/wiki/Load_balancing_(computing))
* being absolutely essential to website functionality (eg used to store shopping cart contents)

Still notify users that these cookies are in use, even though they are exempt from the [Privacy and Electronic Communications Regulations](https://ico.org.uk/for-organisations/guide-to-pecr/).

## Cookie information and warnings

All services on the `service.gov.uk` subdomain must include a cookie information page. This page must contain information about the cookies used throughout the site, followed by an explanation of each cookie's purpose and how long it's stored for.

The [GOV.UK cookies page](https://www.gov.uk/support/cookies) is an example of how to do this. You must link back to your service's cookies page from the footer of your website. Your service information page must also link to the GOV.UK cookies page.

Your service must also tell users on their first visit that cookies are used and regularly remind them of this. This is particularly important when your service relies on implied consent. GOV.UK does this with a blue information banner that is displayed at least once every 3 months with this message:

> "GOV.UK uses cookies to make the site simpler. [Find out more about cookies](https://www.gov.uk/support/cookies)."

Where explicit consent is required, your service must notify their users before the cookie is set. To do this, use the `sets a cookie` text linked to the appropriate details on your cookie information page.

## Cookie scoping and attributes

Cookies must only apply to your originating domain name, eg `www.servicename.service.gov.uk` not `.gov.uk`.

Don't use cookies on domains that host only static assets (they [introduce a browser overhead that slows down the response time for users](https://developer.yahoo.com/performance/rules.html#cookie_free) without providing any benefit).

Cookies must be sent with the `Secure` attribute and, where appropriate, be sent with the `HttpOnly` attribute. These [flags provide additional assurances about how cookies will be handled by browsers](https://en.wikipedia.org/wiki/HTTP_cookie#Secure_and_HttpOnly).


## Further reading

This [blog post by GDS developer Dafydd Vaughan](https://gds.blog.gov.uk/2012/01/12/cookies-on-the-beta/) explains how cookies were used on the beta version of [GOV.UK](https://www.gov.uk).
