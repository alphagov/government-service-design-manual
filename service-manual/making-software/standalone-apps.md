---
layout: detailed-guidance
title: Standalone mobile apps
subtitle: What is the government position on mobile apps?
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

## What is the government's position on apps?

The government’s position is that native and hybrid apps are currently rarely justified. At the October 2012 Digital Leaders’ meeting, the position was clarified: native apps could not be developed without Cabinet Office approval. The November 2012 [Digital Strategy](http://www.publications.cabinetoffice.gov.uk/digital/strategy/) says:

> “Stand-alone mobile apps will only be considered once the core web service works well on mobile devices, and if specifically agreed with the Cabinet Office”

Ensuring your service meets the [Digital by Default Service Standard](/service-manual/digital-by-default) means it will work well on mobile devices. Making your data and API available for re-use will stimulate the market if there is real demand for native apps.

## Brief background on apps landscape

So-called “apps” come in several different flavours, therefore confusion is understandable. When people are talking about "apps", they can mean device-specific "download and install" apps (aka native apps) or websites that respond to various screen sizes (aka responsive websites, web apps or HTML5) or even various hybrids of the two.

### What is a native app?

Native apps are downloadable software applications that run using the device's operating system code and APIs.

They persist on the device and can access all the hardware features (camera, storage, phone capabilities, etc). Because they run using native code, different versions must be created for each operating system. Examples of native apps include Spotify, Angry Birds, Instagram and Skype.

#### Pros
- Revenue (download and buy)
- Persistent presence on device
- Can access all functions on a device
- Snappier performance in general
- Can be used offline, in some cases

#### Cons
- Expensive to develop and maintain
- Needs several different versions (Android, iPhone, iPad, Blackberry etc)
- Service iteration more complex (at least 3 types of app to deploy)
- Can only be downloaded via gatekeeper app stores (Apple, Google)
- Most apps are rarely downloaded, and even then hardly used

### What is a web app?

A web app is built using open web standards (HTML, CSS, javascript, etc) and is run inside a device's web browser. Web apps can be downloaded locally or accessible over the internet.

Examples are GOV.UK, PM's dashboard, FT webapp, Virgin Active's 'My Locker', bbc.co.uk/sport.

#### Pros
- It is your website, so costs are minimised and service iteration simplified
- Uses open standards (HTML5)
- No gatekeepers to constrain access
- Performance still good
- Mobile web outstripping mobile app reach
- Clear winning strategy for ‘utility’ services which do not require complex device  features or persistence

#### Cons
- Not persistent on device
- Some device features unavailable (camera, address book)
- Requires internet connection
- Not snappy enough for some complex services (eg Spotify, Facebook, Skype)
- No ‘download and buy’ revenue stream

## Rationale for the government's position

If there is a market for native or hybrid apps, why should the government monopolise it? There is a vibrant market of 3rd party native app developers using government data and APIs. Government’s position is that native and hybrid apps currently rarely justified.

We are [backing open standards](/service-manual/making-software/open-standards-and-licencing.html) rather than risking proliferation of parallel versions of services as devices proliferate.

And while people spend as much time using apps as using mobile web the vast majority of app use is for gaming and social networking. For “utility” needs, such as those met by government services, the mobile web is preferred to native apps

## Exceptions

We recognise that there’ll be a few exceptions. To help you assess whether your case is likely to be considered an exception, consider the following:

### Have you met the necessary conditions?

**Condition 1:** Your web service is already designed to be responsive

**Condition 2:** The service or the content you’re looking to build an app for is already open to third-parties via APIs or as open data

> **NOTE:** If these conditions are not in place, it is unlikely that your app proposal will be approved. If you believe there are _compelling_ reasons why these conditions have not been met, please set them out in your proposal.

### More questions

1. [What’s the user need?](/service-manual/users/user-needs.html) Please provide supporting evidence
2. Which 3rd-party native/hybrid apps have already been developed to meet this user need? If  there are none and condition 2 has been met, please provide your thoughts on why this might be the case. If there are 3rd-party alternatives, please state why you believe a government-developed app is required.
3. Is this user need of sufficient importance to (your users to) justify the lifetime cost of your proposed app? If you believe it is, how have you determined this? You might find it useful to review articles within the service manual such as, [Know Your Users](/service-manual/users) and [Writing User Stories](/service-manual/agile/writing-user-stories.html).
4. Is there evidence of demand for this type of app amongst your target users? If you believe there is, please provide supporting evidence e.g. examples of similar apps that have proven popular with your target audience and evidence of their popularity.
5. Is there evidence to justify building an app for the platform you’re proposing to do this for? If so, please provide supporting evidence e.g. analytics data that shows proportion of visitors to your content/service that currently access it using relevant devices.

### The Process

In addition to the evidence requested above, all digital spend for the development of standalone mobile apps is subject to the GDS spend approval. [Guidance](https://www.gov.uk/government/publications/cabinet-office-controls-guidance-version-3-1) (including details about [response times](https://www.gov.uk/government/uploads/system/uploads/attachment_data/file/60699/Annex-4-2-Service-Level-Agreement.doc)) on the process can be found on GOV.UK. If you have any queries, [please contact GDS](mailto:pmo@digital.cabinet-office.gov.uk).

## Further Reading

* [Government approach to apps presentation](http://www.slideshare.net/DigEngHMG/government-approach-to-apps)
* Tom Loosemoore's [blog post about standalone apps](http://digital.cabinetoffice.gov.uk/2013/03/12/were-not-appy-not-appy-at-all/)
