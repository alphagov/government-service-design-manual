---
layout: detailed-guidance
title: CAPTCHA
subtitle: Build to the GOV.UK style
category: design-and-development-resources
type: resource
audience:
    primary: designers
    secondary: developers
status: draft
phases:
  - alpha
  - beta
  - live
page_class: buttons
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

[CAPTCHA](https://en.wikipedia.org/wiki/CAPTCHA) stands for Completely Automated Public Turing test to tell Computers and Humans Apart. They are used to prevent bots (automated software) from completing a form or accessing a system and usually take the form of jumbled up text for the user to decipher and enter before submitting a form.

## Why you shouldn't use them

CAPTCHAs introduce significant problems to online services:

* **Usability** -- they put the burden of detecting bots on the user rather than the system. As CAPTCHAs are designed to be hard to read and understand, this makes the service much more difficult to use.
* **Accessibility** -- they are inaccessible by design. This effectively makes the service unusable by people with certain disabilities. Even CAPTCHAs that provide audio versions do not completely resolve this issue.

Additionally, if a 3rd party CAPTCHA service is used, there are further problems to consider:

* **Privacy** -- 3rd party CAPTCHA services set cookies, collect analytics and can track users across multiple sites. This introduces significant privacy concerns
* **Performance** -- use of a 3rd party CAPTCHA service ties your performance to theirs. If their service goes offline, so does access to your service
* **Security** -- the security of your service is tied to that of the 3rd party. If they are compromised, your service and its users may also be

## Alternatives to CAPTCHA

Many of the risks that CAPTCHAs are aimed to mitigate can be addressed in other ways:

* rate and connection limiting
* use of honey pots
* protective monitoring

It's important to note that even with a CAPTCHA in place bots will still get through due to advances in computer imaging and the use of CAPTCHA farms. A combination of different approaches generally gives the best results.

## Further reading

* [CAPTCHA and the BBC](http://www.bbc.co.uk/blogs/legacy/bbcinternet/2010/10/captcha_and_bbc_id.html)
* [Ticketmaster ditches CAPTCHA for something simpler](http://thenextweb.com/insider/2013/01/30/good-news-music-fans-ticketmaster-is-ditching-its-captcha-conundrums-for-something-simpler/)
