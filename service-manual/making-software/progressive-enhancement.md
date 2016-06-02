---
layout: detailed-guidance
title: Progressive enhancement
subtitle: How to create pages that work regardless of browser capability
category: making-software
type: guide
audience:
  primary: designers, developers
  secondary:
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
exclude_from_search: true
---

When creating web pages, the only part of a page that you can rely upon working is the HTML. Even that can fail, but without HTML there is no web page and everything else becomes moot. The attitude towards building for the web with this in mind is called [progressive enhancement](https://en.wikipedia.org/wiki/Progressive_enhancement).

This means in essence that each extra layer (images, styling, behaviour, video, audio) of a page should be seen as optional. If you build pages with the idea that parts other than HTML are optional, you’ll create a better and stronger web page.

## First, just make it work

Make any new page or feature work with HTML alone -- no images, no CSS, no JavaScript, nothing but HTML. Interactive elements are only those capable of being implemented with [forms](https://en.wikipedia.org/wiki/Form_(HTML)) and server-side processing.

This gives a baseline experience, which will work in practically every browser. It allows your site to work for as many people and devices as possible, including older legacy browsers and devices.

## Second, make it work better

From this baseline, extra layers can then be added:

* the page can have images
* advance styling can be applied
* interactions can be made smoother and faster without the need to refresh the entire page
* validation of submitted data can be performed before it hits the network
* charts and data tables can be turned into visualisations with interactive elements

Remember to see each addition as just that -- an addition. Additions are just something extra that the more modern browsers are capable of doing. That’s along with an already accessible and usable experience that has as much content and interactivity available as possible.

## It isn't about "JavaScript off"

A common misunderstanding is that designing sites to work without CSS, JavaScript or anything else is because a user chooses to switch it off, considering it as an error or mistake on their part. It isn’t just a case of "fix your browser".

There are many scenarios where the extra layers can fail to load or are deliberately filtered, which can be due to:

* temporary network errors
* DNS lookup failures
* server that the resource is found on could be overloaded or down, and fail to respond in time or at all
* large institutions (eg banks and financial institutions, some government departments) having corporate firewalls that block, remove or alter content from the Internet
* mobile network providers resampling images and altering content so that load times faster and reduce bandwidth consumed
* antivirus and personal firewall software that will alter and/or block content

Of course, there are people who do turn off features in their browsers deliberately. Respect their decision and still provide them with a usable and useful service.

## It isn't only about JavaScript

Another common misunderstanding is that progressive enhancement is only about JavaScript - it’s not. Treat anything other than HTML as an enhancement, not just JavaScript.

Adding an image to a page requires suitable alternative text (and may also need a longer description) for people with visual impairment, and simply for when the image fails to load.

Styling cannot be the only way to share information:

* "text in bold" is not enough when you can't tell the difference between bold and normal text
* "red items are required" is not enough when you don't have colours applied, cannot see the colours or cannot distinguish them because of colour blindness
* video and audio without transcripts and/or subtitles are unsuitable for hearing impaired people
* information appearing in a video in a visual-only form (eg a question answered by people on screen with no audio) is unsuitable for people with visual impairments
* making an interactive element that requires a mouse (eg a hover effect, a drag and drop operation) is unusable by someone who only uses a keyboard or a touch-only device such as a smartphone or tablet

## Further reading

* [Progressive enhancement](https://en.wikipedia.org/wiki/Progressive_enhancement) on Wikipedia
* [Understanding progressive enhancement](http://alistapart.com/article/understandingprogressiveenhancement) on A List Apart
* [For a Future-Friendly Web](http://alistapart.com/article/for-a-future-friendly-web) on A List Apart
* [Cutting the mustard](http://responsivenews.co.uk/post/18948466399/cutting-the-mustard) on the BBC News responsive design blog
