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

The government’s position is that native & hybrid apps are currently rarely justified. At the October 2012 Digital Leaders’ meeting, the position was clarified: native apps could not be developed without Cabinet Office approval. The November 2012 [Digital Strategy](http://www.publications.cabinetoffice.gov.uk/digital/strategy/) says:

> “Stand-alone mobile apps will only be considered once the core web service works well on mobile devices, and if specifically agreed with the Cabinet Office”

Ensuring your service meets the [Digital by Default Service Standard](/service-manual/digital-by-default) means it will work well on mobile devices. Making your data & API available for re-use will stimulate the market if there is real demand for native apps.

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
- Expensive to develop & maintain
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

If there is a market for native or hybrid apps, why should the government monopolise it? There is a vibrant market of 3rd party native app developers using government data & APIs. Government’s position is that native & hybrid apps currently rarely justified.

We are [backing open standards](/service-manual/making-software/open-standards-and-licencing.html) rather than risking proliferation of parallel versions of services as devices proliferate.

And while people spend as much time using apps as using mobile web the vast majority of app use is for gaming & social networking. For “utility” needs, such as those met by government services, the mobile web is preferred to native apps
