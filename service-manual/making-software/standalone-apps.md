---
layout: detailed-guidance
title: Standalone apps
subtitle: What is the government position on apps?
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

## What is the government position on apps?

At the Oct 2012 Digital Leaders’ meeting, the position was clarified: native apps could not be developed without Cabinet Office approval.

The Nov 2012 [Digital Strategy](http://www.publications.cabinetoffice.gov.uk/digital/strategy/) says:
“Stand-alone mobile apps will only be considered once the core web service works well on mobile devices, and if specifically agreed with the Cabinet Office”

Government’s position is that native & hybrid apps are currently rarely justified.

Ensure your service meets the [Digital by Default Service Standard](/service-manual/digital-by-default) and it will work well on mobile devices.

Make your data & API available for re-use and you will stimulate market if there is real demand for native apps.

## Brief background on apps landscape

Confusion is understandable.

So-called “Apps” come in several very, very different flavours.

## What is meant by ‘an app’?

  1. Device-specific ‘download and install’ apps (aka ‘native apps’)
  2. Websites that respond to various screen sizes (aka ‘responsive websites’, ‘web apps’ or ‘HTML5’)
  3. Various hybrids of the above

### What is a ‘Native’ App ?
  - Downloaded, installed piece of bespoke software.
  - Free or paid-for.
  - Persists on the device.
  - Can access all device features
  - Requires completely different software for iPhone/iPad, Android, Blackberry etc.
  - Spotify, Angry Birds, Instagram, Skype etc.

#### ‘Native’ App - Pros
  - Revenue (download and buy)
  - Persistent presence on device
  - Can access all functions on a device
  - Snappier performance in general
  - Can be used offline, in some cases.

#### ‘Native’ App - Cons
  - Expensive to develop & maintain
  - Needs several different versions (Android, iPhone, iPad, Blackberry etc)
  - Service iteration more complex (x3) -Can only be downloaded via gatekeeper app stores (Apple, Google)
  - Most apps are little downloaded, and even then rarely used.

### What is a‘ Web App’
  - Uses device’s built-in web browser
  - Is a website which optimises its layout & functionality for each device
  - Uses open standards (HTML5) -Examples: GOV.UK, PM’s dashboard, FT webapp, Virgin Active’s ‘My Locker’, bbc.co.uk/sport

#### Web App - Pros
  - It is your website, so costs are minimised and service iteration simplified.
  - Uses open standards (HTML5)
  - No gatekeepers to constrain access -Performance still good
  - Mobile web outstripping mobile app reach
  - Clear winning strategy for ‘utility’ services which do not require complex device features or persistence.

#### Web App - Cons
  - Not persistent on device
  - Some device features unavailable (camera, address book)
  - Requires internet connection
  - Not snappy enough for some complex services (e.g. Spotify, Facebook, Skype)
  - No ‘download and buy’ revenue stream.

### Hybrid App
  - A small native app which then loads up a bespoke website
  - So, can use device features (e.g. camera) that Web apps can’t.
  - Requires a stand-alone website Examples “Bing for Mobile”, Netflix app, LinkedIn app, BBC News app

#### Hybrid App – Pros & Cons
Not as expensive to maintain as a native app, can access device functions and be persistent but still requires a new parallel version of your web service, and multiple versions to be developed for each device.

## Rationale for government's position

If there is a market for native or hybrid apps, why should the government monopolise it? There is a vibrant market of 3rd party native app developers using government data & APIs. Government’s position is that native & hybrid apps currently rarely justified.

Ensure your service meets the [Digital by Default Service Standard](/service-manual/digital-by-default) and it will work well on mobile devices.

Make your [data](/service-manual/making-software/open-source.html) & [API](/service-manual/making-software/apis.html) available for re-use and you will stimulate market if there is real demand for native apps.

We are [backing open standards](/service-manual/making-software/open-standards-and-licensing.html) (HTML5) rather than risking proliferation of parallel versions of services as devices proliferate.

And while people spend as much time using apps as using mobile web the vast majority of app use is for gaming & social networking.

For “utility” needs, such as those met by government services, the mobile web is preferred to native apps

## Exceptions

But what about exceptions ie those times when only an app will do? 

We recognise that there'll be a few exceptions, to help you determine whether your case is likely to be considered an exception, consider the following: 

### Have the necessary conditions been met? 

  1. Your web service is already designed to be responsive 
  2. The service you're looking to build an app for is already open to third-parties via APIs or open data

NOTE: If these conditions are not in place, it is unlikely you'll obtain approval for your app proposal. If you believe there are compelling reasons why these conditions have not been met, please be prepared to provide them. 
 
### More questions

  3. [What's the user need](/service-manual/users/user-needs.html)? Please provide evidence for this.

  4. Are you sure there are no 3rd party native or hybrid apps that meet this user need? If 'yes' and condition 2 has been met, please provide your thoughts on why no 3rd party has developed a native/hybrid app.
   
  5. Is this user need of sufficient importance to (your users) justify the lifetime cost of your proposed app? If yes, how have you determined this? You might want to review articles within the service manual such as, [Know Your Users](/service-manual/users) and [Writing User Stories](/service-manual/agile/writing-user-stories.html). 

  6. Is there evidence of demand for this app amongst your target users?  If 'yes', please provide supporting  evidence. 

  7. Given the sheer number of apps out there is there evidence to suggest that this app will actually be downloaded? If 'yes', please provide supporting evidence e.g. examples of similar apps that have proven popular with your target audience and evidence of their popularity.

  8. Is there evidence of platform penetration amongst your users to justify building an app for the platform you're proposing to do this for? If 'yes', please provide supporting evidence e.g. analytics data that shows proportion of visitors to your content/service that currently access it using relevant devices
