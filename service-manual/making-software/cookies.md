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


This short guide tells you what to keep in mind when including cookies into your services, and how and why you must notify users about cookies on your service.

## Cookies explained

Cookies are small data files that are sent from a website and stored on a user's computer. They are used to store information that can be retrieved later in the visit or in future visits to the website.

Many uses of cookies are harmless, but sometimes they are used to track users and their browsing habits across multiple websites and target them with relevant advertising.

## Privacy and Electronic Communications Regulations

In May 2011, the [Privacy and Electronic Communications Regulations](http://ico.org.uk/for_organisations/privacy_and_electronic_communications) were updated to require website operators to gain consent before storing or retrieving data from a user's computer (or other device). This change directly affects the use of cookies and other similar technologies such as HTML5 local storage.

Before using cookies the website operator needs to either:

- get explicit, informed consent from the user before storing cookies on a user's computer
- be satisfied that the user understands that their actions will result in cookies being stored (implied consent)
- be satisfied that the cookie is "absolutely essential" to the operation of the website (eg cookies used for operating a shopping cart)

Responsibility for complying with these regulations lies with the website operator.
[The Information Commissioner's Office provides guidance on cookies](http://ico.org.uk/for_organisations/privacy_and_electronic_communications/the_guide/cookies).

## Using cookies

This guide covers how to use cookies on government services, but the principles also apply to other technologies such as HTML5 local storage.

You should minimise the use of cookies on services, store as little information as you require for as short a time as necessary to provide a good service to users.

If your service requires cookies to be stored then you need to make sure that they can be explained simply and clearly, in a way that the majority of users can understand.

You must notify users that cookies are being stored.

## Types of cookies

### First party cookies

These are cookies that are set by the website that the user is currently viewing. They are under the control of the website operator and can only be accessed by the website. Data stored in the cookie is not shared with other websites.

Examples of first party cookies include:

- cookies for storing logged in status
- cookies for storing user preferences
- some types of analytics cookies

These types of cookies are minimally intrusive as the website owner has complete control over what data is stored within them, how long the data is stored for and what the data is used for.


### Third party cookies

These are cookies set by external services used on the website. The cookies are under the control of the third party service and can be accessed on any website that makes use of the service. These cookies are not controlled by the website operator and can be used to track a user from one site to another.

Examples of third party cookies include:

- cookies from social media sharing services
- cookies from advertising campaign management services
- cookies from embedded document sharing services
- cookies from some analytics services

These types of cookies are intrusive as the website owner usually has no control over what data is collected or how it is used.

### Exempt cookies

A number of uses of cookies are exempt from the requirement to gain consent. These include cookies that are used for [load balancing](http://en.wikipedia.org/wiki/Load_balancing_(computing)) or cookies that are "absolutely essential" to the use of a website (eg used to store shopping cart contents).

While these cookies are exempt from the Privacy and Electronic Communications Regulations, you should still notify users that these cookies are in use.

## Cookie information and warnings

All services on the `service.gov.uk` subdomain must include a cookie information page. This page must contain information about the cookies used throughout the site, followed by an explanation of each cookie's purpose and how long it stored for.

You can see an example of how to do this on the [GOV.UK cookies page](https://www.gov.uk/support/cookies).

Each service must include a link to this page on the footer of the website. The information page must also include a link back to the main GOV.UK cookies page.

Services must also tell users on their first visit that cookies are used and regularly remind them of this. This is particularly important when the service relies on implied consent. GOV.UK does this with a blue information banner that is displayed at least once every 3 months with this message:

> "GOV.UK uses cookies to make the site simpler. [Find out more about cookies](https://www.gov.uk/support/cookies)."

Where explicit consent is required, services must notify their users before the cookie is set. You should do this with the `sets a cookie` text linked to the appropriate details on the cookie information page.

## Cookie scoping and attributes

Cookies must be scoped to their originating domain name only eg `www.servicename.service.gov.uk` not `.gov.uk`.

Cookies should not be used on domains that host only static assets (they [introduce a browser overhead that slows down the response time for users](https://developer.yahoo.com/performance/rules.html#cookie_free) without providing any benefit).

Cookies must be sent with the `Secure` attribute and should, where appropriate, be sent with the `HttpOnly` attribute. These [flags provide additional assurances about how cookies will be handled by browsers](https://en.wikipedia.org/wiki/HTTP_cookie#Secure_and_HttpOnly).


## Further reading

This [blog post by GDS developer Dafydd Vaughan](https://gds.blog.gov.uk/2012/01/12/cookies-on-the-beta/) explains how cookies were used on the beta version of [GOV.UK](https://www.gov.uk).
