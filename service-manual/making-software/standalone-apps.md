---
layout: detailed-guidance
title: Standalone mobile apps
subtitle: Government's position on mobile apps
category: making-software
type: guide
audience:
  primary: service-managers
  secondary: designers, developers, performance-analysts, qa, content-designers
phases:
  - discovery
  - alpha
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

## The government's position on apps

The government’s position is that native apps are currently rarely justified. At the October 2012 Digital Leaders meeting, the position was clarified: native apps could not be developed without Cabinet Office approval. The November 2012 [Digital Strategy](/government/publications/government-digital-strategy) says:

> “Standalone mobile apps will only be considered once the core web service works well on mobile devices, and if specifically agreed with the Cabinet Office”

Making sure your service meets the [Digital by Default Service Standard](/service-manual/digital-by-default) means it will work well on mobile devices. Making your data and application programming interface (API) available for reuse will stimulate the market if there is real demand for native apps.

## Types of apps

So-called 'apps' come in several different forms, which makes confusion understandable. When people are talking about 'apps', they can mean:

* device-specific 'download and install' apps (aka native apps)
* websites that respond to various screen sizes (aka responsive websites, web apps or HTML5)
* various combinations of the two

### Native apps

Native apps are downloadable software applications that run using the device's operating system code and APIs.

Native apps remain on the device and can access all the hardware features (camera, storage, phone capabilities etc). Because they run using native code, different versions must be created for each operating system.
Examples of native apps include Spotify, Angry Birds, Instagram, and Skype.

Pros of native apps:

- generate revenue (download and buy)
- have a constant presence on device
- can access all functions on a device
- better performance in general
- can be used offline, in some cases

Cons of native apps:

- expensive to develop and maintain
- needs several different versions (Android, iPhone, iPad, Blackberry etc)
- service iteration is more complex (at least 3 types of app to deploy)
- can only be downloaded via gatekeeper app stores (Apple, Google)
- most apps are rarely downloaded, and even then hardly used

### Responsive web design

Responsive web design is a design approach that optimises users' viewing experiences across a wide range of devices.

When a responsive website is accessed via a mobile phone, it's sometimes referred to as a ‘mobile web app’.
Responsive websites are built using open web standards (HTML, CSS, JavaScript etc) and they run inside a device’s web browser.

Examples include [GOV.UK](https://www.gov.uk), PM's dashboard, [Financial Times app](http://app.ft.com/), and [BBC Sport](http://www.bbc.co.uk/sport/0/).

Pros of responsive web design:

- it’s your website, so costs are minimised and service iteration simplified
- uses open standards (HTML5)
- no gatekeepers to restrict access
- performance is maintained
- mobile web exceeding the reach of mobile apps
- clear winning strategy for ‘utility’ services which do not require complex device features or persistence

Cons of responsive web design:

- not continuously on device
- some device features unavailable (camera, address book)
- requires internet connection
- not snappy enough for some complex services (eg Spotify, Facebook, Skype)
- no ‘download and buy’ revenue stream

## Reasons for government's position

If there’s a market for native apps, why should government monopolise it? There’s a vibrant market of third party native app developers using government data and APIs. The government’s position is that native and hybrid apps are currently rarely justified.

GDS [back open standards](/service-manual/making-software/open-standards-and-licensing.html) -- this removes the risk of having to create different, yet similar, versions of a service so an app is compatible with new devices that enter the market.

While people spend as much time using apps as using mobile web, the vast majority of app use is for gaming and social networking. For 'utility' needs, such as those met by government services, the mobile web is preferred to native apps.

## Exceptions

GDS recognise that there'll be a few exceptions. To help you assess whether your case is likely to be considered an exception, consider if you've met the following conditions:

* condition 1 -- your web service is already designed to be responsive
* condition 2 -- the service or the content you're looking to build an app for is already open to third-parties via APIs or as open data

If these conditions are not in place, it's unlikely that your app proposal will be approved. If you believe there are compelling reasons why these conditions have not been met, please set them out in your proposal.

### Setting out your proposal

In your proposal, please provide answers and evidence on the following:

* [what’s the user need?](/service-manual/user-centred-design/user-needs.html)
* which third-party native/hybrid apps have already been developed to meet this user need?
  * if there are none and condition 2 has been met, please provide your thoughts on why this might be the case
  * if there are third-party alternatives, please state why you believe a government-developed app is required
* is this user need of sufficient importance to (your users to) justify the lifetime cost of your proposed app?
  * if you believe it is, how have you determined this? You might find it useful to review articles within the service manual such as, [know your users](/service-manual/user-centred-design) and [writing user stories](/service-manual/agile/writing-user-stories.html)
* is there evidence of demand for this type of app amongst your target users? (provide supporting evidence, eg similar apps that have proven popular with your target audience and evidence of their popularity)
* is there evidence to justify building an app for the platform you’re proposing to do this for? (please provide supporting evidence, eg analytics data that shows proportion of visitors to your content/service that currently access it using relevant devices)

### The process

In addition to the evidence requested above, all digital spend for the development of standalone mobile apps is subject to the GDS spend approval. [Guidance](https://www.gov.uk/government/publications/cabinet-office-controls) (including details about response times) on the process can be found on GOV.UK. If you have any queries, [please contact GDS](mailto:pmo@digital.cabinet-office.gov.uk).

## Further reading

* Tom Loosemoore's [blog post about standalone apps](https://gds.blog.gov.uk/2013/03/12/were-not-appy-not-appy-at-all/)

*[APIs]: application programming interfaces
*[API]: application programming interface

